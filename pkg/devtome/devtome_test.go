package devtome_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/nenitf/devtome/pkg/devtome"
)

func TestAcessoDaAPI(t *testing.T) {
	is := is.New(t)

    setupServer := func() (mux *http.ServeMux, svr *httptest.Server) {
        mux = http.NewServeMux()
        svr = httptest.NewServer(mux)
        return
    }

	t.Run("Retorna artigos conforme a API", func(t *testing.T) {
        mux, svr := setupServer()
        defer svr.Close()
        p := []devtome.Article{
            devtome.Article{
                Id:    1,
                Title: "Exemplo De Post",
            },
        }
        expected, _ := json.Marshal(p)

        mux.HandleFunc("/api/articles/me/all", func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(p)
        })

        c := devtome.NewClient(svr.URL, "token-desnecessário")
        res, err := c.GetAll()
        is.NoErr(err)
        is.Equal(strings.TrimSpace(res), string(expected))
    })
}

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
