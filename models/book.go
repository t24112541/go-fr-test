package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Book struct {
	bun.BaseModel `bun:"table:books"`

	BookID    int       `json:"book_id" form:"book_id" sql:"auto_increment" bun:"book_id,pk,autoincrement,unique"`
	Name      string    `json:"name" form:"name" sql:"not_null" bun:"name,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete"`
	Users     []*User   `bun:"m2m:user_books,join:Book=User"`
}

func (u *Book) RestPath() string {
	return "book"
}

func (u *Book) TableName() string {
	return "books"
}
