package entiy

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	name string
	pwd  string
	age  string
}
