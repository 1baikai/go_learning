package entity

type User struct {
	ID       int64  `json:"id" gorm:"id"`
	Username string `json:"username" gorm:"username"`
	Fraction string `json:"fraction" gorm:"fraction"`
}
