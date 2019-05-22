package structs

import "github.com/jinzhu/gorm"

type Person struct {
  gorm.Model
  First_Name string `json: first_name`
  Last_Name string `json: last_name`
}