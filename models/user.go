package models

type User struct {
	Id        int
	Name      string
	Telephone string
	Password  string
}

// TableName 防止去操作users表
func (User) TableName() string {
	return "user"
}
