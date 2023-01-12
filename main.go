package main

import (
	"fmt"
	"go-mvc-mvp/controllers"
	"go-mvc-mvp/templates"
	"go-mvc-mvp/views"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func main() {
	/* usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "users/signup.gohtml", "tailwind.gohtml")) */

	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000...")
	CSRF := csrf.Protect([]byte("gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"), csrf.Secure(false))
	http.ListenAndServe("127.0.0.1:3000", CSRF(r))
}
