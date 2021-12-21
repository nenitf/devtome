package devtome_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/nenitf/devtome/pkg/devto"
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
		p := []devto.Article{
			devto.Article{
				Id:    1,
				Title: "Exemplo De Post",
			},
		}

		mux.HandleFunc("/api/articles/me/all", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
		})

		c := devtome.NewClient(svr.URL, "token-desnecessário")
		art, err := c.GetAll()
        arts, _ := json.Marshal(art)
		is.NoErr(err)
		is.Equal(string(`[{"id":1,"title":"Exemplo De Post","description":""}]`), string(arts))
	})
}

func TestPersistenciaDosArtigosEmArquivos(t *testing.T) {
	is := is.New(t)

	t.Run("Persiste multiplos arquivos", func(t *testing.T) {
		tmpdir := t.TempDir()

		p := []devtome.Article{
			devtome.Article{
				Id:    1,
				Title: "Exemplo De Post",
			},
			devtome.Article{
				Id:    1,
				Title: "Mais um Exemplo De Post",
			},
		}

		err := devtome.ArticlesPersistence(tmpdir, p)
		is.NoErr(err)

		_, err = os.Stat(filepath.Join(tmpdir, "exemplo_de_post.md"))
		is.NoErr(err)

		_, err = os.Stat(filepath.Join(tmpdir, "mais_um_exemplo_de_post.md"))
		is.NoErr(err)
	})
}
