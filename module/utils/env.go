package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Env struct {
	Jwt_Secret_Key []byte
	PORT           string
}

func InitEnv() (*Env, error) {

	e := Env{}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}
	dir := filepath.Dir(filename)
	envPath := filepath.Join(dir, "../../.env")
	if os.Getenv("EnvFile") == "" {
		EnvErr := godotenv.Load(envPath)
		if EnvErr != nil {
			return nil, EnvErr
		}
	}
	//check the variables
	Jwt_Secret_Key := os.Getenv("JWT_SECRET_KEY")
	if Jwt_Secret_Key == "" {
		return nil, errors.New("no JWT_SECRET_KEY was found in env file")
	}
	e.Jwt_Secret_Key = []byte(Jwt_Secret_Key)
	e.PORT = os.Getenv("PORT")
	if e.PORT == "" {
		return nil, errors.New("no PORT was found in env file")
	}
	return &e, nil
}
