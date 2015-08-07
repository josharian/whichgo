package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	path := os.Args[1]
	path = strings.TrimSuffix(path, "/bin/go")
	path = filepath.Base(path)
	if path == "go" {
		path = "go-tip"
	}
	path = strings.TrimPrefix(path, "go-")
	path = strings.TrimPrefix(path, "go.")
	fmt.Print("(", path, ")")
}
