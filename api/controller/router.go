package controller

import (
	"github.com/gorilla/mux"
)

type routes struct {
	root    *mux.Router
	apiRoot *mux.Router
}

type api struct {
	routes *routes
}

func Init(root *mux.Router) {
	r := routes{
		root:    root,
		apiRoot: root.PathPrefix("/api").Subrouter(),
	}

	_ = api{
		routes: &r,
	}
}

func (api *api) initProject() {
	api.routes.apiRoot.HandleFunc("/user", api.UserSignUp).Methods("POST")
}
