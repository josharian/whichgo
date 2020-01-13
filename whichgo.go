package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	path, err := exec.LookPath("go")
	if err != nil {
		return
	}
	path = strings.TrimSuffix(path, "/bin/go")
	path = filepath.Base(path)
	if path == "go" {
		path = "go-tip"
	}
	path = strings.TrimPrefix(path, "go-")
	path = strings.TrimPrefix(path, "go.")
	path = strings.TrimPrefix(path, "go")
	fmt.Print("(", path, ")")
}
