package main

import (
	"fmt"
	"time"
	"github.com/belajar/go-mongo-crud/config"
	"github.com/belajar/go-mongo-crud/src/modules/profile/model"
	"github.com/belajar/go-mongo-crud/src/modules/profile/repository"
	
)

func main() {
	fmt.Println("Go Mongo Db")
  
	db, err := config.GetMongoDB()
  
	if err != nil {
	  fmt.Println(err)
	}
  
	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")
  
	saveProfile(profileRepository)

  }
  
  func saveProfile(profileRepository repository.ProfileRepository)  {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Faiz"
	p.LastName = "Damar"
	p.Email = "admin@mail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
  
	err := profileRepository.Save(&p)
  
	if err != nil {
	  fmt.Println(err)
	} else {
	  fmt.Println("Profile saved..")
	}
  }
  
