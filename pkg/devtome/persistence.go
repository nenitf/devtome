package devtome

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type BackupFile struct {
	Name string
}

func ArticlesPersistence(path string, articles []Article) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	for _, a := range articles {
		filename := translateTitle2Filename(a.Title) + ".md"
		f, err := os.Create(filepath.Join(path, filename))
		defer f.Close()
		if err != nil {
			return err
		}

		_, err = f.WriteString(a.Content)
		if err != nil {
			return err
		}
	}

	return
}

func translateTitle2Filename(origin string) string {
	s := strings.ToLower(origin)
	var re = regexp.MustCompile(`\s`)
	return re.ReplaceAllString(s, "_")
}
