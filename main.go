package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"github.com/acheong08/OpenAIAuth/auth"
)

<<<<<<< HEAD

func getToken(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /token request\n")

	auth := auth.NewAuthenticator(os.Getenv("OPENAI_EMAIL"), os.Getenv("OPENAI_PASSWORD"), os.Getenv("OPENAI_PUID"), os.Getenv("PROXY"))
=======
func main() {
	auth := auth.NewAuthenticator(os.Getenv("OPENAI_EMAIL"), os.Getenv("OPENAI_PUID"), os.Getenv("PROXY"))
>>>>>>> upstream/main
	err := auth.Begin()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		io.WriteString(w, "Error")
		return
	}
	token, err := auth.GetAccessToken()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		io.WriteString(w, "Error")
		return 
	}

	fmt.Println("token=" + token)

	io.WriteString(w, token)
}

func main() {
	http.HandleFunc("/token", getToken)

	http.ListenAndServe(":7555", nil)

}
