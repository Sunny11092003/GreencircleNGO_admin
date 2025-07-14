package main

import (
	"log"
	"net/http"

	"admin-panel/treehandler"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Admin Dashboard route
	r.HandleFunc("/admin/dashboard", treehandler.AdminDashboardHandler)

	r.HandleFunc("/admin/edit/{id}", treehandler.EditTreeHandler)
	r.HandleFunc("/admin/delete/{id}", treehandler.DeleteTreeHandler)

	// main.go  (add just after your dashboard route)
	r.HandleFunc("/admin/edit/{id}", treehandler.EditTreeHandler).Methods("GET", "POST")

	// Serve static files (CSS, JS, images)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	log.Println("✅ Server running at: http://localhost:8081/admin/dashboard")
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
