package main

import (
	"fmt"
	"time"
  
	"github.com/faizdamar1/go-mongo-crud/config"
	"github.com/faizdamar1/go-mongo-crud/src/modules/profile/model"
	"github.com/faizdamar1/go-mongo-crud/src/modules/profile/repository"
)

func main() {
	fmt.Println("Go Mongo Db")
  
	db, err := config.GetMongoDB()
  
	if err != nil {
	  fmt.Println(err)
	}
  
	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")
	
	//just run MONGO_HOST=localhost MONGO_DB_NAME=tutorial1  go run main.go
	saveProfile(profileRepository)
	//updateProfile(profileRepository)
	//deleteProfile(profileRepository)
	//getProfile("U2", profileRepository)
	// getProfiles(profileRepository)
  }
  
  func saveProfile(profileRepository repository.ProfileRepository)  {
	var p model.Profile
	p.ID = "U3"
	p.FirstName = "Robertttt"
	p.LastName = "Griezmer"
	p.Email = "robert@gmail.com"
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
  
  func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Wuriyanto"
	p.LastName = "Musobar"
	p.Email = "wuriyanto_musobar@gmail.com"
	p.Password = "12345678"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
  
	err := profileRepository.Update("U1", &p)
  
	if err != nil {
	  fmt.Println(err)
	} else {
	  fmt.Println("Profile updated..")
	}
  }
  
  func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")
  
	if err != nil {
	  fmt.Println(err)
	} else {
	  fmt.Println("Profile deleted..")
	}
  }
  
//   func getProfile(id string, profileRepository repository.ProfileRepository) {
// 	profile, err := profileRepository.FindByID(id)
  
// 	if err != nil {
// 	  fmt.Println(err)
// 	}
  
// 	fmt.Println(profile.ID)
// 	fmt.Println(profile.FirstName)
// 	fmt.Println(profile.LastName)
// 	fmt.Println(profile.Email)
//   }
  
  func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()
  
	if err != nil {
	  fmt.Println(err)
	}
  
	for _, profile := range profiles{
	  fmt.Println("-----------------------")
	  fmt.Println(profile.ID)
	  fmt.Println(profile.FirstName)
	  fmt.Println(profile.LastName)
	  fmt.Println(profile.Email)
	}
  }
