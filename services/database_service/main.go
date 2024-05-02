package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/create", CreateDatabaseHandler)
	http.HandleFunc("/migrate", MigrateDatabaseHandler)

	fmt.Println("Database service is running on port 5432...")
	if err := http.ListenAndServe(":5432", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func CreateDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cmd := exec.Command("/bin/bash", "scripts/create_database.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		http.Error(w, "Failed to create database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database created successfully")
}

func MigrateDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cmd := exec.Command("/bin/bash", "scripts/migrate_database.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		http.Error(w, "Failed to migrate database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database migrated successfully")
}
