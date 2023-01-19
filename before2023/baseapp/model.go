package main

import "time"

type Base struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (p *Base) IsPersistent() bool {
	return p.ID > 0
}

func (p *Base) IsDelete() bool {
	return p.DeletedAt != nil
}
