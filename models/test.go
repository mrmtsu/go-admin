package models

import "github.com/joho/godotenv"

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
}
