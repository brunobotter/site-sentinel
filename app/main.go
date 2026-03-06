package main

import (
	"github.com/brunobotter/site-sentinel/main/app"
	"github.com/brunobotter/site-sentinel/main/providers"
)

func main() {
	app.NewApplication(providers.List()).Bootstrap()
}
