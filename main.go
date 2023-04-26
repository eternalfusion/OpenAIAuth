package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"github.com/acheong08/OpenAIAuth/auth"
	"encoding/json"
)

type RequestBody struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func getToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Printf("Handle /token request\n")

        var reqBody RequestBody
        err := json.NewDecoder(r.Body).Decode(&reqBody)

        if err != nil {
            http.Error(w, "Error reading request body", http.StatusBadRequest)
            return
        }

        auth := auth.NewAuthenticator(reqBody.Email, reqBody.Password, os.Getenv("PROXY"))

		beginErr := auth.Begin()
		if beginErr.Error != nil {
			println("Error: " + beginErr.Details)
			println("Location: " + beginErr.Location)
			println("Status code: " + fmt.Sprint(beginErr.StatusCode))
			println("Embedded error: " + beginErr.Error.Error())
			w.WriteHeader(500)
			io.WriteString(w, "Error")
			return
		}
		token, tokenErr := auth.GetAccessToken()
		if tokenErr.Error != nil {
			println("Error: " + tokenErr.Details)
			println("Location: " + tokenErr.Location)
			println("Status code: " + fmt.Sprint(tokenErr.StatusCode))
			println("Embedded error: " + tokenErr.Error.Error())
			w.WriteHeader(500)
			io.WriteString(w, "Error")
			return 
		}

		fmt.Println("token=" + token)

		io.WriteString(w, token)
	} else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }	
}

func main() {
	http.HandleFunc("/token", getToken)

	http.ListenAndServe(":7555", nil)

}
