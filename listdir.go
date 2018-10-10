package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func compare(paths []string) {
	var readers []*bufio.Reader
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return
		}
		defer f.Close()

		readers = append(readers, bufio.NewReader(f))
	}

	eof := false
	offset := 0
	for done := false; !done; {
		unique := make(map[byte]struct{}, 0)

		for _, reader := range readers {
			c, err := reader.ReadByte()
			if err == io.EOF {
				eof = true
			} else if err != nil {
				fmt.Println("err:", err)
				continue
			}
			unique[c] = struct{}{}
		}

		if eof || (len(unique) != 1) {
			done = true
			continue
		}

		offset = offset + 1
	}
	if eof {
		fmt.Println("end of file reached at:", offset)

	} else {
		fmt.Println("differs at:", offset)
	}
}

func main() {
	var dirs []string
	if len(os.Args) > 1 {
		dirs = os.Args[1:]
	} else {
		dirs = append(dirs, ".")
	}

	f := make(map[int64][]string)

	for _, root := range dirs {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if info.IsDir() {
				return nil
			}

			relative, _ := filepath.Rel(root, path)
			f[info.Size()] = append(f[info.Size()], relative)

			fmt.Println(relative)

			return nil
		})

		fmt.Println()

		for k, v := range f {
			if len(v) > 1 {
				fmt.Println(k)

				var paths []string
				for _, n := range v {
					fmt.Println(n)
					paths = append(paths, n)
				}

				fmt.Println()

				compare(paths)

				fmt.Println()
			}
		}
	}
}
