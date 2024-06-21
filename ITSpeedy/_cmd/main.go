package main

import (
	"fmt"
	"os"
	"templategoapi"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("template go api")
	if os.Getenv("APIKEY") == "" {
		_ = godotenv.Load()
	}
	templategoapi.StartServer()

}
