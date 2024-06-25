package config

import (
	"affiliate/internal/constants"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var env ENV

// Header keys
const (
	HeaderOrigin         = "Origin"
	HeaderContentLength  = "Content-Length"
	HeaderContentType    = "Content-Type"
	HeaderAuthorization  = "Authorization"
	HeaderAcceptLanguage = "Accept-Language"
	ResponseType         = "responseType"
	HeaderApiKey         = "Api-Key"
)

// Init ...
func Init() {
	env = ENV{
		Env: os.Getenv("ENV"),
	}
	var (
		ctx = context.Background()
	)
	envFile := ".env"
	if err := godotenv.Load(envFile); err != nil {
		fmt.Println("Load env file err: ", err)
	}
	if err := envconfig.Process(ctx, &env); err != nil {
		log.Fatal("Assign env err: ", err)
	}
}

// GetENV ...
func GetENV() *ENV {
	return &env
}

// IsEnvDevelop ...
func IsEnvDevelop() bool {
	return env.Env == constants.EnvDevelop
}

func CheckAuthen(pass string) bool {
	return pass == env.PASS_AUTHEN
}

// IsEnvStaging ...
func IsEnvStaging() bool {
	return env.Env == constants.EnvStaging
}

// IsEnvProduction ...
func IsEnvProduction() bool {
	return env.Env == constants.EnvProduction
}
