package api

import (
	"Anastasia/skillfactory/advanced/news-gathering-service/pkg/models"
	"Anastasia/skillfactory/advanced/news-gathering-service/pkg/repo"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	db     repo.Interface
	router *mux.Router
}

func New(db repo.Interface) *API {
	api := API{
		db: db,
	}

	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.router.HandleFunc("/posts", api.createPostHandler).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("posts/{n}", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = api.db.CreatePost(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["n"]
	n, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	posts, err := api.db.Posts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
