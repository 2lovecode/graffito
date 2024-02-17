package frontend

import (
	"embed"
	"io/fs"
)

//go:embed prod/*
var staticDir embed.FS

func AssetsDir() fs.FS {
	releaseDir := WebDir()
	assetsDir, err := fs.Sub(releaseDir, "assets")

	if err != nil {
		panic(err)
	}
	return assetsDir
}

func WebDir() fs.FS {
	releaseDir, releaseErr := fs.Sub(staticDir, "prod")

	if releaseErr != nil {
		panic(releaseErr)
	}
	return releaseDir
}
