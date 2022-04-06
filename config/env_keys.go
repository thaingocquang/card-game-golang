package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

// env ...
var env ENV

// projectDirName ...
const projectDirName = "card-game-golang"

// InitDotEnv ...
func InitDotEnv() {
	// load env
	if err := godotenv.Load(GetEnvPath()); err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("abcxyz")

	// database ...
	database := Database{Uri: GetEnvString("DB_URI"), Name: GetEnvString("DB_Name"), TestName: GetEnvString("DB_Name_Test")}

	// appPort ...
	appPort := GetEnvString("APP_PORT")

	// Jwt
	jwt := Jwt{SecretKey: GetEnvString("SECRET_KEY")}

	env = ENV{
		Database: database,
		AppPort:  appPort,
		Jwt:      jwt,
	}

}

// GetEnvPath ...
func GetEnvPath() string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)

	fmt.Println("hello", projectName)

	curWorkDir, _ := os.Getwd()
	rootPath := projectName.Find([]byte(curWorkDir))
	return string(rootPath) + `/.env`
}

// GetEnvString ...
func GetEnvString(key string) string {
	return os.Getenv(key)
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
