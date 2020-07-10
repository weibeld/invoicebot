package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"log"
)

func main() {
	stripe.Key = "sk_test_Jdv3kgamHRXSygTcXK7g5nwn00O4bCvXL9"

	/*params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyCHF)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		ReceiptEmail: stripe.String("jenny.rosen@example.com"),
	}*/
	//paymentintent.New(params)

	cancelParams := &stripe.PaymentIntentCancelParams{
		CancellationReason: stripe.String("abandoned"),
	}
	pi, err := paymentintent.Cancel("pi_1G9Cd0GxzMsXyrgoUAJ7r33z", cancelParams)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s cancelled\n", pi.ID)
	}

	listParams := &stripe.PaymentIntentListParams{}
	i := paymentintent.List(listParams)
	for i.Next() {
		pi := i.PaymentIntent()
		fmt.Printf("%s: %s", pi.ID, pi.Status)
		if pi.Status == "canceled" {
			fmt.Printf(" because of: %s", pi.CancellationReason)
		}
		fmt.Println()
		//spew.Dump(pi)
	}

	/*pi, err := paymentintent.Get("pi_asdfsdaf", nil)
	if err != nil {
		log.Fatal("foo bar")
	} else {
		spew.Dump(pi)
	}*/
}
