package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func compare(paths []string) {
	m, err := os.Open(paths[0])
	if err != nil {
		return
	}

	offset := 0
	mr := bufio.NewReader(m)

	var o []*bufio.Reader
	for _, n := range paths[1:] {
		f, err := os.Open(n)
		if err != nil {
			return
		}

		o = append(o, bufio.NewReader(f))
	}

	c, err := mr.ReadByte()
	if err != nil {
		return
	}

	for _, t := range o {
		offset = offset + 1
		cp, err := t.ReadByte()
		if err != nil {
			return
		}

		if cp != c {
			fmt.Println("differs at:", offset)

		}
	}
	fmt.Println("finished compare at offset:", offset)

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
