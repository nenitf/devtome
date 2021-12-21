package devtome

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type BackupFile struct {
	Name string
}

func FilePersist(path string, articles []Article) (err error) {
	var sb strings.Builder

	for _, a := range articles {
		sb.WriteString(path)
		sb.WriteString("/")
		sb.WriteString(translateTitle2Filename(a.Title))
		sb.WriteString(".md")
		log.Fatal(sb.String())
		f, err := os.OpenFile(sb.String(), os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		// _, err = f.Write([]byte("Hello"))
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
