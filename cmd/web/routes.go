package main

import (
	"github.com/bworkman1/bookings/pkg/config"
	"github.com/bworkman1/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/general-quarters", handlers.Repo.GeneralQuarters)
	mux.Get("/majors-suite", handlers.Repo.MajorsSuite)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/contact-us", handlers.Repo.ContactUs)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
