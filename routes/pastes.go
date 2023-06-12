package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"xylight.dev/pastebin/db"
)

type Paste struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func PasteRoute() chi.Router {
	router := chi.NewRouter()
	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		paste := Paste{}
		row := db.DB.QueryRow("SELECT content, title FROM pastes WHERE id = $1", id)
		if err := row.Scan(&paste.Content, &paste.Title); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "paste not found")
			return
		}

		re.JSON(w, 200, paste)
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var newPaste Paste

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			return
		}

		if err := json.Unmarshal(body, &newPaste); err != nil {
			w.WriteHeader(400)
		}

		_, err = db.DB.Exec("INSERT INTO pastes (content, title) VALUES ($1, $2)", newPaste.Content, newPaste.Title)

		if err != nil {
			return
		}

		re.JSON(w, 200, newPaste)
	})

	return router
}
