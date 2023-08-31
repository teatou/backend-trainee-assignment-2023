package main

import (
	"backend-trainee-assignment-2023/internal/app"
	"os"
)

const configEnv = "CONFIG"

func main() {
	val, ok := os.LookupEnv(configEnv)
	if !ok {
		panic("no config env")
	}
	if err := app.Avito(val); err != nil {
		panic(err)
	}
}
