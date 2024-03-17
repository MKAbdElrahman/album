package handlers

import (
	"album/components"
	"context"
	"net/http"
)

type HTML struct {
}

func (h HTML) GetHomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	components.Home("MK").Render(context.TODO(), w)
}

func (h HTML) GetContactPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	components.Contact().Render(context.TODO(), w)
}
