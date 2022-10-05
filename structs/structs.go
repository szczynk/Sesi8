package structs

import (
	"errors"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	First_Name string `json:"first_name" binding:"required"`
	Last_Name  string `json:"last_name" binding:"required"`
}

func (p *Person) BeforeCreate(tx *gorm.DB) (err error) {
	// fmt.Println("Person BeforeCreate()")
	if len(p.First_Name) < 4 {
		err = errors.New("Person first name is too short")
	}
	if len(p.Last_Name) < 4 {
		err = errors.New("Person last name is too short")
	}
	return
}
