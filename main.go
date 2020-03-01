package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os/exec"
	"strings"
)

type BuildResponse struct {
	SCM             string
	BuildCommand    string
	BuildParameters []string
	EmailAddresses  []string
	SCMLocation     string
	CustomCommands  []string
}

type smtpServer struct {
	host string
	port string
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	var buildResponse BuildResponse

	jsonError := json.NewDecoder(r.Body).Decode(&buildResponse)
	if jsonError != nil {
		http.Error(w, jsonError.Error(), http.StatusBadRequest)
		return
	}

	from := email
	password := password

	to := []string{
		emailResponse.Recipient,
	}

	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	emailMessage := []byte("Subject: " + emailResponse.Subject + "\r\n" + emailResponse.Message)
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	err := smtp.SendMail(smtpServer.host+":"+smtpServer.port, auth, from, to, emailMessage)

	if err != nil {
		fmt.Println(err)
		return
	}

	message := r.URL.Path

	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message + " and name: " + emailResponse.Recipient
	//w.Write([]byte(message))
	fmt.Println(message)
}

func main() {

	//fmt.Printf("Hello")
	//args := []string{"/C", "git", "config", "user.email"}

	//argString := "/C git config user.email"
	argString := "cd"

	out, err := exec.Command("cmd", argString).Output()
	if err != nil {
		fmt.Printf("Error has Occured")
		fmt.Printf("%s", err)
	}

	fmt.Printf("Output: %s", out)

}
