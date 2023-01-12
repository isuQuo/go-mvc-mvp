package controllers

import (
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// get correct data and supply to template
	/* var data struct {
		...
	}
	data. = r.FormValue("...") */
	//m.View.Execute(w, r, data)
}
