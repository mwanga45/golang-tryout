package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/gomail.v2"
)

// Email struct to parse email details from the request
type Email struct {
	To       string `json:"to"`        // Recipient's email address
	Message  string `json:"message"`   // Email message content
	Subject  string `json:"subject"`   // Email subject
	Sender   string `json:"senderEmail"` // Sender's email address
	
}

// SendEmail sends an email using gomail
func SendEmail(to, message, senderEmail, subject string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail) // Set the sender's email address
	m.SetHeader("To", to)            // Set the recipient's email address
	m.SetHeader("Subject", subject)  // Set the email subject
	m.SetBody("text/plain", message) // Set the email body as plain text

	// Configure the SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, "nargvbxkrjzovlcj")
	// Bypass the TLS server certificate verification (not recommended for production)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email and handle errors
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("Failed to send email due to:", err)
		return err
	}
	fmt.Println("Email successfully sent")
	return nil
}

// HandlerEmail handles the /Email endpoint
func HandlerEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request method:", r.Method) 
	// Log the request method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body into the Email struct
	var reqEmail Email
	err := json.NewDecoder(r.Body).Decode(&reqEmail)
	if err != nil {
		fmt.Println("Failed to decode JSON due to:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Decoded request: %+v\n", reqEmail) // Log the decoded request

	// Call SendEmail with the decoded data
	err = SendEmail(reqEmail.To, reqEmail.Message, reqEmail.Sender, reqEmail.Subject)
	if err != nil {
		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email successfully sent"})
}

func main() {
	r := mux.NewRouter()

	// Define the route and attach the handler
	r.HandleFunc("/Email", HandlerEmail).Methods("POST")

	// Add CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Frontend origin
		handlers.AllowedMethods([]string{"PUT", "GET", "POST", "DELETE"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
	)
   
	// Start the server with CORS
	fmt.Println("Server is running on port 8010...")
	log.Fatal(http.ListenAndServe(":8010", corsMiddleware(r)))
}
