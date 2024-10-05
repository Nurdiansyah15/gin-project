package seeder

import (
	"gin-project/repository"
	"log"
)

func SeedUsers() {
	users := []struct {
		Username string
		Password string
	}{
		{"admin", "password"},
		{"user1", "password123"},
		{"user2", "mysecret"},
	}

	for _, user := range users {
		_, err := repository.CreateUser(user.Username, user.Password)
		if err != nil {
			log.Printf("Failed to seed user %s: %v", user.Username, err)
		}
	}

	log.Println("User seeding completed.")
}
