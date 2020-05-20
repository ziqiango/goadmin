package models

type User struct {
	ID        uint32 `gorm:"primary_key"`
	Name string `gorm:"type:varchar(16)"`
}

func (User) TableName() string  {
	return "user"
}
