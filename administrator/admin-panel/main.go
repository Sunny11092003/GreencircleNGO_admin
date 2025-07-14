package main

import (
	"log"
	"net/http"
	"os"

	"admin-panel/treehandler"

	"github.com/gorilla/mux"
)

func main() {
	// Use PORT from environment for Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // fallback for local dev
	}

	r := mux.NewRouter()

	// Admin Dashboard route
	r.HandleFunc("/admin/dashboard", treehandler.AdminDashboardHandler)

	// Edit and Delete Tree routes
	r.HandleFunc("/admin/edit/{id}", treehandler.EditTreeHandler).Methods("GET", "POST")
	r.HandleFunc("/admin/delete/{id}", treehandler.DeleteTreeHandler)

	// Serve static files (CSS, JS, images)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	log.Println("✅ Server running at: http://localhost:" + port + "/admin/dashboard")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
