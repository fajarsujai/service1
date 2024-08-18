package main

import (
	"log"
	"net/http"
	// "os"
	"time"

	// "github.com/joho/godotenv"
)

// loggingMiddleware mencatat setiap permintaan dan respons
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom ResponseWriter to capture the status code
		lrw := &loggingResponseWriter{w, http.StatusOK}

		// Call the next handler
		start := time.Now()
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)

		// Log status and other details
		log.Printf("Status: %d, Method: %s, URL: %s, Duration: %v", lrw.statusCode, r.Method, r.URL.Path, duration)
	})
}

// loggingResponseWriter adalah custom ResponseWriter untuk menangkap status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}



func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	// log.Fatal("Error loading .env file")
	// 	log.Println("Error occurred:", err)

	// }
	
	// err := doSomething()
    // if err != nil {
		//     log.Println("Error occurred:", err)
    // }
	// godotenv package
	// appenv := os.Getenv("APPENV")
	// apport := os.Getenv("PORT")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Update 1 dan Hello ")
		log.Println("Update 1 dan Hello ")
	})
	
	// Periksa kode status HTTP
	// Handler untuk rute yang ada
	// mux := http.NewServeMux()
	// mux.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.URL.Path == "/service1" {
	// 		w.Write([]byte("service 1 "))
	// 	} else {
	// 		http.NotFound(w, r)
	// 	}
	// })

	// Tambahkan middleware logging
	// loggedMux := loggingMiddleware(mux)

	// Mulai server
	fmt.Printf("Server running (port=3001), route: http://localhost:3001")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
	
	// log.Println("Server listening on port 3001")
	// if err := http.ListenAndServe(":3001", loggedMux); err != nil {
	// 	log.Fatalf("Server failed: %v", err)
	// }

}

