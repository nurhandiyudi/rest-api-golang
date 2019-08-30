package structs

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Email      string `json:email"size:255"`
	Password   string `json: password`
	First_Name string `json: first_name`
	Last_Name  string `json: last_name`
}
