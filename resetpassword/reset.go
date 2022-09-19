package resetpassword

import (
	//"bytes"

	"encoding/json"
	"log"
	"net/http"
	"strconv"

	dbs "example.com/database"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Users_DETAILS struct {
	User_ID  int    `json:"User_ID"`
	Password string `json:"Password"`
}

// Reset password function resets the user password based on user id.
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	//cors.SetupCORS(&w, r)
	db := dbs.Connect()

	var p Users_DETAILS //declare a variable p for type Users_DETAILS
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To convert the password in the encrypted form
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)

	id := strconv.Itoa(p.User_ID)
	row := db.QueryRow("SELECT User_ID from Users where User_ID =" + id)

	var UserId int
	err_scan := row.Scan(&UserId)
	if err_scan != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "User doesn't exist", "Status Code": "400 BadRequest"})
		return
	}
	_, err2 := db.Exec("UPDATE Users SET Password=$2 WHERE User_ID=$1;", p.User_ID, string(hashedPassword))

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": err2, "Status Code": "400 BadRequest"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Password Reset successfully!", "Status Code": "200 OK"})

}
func HandleFunc() {

	http.HandleFunc("/ResetPassword", ResetPassword)
	// http.HandleFunc("/View", list_requestTable)
	//start the server on port 5300
	log.Fatal(http.ListenAndServe(":5000", nil))

}
