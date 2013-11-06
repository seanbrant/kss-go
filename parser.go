package kss

import (
	"os"
	"path/filepath"
)

func Parser(paths ...string) map[string]Section {
	sections := map[string]Section{}

	for i := range paths {
		path := paths[i]

		filepath.Walk(path, func(p string, f os.FileInfo, err error) error {

			if (!f.IsDir()) {
				comments := CommentParser(p)

				for i := range comments {
					comment := comments[i]

					section := NewSection(comment, p)
					sections[section.Reference] = section
				}
			}

			return nil
		})
	}

	return sections
}
