package main

import (
	"embed"
	"log"

	"eboek/internal/cloud"
	"eboek/internal/storage"
	"eboek/internal/sync"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// logs any error that might occur.
func main() {
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	store := storage.NewStore("~/.local/share/eboek/eboek.db")
	client := cloud.NewClient(store)
	engine := sync.NewEngine(client, store)

	app := application.New(application.Options{
		Name:        "eboek",
		Description: "eReader made easy",
		Services: []application.Service{
			application.NewService(engine),
			application.NewService(store),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "eboek",
		URL:   "/",
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
