package main

import (
	"math/rand"
	"path/filepath"
	"time"
)

var supportedExtensions map[string]bool

func init() {
	rand.Seed(time.Now().UnixNano())
	supportedExtensions = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".webp": true,
	}
}

func IsValidExtension(name string) bool {
	extension := filepath.Ext(name)
	_, ok := supportedExtensions[extension]
	return ok
}

func RandomName() string {
	b := make([]byte, maxCharacters)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
