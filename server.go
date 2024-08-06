package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/stripe/stripe-go/v79"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on localhost:4242...")
	var err = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func handleHealth(writer http.ResponseWriter, request *http.Request) {
	response := []byte("Server is up and running!")

	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(calculateOrderAmount(req.ProductId)),

	}
}
func calculateOrderAmount(productID string)int64{
	switch  productID {
	case "Forever pants":
		return 26000
	case  "Forever shirt":
		return 15500
	case  "Forever hood":
		return 30000

}
return 0
}