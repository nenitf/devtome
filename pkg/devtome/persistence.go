package devtome

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type BackupFile struct {
	Name string
}

func ArticlesPersistence(path string, articles []Article) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "unable to create path")
	}

	for _, a := range articles {
		filename := translateTitle2Filename(a.Title) + ".md"
		f, err := os.Create(filepath.Join(path, filename))
		defer f.Close()
		if err != nil {
			return errors.Wrap(err, "unable to create file")
		}

		_, err = f.WriteString(a.Content)
		if err != nil {
			return errors.Wrap(err, "unable to write file")
		}
	}

	return
}

func translateTitle2Filename(origin string) string {
	s := strings.ToLower(origin)
	var re = regexp.MustCompile(`\s`)
	return re.ReplaceAllString(s, "_")
}
