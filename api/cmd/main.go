package main

import (
	"flag"
	"fmt"
	"log"
  "os"

  "github.com/joho/godotenv"
  api "api/pkg/api"
  storage "api/pkg/storage"
  types "api/pkg/types"
)

func seedUser(store storage.Storage, username, email, pw string) *types.User {
	user, err := types.NewUser(email, username, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateUser(user); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new user => ", user.ID)

	return user
}

func seedUsers(s storage.Storage) {
	seedUser(s, "bob", "bob@gmail.com", "hunter88888")
	seedUser(s, "tom", "tom@gmail.com", "password")
}

func seedSoftware(store storage.Storage, name, title, description, image, url, username string) *types.Software {
	software, err := types.NewSoftware(name, title, description, image, url, username)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateSoftware(software); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new software => ", software.ID)

	return software
}

func seedSoftwares(s storage.Storage) {
  seedSoftware(s, "brave", "Brave Browser", "Secure, fast and private web browser with adblocker.", "brave-logo.svg", "https://www.brave.com", "tom")
  seedSoftware(s, "session", "Session", "End-to-end encrypted messenger that minimises sensitive metadata, designed and built for people who want absolute privacy and freedom from any form of surveillance.", "Session_Logo.svg", "https://www.getsession.org", "tom")
  seedSoftware(s, "telegram", "Telegram Messenger", "Cloud-based, cross-platform, encrypted instant messaging (IM) service","Telegram_logo.svg", "https://www.telegram.org", "tom")
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	store.Init()

	if *seed {
		fmt.Println("seeding the database")
		seedUsers(store)
		seedSoftwares(store)
	}

  port := os.Getenv("API_PORT")
  if port == "" {
    port = "3000"
  }
  server := api.NewAPIServer(":" + port, store)
	server.Run()
}
