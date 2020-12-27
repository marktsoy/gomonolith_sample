package apiserver

import (
	"net/http"

	"github.com/marktsoy/gomonolith_sample/internal/app/store"
	"github.com/marktsoy/gomonolith_sample/internal/app/utils"
)

type Server struct {
	msgs  chan *utils.Action
	mux   *http.ServeMux
	store store.Store
}

func New(messageQueue chan *utils.Action, store store.Store) *Server {
	serv := &Server{
		msgs:  messageQueue,
		mux:   http.NewServeMux(),
		store: store,
	}

	serv.configureRouter()
	return serv
}

func (s *Server) configureRouter() {
	s.mux.HandleFunc("/", s.createBundle())
	s.mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PONG"))
		w.WriteHeader(200)
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
