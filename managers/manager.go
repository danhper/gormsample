package managers

import (
	"github.com/jinzhu/gorm"
)

type manager struct {
	db *gorm.DB
}

// UserManager interacts with users
var UserManager *userManager

// Boot setups the managers
func Boot(db *gorm.DB) {
	baseManager := &manager{db: db}
	UserManager = &userManager{manager: baseManager}
}
