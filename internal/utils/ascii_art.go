package utils

import (
	"embed"
	"io/fs"
	"strings"
	"sync"
)

//go:embed ascii/*
var asciiFS embed.FS

var asciiCache = map[string]string{}
var cacheOnce sync.Once

func FetchAscii(id string, config Config) string {
	asciiMode := strings.ToLower(config.Ascii)
	id = strings.ToLower(id)
	loadAsciiFiles()

	if asciiMode == "none" || asciiMode == "disabled" {
		return loadAscii("none.txt")
	}

	if asciiMode != "default" && asciiMode != "" {
		if file, ok := asciiCache[asciiMode]; ok {
			return loadAscii(file)
		}
		return loadAscii("none.txt")
	}

	if file, ok := asciiCache[id]; ok {
		return loadAscii(file)
	}
	return loadAscii("none.txt")
}

func loadAsciiFiles() {
	cacheOnce.Do(func() {
		_ = fs.WalkDir(asciiFS, "ascii", func(path string, asciiDir fs.DirEntry, err error) error {
			if err != nil || asciiDir.IsDir() || !strings.HasSuffix(asciiDir.Name(), ".txt") {
				return nil
			}
			asciiCache[strings.ToLower(strings.TrimSuffix(asciiDir.Name(), ".txt"))] = asciiDir.Name()
			return nil
		})
	})
}

func loadAscii(file string) string {
	data, err := asciiFS.ReadFile("ascii/" + file)
	if err != nil {
		return ""
	}
	return string(data)
}
