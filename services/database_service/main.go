package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	_ "github.com/lib/pq"
)

const (
	defaultHost     = "localhost"
	defaultPort     = "5432"
	defaultUser     = "postgres"
	defaultPassword = "mindmentor"
	defaultDBName   = "mindmentor"
	httpPort        = "8000"
)

var (
	db  *sql.DB
	err error
)

func startService(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Initialize database connection
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", defaultHost, defaultPort, defaultUser, defaultPassword, defaultDBName)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close() // Close the database connection

	http.HandleFunc("/create", CreateDatabaseHandler)
	http.HandleFunc("/migrate", MigrateDatabaseHandler)

	fmt.Println("Database service is running on port", httpPort, "...")
	if err := http.ListenAndServe(":"+httpPort, nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func CreateDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Resolve the path relatively
	scriptPath := filepath.Join("scripts", "create_database.sh")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		http.Error(w, "Script not found", http.StatusInternalServerError)
		return
	}

	cmd := exec.Command("/bin/bash", scriptPath)
	cmd.Dir = "scripts" // Set the working directory relative to the current directory
	cmd.Env = os.Environ()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startService(cmd); err != nil {
			http.Error(w, "Failed to create the database", http.StatusInternalServerError)
			return
		}
	}()

	wg.Wait()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database created successfully")
}

func MigrateDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Use relative path resolution for the migration script
	scriptPath := filepath.Join("scripts", "migrate_database.sh")
	cmd := exec.Command("/bin/bash", scriptPath)
	cmd.Dir = "scripts"
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		http.Error(w, "Failed to migrate the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database migrated successfully")
}
