package main

import (
	"github.com/joho/godotenv"
	"github.com/simple-automation-with-golang-and-cucumber/features/helper"
)

func init() {
	env := godotenv.Load()
	helper.LogPanicln(env)
}

func main() {
	// can be blank for now
}
