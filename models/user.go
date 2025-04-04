package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"-"`

	UserID    int       `json:"user_id" form:"user_id" sql:"auto_increment" bun:"user_id,pk,autoincrement,unique"`
	FirstName string    `json:"first_name" form:"first_name" sql:"not_null"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	BOoks     []*Book   `bun:"m2m:user_book,join:User=Book"`
}

func (u *User) RestPath() string {
	return "user"
}

func (u *User) TableName() string {
	return "users"
}
