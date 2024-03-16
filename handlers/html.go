package handlers

import (
	"album/components"
	"context"
	"net/http"
)

type HTML struct {
}

func (h HTML) GetHomePage(w http.ResponseWriter, r *http.Request) {
	components.Home("MK").Render(context.TODO(), w)
}
