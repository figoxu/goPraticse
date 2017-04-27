package main

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB

/* ----------------------------------------------------- */
//db, err := sqlx.Open("postgres", "postgres://postgres:postgres@192.168.56.101/test_db?sslmode=disable")

type Address struct {
	Home string
	Work string
}

type Person struct {
	Id      int             `db:"id"`
	Address json.RawMessage `db:"address"`
}

func (p *Person) CreateTable() {
	cmd := `CREATE TABLE people (
        id   SERIAL PRIMARY KEY,
        address  JSONB
    )
    `
	DB.MustExec(cmd)
}

func (p *Person) DropTable() {
	cmd := `DROP TABLE IF EXISTS people`
	DB.MustExec(cmd)
}

/* ----------------------------------------------------- */

func init() {
	db, err := sqlx.Connect("postgres", "user=Reza dbname=sample sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db
}

/* ----------------------------------------------------- */

func main() {
	p := Person{}
	p.DropTable()
	p.CreateTable()

	addresses := []Address{
		{"11 Home St", "11 Work St"},
		{"12 Home St", "12 Work St"},
		{"13 Home St", "13 Work St"},
	}

	// insert to the db
	tx := DB.MustBegin()
	for _, a := range addresses {
		b, err := json.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}

		j := types.JSONText(string(b))

		v, err := j.Value()
		if err != nil {
			log.Fatal(err)
		}

		tx.MustExec("INSERT INTO people (address) VALUES ($1)", v)
	}
	tx.Commit()

	// get the data back
	people := []Person{}
	if err := DB.Select(&people, "SELECT id,address FROM people"); err != nil {
		log.Fatal(err)
	}

	for i, p := range people {
		log.Printf("%d => %v , %v", i, p.Id, string(p.Address))
	}
}