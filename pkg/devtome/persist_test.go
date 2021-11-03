package devtome_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/nenitf/devtome/pkg/devtome"
)

func TestFilePersist(t *testing.T) {
	tmpdir := t.TempDir()

	p := []devtome.Article{
		devtome.Article{
			Id:    1,
			Title: "Exemplo De Post",
		},
	}

	// log.Fatal(tmpdir)
	devtome.FilePersist(tmpdir, p)

	// tmpdir will be cleaned up
}

func TestFilename(t *testing.T) {
	is := is.New(t)
	o := "Título original do post"

	fn := devtome.TranslateTitle2Filename(o)

	is.Equal(fn, "título_original_do_post")
}
