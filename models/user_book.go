package models

type UserBook struct {
	Id     int   `json:"id" form:"id" sql:"auto_increment" bun:"id,pk,autoincrement,unique"`
	UserID int   `json:"user_id" form:"user_id" bun:",notnull"`
	BookID int   `json:"book_id" form:"book_id" bun:",notnull"`
	User   *User `bun:"rel:belongs-to,join:user_id=user_id"`
	Book   *Book `bun:"rel:belongs-to,join:book_id=book_id"`
}

func (u *UserBook) RestPath() string {
	return "user_book"
}

func (u *UserBook) TableName() string {
	return "user_book"
}
