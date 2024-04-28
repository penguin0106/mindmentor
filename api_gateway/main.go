package main

import (
	"fmt"
	"io"
	"log"
	"mindmentor/api_gateway/middleware"
	"net/http"
)

const url = "http://localhost"

func handleAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.AuthMiddleware(http.HandlerFunc(next)).ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/create-db", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Call the Auth microservice here
		dbURL := "url:8082/create-db"
		resp, err := http.Get(dbURL)
		if err != nil {
			http.Error(w, "Failed to call wallet service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/migrate", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Call the Auth microservice here
		dbURL := "url:8082/migrate"
		resp, err := http.Get(dbURL)
		if err != nil {
			http.Error(w, "Failed to call wallet service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/service-orders", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		tradeURL := "url:8083/orders"
		resp, err := http.Get(tradeURL)
		if err != nil {
			http.Error(w, "Failed to call trade service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/service-place", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		tradeURL := "url:8083/place-order"
		resp, err := http.Get(tradeURL)
		if err != nil {
			http.Error(w, "Failed to call trade service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/service-profile", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Call the User microservice here
		userURL := "url:8084/profile"
		resp, err := http.Get(userURL)
		if err != nil {
			http.Error(w, "Failed to call trade service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/service-update", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Call the User microservice here
		userURL := "url:8084/profile/update"
		resp, err := http.Get(userURL)
		if err != nil {
			http.Error(w, "Failed to call trade service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/service-wallet", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Вызов микросервис Wallet
		walletURL := "url:8085/wallet"
		resp, err := http.Get(walletURL)
		if err != nil {
			http.Error(w, "Failed to call wallet service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/service-transactions", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Вызов микросервис Wallet
		walletURL := "url:8085/wallet/transactions"
		resp, err := http.Get(walletURL)
		if err != nil {
			http.Error(w, "Failed to call wallet service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/auth", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Call the Auth microservice here
		loginURL := "url:8085/login"
		resp, err := http.Get(loginURL)
		if err != nil {
			http.Error(w, "Failed to call wallet service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))
	http.HandleFunc("/auth-register", handleAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		// Call the Auth microservice here
		registerURL := "url:8085/register"
		resp, err := http.Get(registerURL)
		if err != nil {
			http.Error(w, "Failed to call wallet service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(body))
	}))

	http.HandleFunc("/gateway", func(w http.ResponseWriter, r *http.Request) {
		// Call the Gateway microservice here
	})
	fmt.Println("MainAPI service is running on port 8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
