package frontend

import (
	"embed"
	"io/fs"
)

//go:embed release/*
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
	releaseDir, releaseErr := fs.Sub(staticDir, "release")

	if releaseErr != nil {
		panic(releaseErr)
	}
	return releaseDir
}
