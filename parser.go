package kss

import (
	"os"
	"path/filepath"
	"strings"
)

func Parser(paths ...string) map[string]Section {
	sections := map[string]Section{}

	for i := range paths {
		path := paths[i]

		_, err := os.Stat(path)
		if err != nil {
			continue
		}

		filepath.Walk(path, func(p string, f os.FileInfo, err error) error {

			if (!f.IsDir()) {
				comments := CommentParser(p)

				for i := range comments {
					comment := comments[i]
					filename := strings.TrimLeft(strings.Replace(p, path, "", 1), "/")
					section := NewSection(comment, filename)
					sections[section.Reference] = section
				}
			}

			return nil
		})
	}

	return sections
}
