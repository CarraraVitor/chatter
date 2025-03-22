package main

import (
	"log"
	"net/http"

	"chatter/app"
	"chatter/app/database"
	"chatter/app/middleware"

	_ "modernc.org/sqlite"
)

func main() {
    database.InitDB()

    router := http.NewServeMux()

    for _, route := range app.AppRouter.Routes {
        router.HandleFunc(route.Path, route.Handler)
    }

    router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	url := "0.0.0.0:8000"
	server := http.Server{
		Addr: url,
		Handler: middleware.CreateStack(
			middleware.Logging,
		)(router),
	}

	log.Printf("[INFO] Listening on %v...\n", url)
	log.Fatalf("[CRITICAL] Server Failed: %s", server.ListenAndServe())
}

