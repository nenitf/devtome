package devtome

import (
	"os"
	"regexp"
	"strings"
)

type BackupFile struct {
	Name string
}

func ArticlesPersistence(path string, articles []Article) (err error) {
	for _, a := range articles {
		var sb strings.Builder
		sb.WriteString(path)
		sb.WriteString("/")
		sb.WriteString(translateTitle2Filename(a.Title))
		sb.WriteString(".md")
		f, err := os.OpenFile(sb.String(), os.O_CREATE, 0644)
		if err != nil {
			return err
		}
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
