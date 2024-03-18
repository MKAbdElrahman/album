package handlers

import (
	"album/components"
	"net/http"

	"github.com/a-h/templ"
)

type HTML struct {
}

func (h HTML) GetHomePage(w http.ResponseWriter, r *http.Request) {
	h.renderComponent(w, r, components.Home())
}

func (h HTML) GetContactPage(w http.ResponseWriter, r *http.Request) {
	h.renderComponent(w, r, components.Contact())
}


func (h HTML) renderComponent(w http.ResponseWriter, r *http.Request, component templ.Component) {
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
