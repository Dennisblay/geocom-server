package api

import (
	"github.com/gin-gonic/gin"
)

// UssdController handles USSD requests and simulates MTN Mobile Money payment.
func UssdController(c *gin.Context) {
	//phoneNumber := c.PostForm("phoneNumber")
	text := c.PostForm("text")

	var response string

	switch text {
	case "":
		// Initial options for the user
		response = "CON Welcome to GeoComp\n1. Buy Credits\n2. Exit"
	case "1":
		// User selects to buy credits
		response = "CON Select Plan\n1. 10 credits ==> GH₵1\n2. 50 credits ==> GH₵5\n3. 100 credits ==> GH₵10"
	case "1*1", "1*2", "1*3":
		// After selecting a plan, prompt MTN Mobile Money options
		amount := ""
		switch text {
		case "1*1":
			amount = "GH₵1"
		case "1*2":
			amount = "GH₵5"
		case "1*3":
			amount = "GH₵10"
		}
		response = "CON You are about to pay " + amount + " using MTN Mobile Money.\n1. Confirm\n2. Cancel"
	case "1*1*1", "1*2*1", "1*3*1":
		// Simulate MTN Mobile Money payment confirmation
		response = "CON Enter your Mobile Money PIN to confirm payment"
	case "1*1*1*<PIN>", "1*2*1*<PIN>", "1*3*1*<PIN>":
		// Simulate successful payment after PIN entry
		response = "END Payment successful! Thank you for your purchase."
	case "1*1*2", "1*2*2", "1*3*2":
		// User cancels the payment
		response = "END You have canceled the payment."
	case "2":
		// User exits
		response = "END Thank you for using GeoComp. Goodbye!"
	default:
		// Handle invalid input
		response = "END Invalid input. Please try again."
	}

	c.String(200, response)
}
