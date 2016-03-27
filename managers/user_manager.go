package managers

import (
	"github.com/tuvistavie/gormsample/models"
)

type userManager struct {
	*manager
}

func (u *userManager) CreateUser(user models.User) error {
	return u.db.Create(user).Error
}

func (u *userManager) CreateUserLike(from, to models.User) {
	// whatever
}
