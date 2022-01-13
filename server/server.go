package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type eventRequest struct {
    Message struct {
        Attributes map[string]string
        Data       []byte
        ID         string `json:"message_id"`
    }
    Subscription string
}

type responseData struct {
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Otp string `json:"otp"`
}

func main() {
	pwd, err := os.Getwd()
    if err != nil {
        os.Exit(1)
    }
	envFilePath := pwd+"/"+"test.env"
	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

    http.HandleFunc("/send_otp", SendOtp)
    fmt.Println("Server started at port 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func SendOtp(w http.ResponseWriter, r *http.Request) {
	msg := &eventRequest{}
    if err := json.NewDecoder(r.Body).Decode(msg); err != nil {
        http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
        return
    }
	data := &responseData{} 
    if e := json.Unmarshal(msg.Message.Data, data); e != nil {
        log.Printf("Eror while parsing repomse")
    } else {
        log.Printf("Name: %s \nPhone Number: %s \nOTP: %s\n", data.Name, data.PhoneNumber, data.Otp)
    }


	client := twilio.NewRestClient()

    params := &openapi.CreateMessageParams{}
	toPhoneNumber := "+91"+data.PhoneNumber
    params.SetTo(toPhoneNumber)
    params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	messageBody := fmt.Sprintf("Hello %s!\nYour OTP is %s", data.Name, data.Otp)
    params.SetBody(messageBody)

    _, err := client.ApiV2010.CreateMessage(params)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("SMS sent successfully!")
    }
}