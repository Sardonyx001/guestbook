//go:build !prod
// +build !prod

package main

import (
	"io/fs"
	"guestbook/internal/embedded"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
