package main

import (
	"log"
	"net/http"
	"os"

	"admin-panel/treehandler"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := mux.NewRouter()

	// 👇 Redirect root path to dashboard
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
	})

	// Admin routes
	r.HandleFunc("/admin/dashboard", treehandler.AdminDashboardHandler)
	r.HandleFunc("/admin/edit/{id}", treehandler.EditTreeHandler).Methods("GET", "POST")
	r.HandleFunc("/admin/delete/{id}", treehandler.DeleteTreeHandler)

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("✅ Server running at: http://localhost:" + port + "/admin/dashboard")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
