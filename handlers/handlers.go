package handlers

import (
	"github.com/fouched/celeritas"
	"myapp/data"
	"net/http"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	//err := h.App.Render.Page(w, r, "home", nil, nil)
	// using convenience func
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
