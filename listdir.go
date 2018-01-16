package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nyxnaut/fastwalk"
)

func main() {
	root := "."

	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	// var paths []string
	err := fastwalk.Walk(root, func(path string, typ os.FileMode) error {
		if typ.IsDir() {
			return nil
		}

		relative, _ := filepath.Rel(root, path)
		if strings.Index(relative, ".") == 0 {
			return nil
		}

		fmt.Println(relative)
		// paths = append(paths, relative)
		return nil
	})

	if err != nil {
		fmt.Println("fastwalk.Walk", err, root)
	}

	// for _, p := range paths {
	// 	fmt.Println(p)
	// }
}
