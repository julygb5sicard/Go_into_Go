package main

import(
"fmt"
"net/http"
)

func main() {
// Defining a request handler
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprint(w, "Hello, World!")
})

// Start the web server on port 8080
fmt.Println("Server is running on http://localhost:8080")
err := http.ListenAndServe(":8080", nil)
if err != nil {
fmt.Println("Error:", err)
}
}
