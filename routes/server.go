package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/unrolled/render"
)

type Server struct {
}

var re *render.Render

func NewServer() *Server {

	re = render.New(render.Options{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/api/pastes", PasteRoute())

	fmt.Println("Mounted routes")

	http.ListenAndServe(":3000", r)

	return &Server{}
}
