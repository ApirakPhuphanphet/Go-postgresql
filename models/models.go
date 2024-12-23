package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Amount   int    `json:"amount"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
