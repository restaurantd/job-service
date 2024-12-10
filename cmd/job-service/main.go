package main

import app2 "job-service/internal/app"

func main() {
	app := app2.Init()
	app.Run()
}
