package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var db *sql.DB
var err error
// var stm *sql.Stmt

func main() {

	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/field")

	if err != nil {
		log.Fatalf("Failed to connect with database %v", err)

	}
	// Ping is built in methode use to check if the connectivity to database.
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect via Ping function %v", err)

	}
	//    using gorilla  Mux package to create new router and url matcher for building  web application 
	r := mux.NewRouter()

	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/", register).Methods("POST")

	//cors creation

	CorsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"PUT","GET","POST","DELETE"}),
		handlers.AllowedHeaders([]string{"content-type", "Authorization"}),
	)
	fmt.Println("listening at  Port :8080....")
	log.Fatal(http.ListenAndServe(":8080", CorsMiddleware(r)))

}

func register(w http.ResponseWriter, r *http.Request) {

	var user User
	var hashedPassword []byte

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return

	}
	// stm,err = db.Prepare("INSERT INTO register (username, password) value(?,?)")
	// if err != nil {
	// 	http.Error(w, "Failed to Prepare statement", http.StatusInternalServerError)
	// 	return
	// }
	// defer stm.Close()
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to Encypt password", err)
	}
	_, err = db.Exec("INSERT INTO register (username, password) value(?,?)", user.Username, hashedPassword)

	if err != nil {
		http.Error(w, "failed--to--insert data to the table", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "successfuly add new user "})

}
func Login(w http.ResponseWriter, r *http.Request) {
    
	var user User
    var storedPassword string

    // Decode JSON payload
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        log.Println("Invalid Request Body:", err)
        http.Error(w, "Invalid Request", http.StatusBadRequest)
        return
    
	}
	
    log.Println("Received Login Request for Username:", user.Username)
    log.Println("Received Login Request for Username:", user.Password)
   
   

    // Query database for user
    Rows := db.QueryRow("SELECT password FROM register WHERE username = ?", user.Username)
    err = Rows.Scan(&storedPassword)
    if err == sql.ErrNoRows {
        log.Println("No User Found for Username:", user.Username)
        http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
        return
    } else if err != nil {
        log.Println("Database Query Error:", err)
        http.Error(w, "Failed to select data", http.StatusInternalServerError)
        return
    }

    log.Println("Stored Password Hash:", storedPassword)

    // Compare hashed password
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
    if err != nil {
        log.Println("Password Comparison Failed:", err)
        http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
        return
    }

    log.Println("Login Successful for Username:", user.Username)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Successfully Login"})

}

