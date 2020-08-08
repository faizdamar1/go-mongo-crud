package repository

import (
	"github.com/belajar/go-mongo-crud/src/modules/profile/model"
)

// Profile ProfileRepository
type ProfileRepository interface{
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindById(string)(*model.Profile,error)
	FindAll()(model.Profiles, error)
}
