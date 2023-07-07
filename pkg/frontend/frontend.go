package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

// Embed the frontend build directory
//
//go:embed build/*
var BuildFs embed.FS

// Get the subtree of the embedded files with `build` directory as a root.
func BuildHTTPFS() (http.FileSystem, error) {
	build, err := fs.Sub(BuildFs, "build")
	if err != nil {
		return nil, err
	}
	return http.FS(build), nil
}
