package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Phone    string
	Password string
	Super    bool
}

func (u User) ToDt() (dto UserDto) {
	dto.ID = u.ID
	dto.Name = u.Name
	dto.Email = u.Email
	dto.Phone = u.Phone
	dto.Password = u.Password
	dto.Super = u.Super

	return
}
