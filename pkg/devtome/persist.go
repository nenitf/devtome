package devtome

import (
	"os"
	"regexp"
	"strings"
	"path/filepath"
)

type BackupFile struct {
	Name string
}

func ArticlesPersistence(path string, articles []Article) (err error) {
	for _, a := range articles {
        filename := translateTitle2Filename(a.Title) + ".md"
		f, err := os.Create(filepath.Join(path, filename))
		if err != nil {
			return err
		}
		f.Close()
	}

	return
}

func translateTitle2Filename(origin string) string {
	s := strings.ToLower(origin)
	var re = regexp.MustCompile(`\s`)
	return re.ReplaceAllString(s, "_")
}
