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

func TestGetAll(t *testing.T) {
	is := is.New(t)
	mux := http.NewServeMux()

	p := devtome.Article{
		Id:    1,
		Title: "AAQ",
	}

	mux.HandleFunc("/api/articles/me/all", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	})
	svr := httptest.NewServer(mux)
	w := httptest.NewRecorder()
	defer svr.Close()

	t.Run("Deve retornar o resultado correto", func(t *testing.T) {
		expected, err := json.Marshal(p)
		is.NoErr(err)

		c := devtome.NewClient(svr.URL, "AAA")
		res, err := c.GetAll()
		is.NoErr(err)
		is.Equal(w.Code, http.StatusOK)
		is.Equal(strings.TrimSpace(res), string(expected))
	})
}
