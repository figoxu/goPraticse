package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/figoxu/utee"
)

func main() {
	db, err := gorm.Open("postgres", "user=jianhui password=q1w2e3r4ASDF dbname=sdz_dev host=47.93.108.90 port=3432 sslmode=disable")
	utee.Chk(err)
	 db.LogMode(true)

	// Create Table
	type ClassRoom struct {
		gorm.Model
		State string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
	}

	// AutoMigrate
	db.DropTable(&ClassRoom{})
	db.Debug().AutoMigrate(&ClassRoom{})

	// JSON to insert
	STATE := `{"uses-kica": false, "hide-assessments-intro": true, "most-recent-grade-skew": 1.5}`

	classRoom := ClassRoom{State: STATE}
	db.Save(&classRoom)

	// Select the row
	var result ClassRoom
	db.First(&result)

	if result.State == STATE {
		fmt.Println("SUCCESS: Selected JSON == inserted JSON")
	} else {
		fmt.Println("FAILED: Selected JSON != inserted JSON")
		fmt.Println("Inserted: " + STATE)
		fmt.Println("Selected: " + result.State)
	}
}