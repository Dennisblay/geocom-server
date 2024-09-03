package api

import (
	"github.com/gin-gonic/gin"
)

// UssdController handles USSD requests.
func UssdController(c *gin.Context) {
	// Read the variables sent via POST from our API
	//phoneNumber := c.PostForm("phoneNumber")
	text := c.PostForm("text")

	var response string

	switch text {
	case "":
		// This is the first request. Start the response with "CON"
		response = "CON Welcome to GeoComp\n1. Buy Credits\n2. Exit"
	case "1":
		// User chose to buy credits, show available plans
		response = "CON Select Plan\n1. 10 credits ==> GH₵1\n2. 50 credits ==> GH₵5\n3. 100 credits ==> GH₵10"
	case "2":
		// User chose to exit, end the session
		response = "END Thank you for using GeoComp. Goodbye!"
	case "1*1":
		// User selected the 10 credits plan
		response = "END You have selected the 10 credits plan for GH₵1.\n You will receive a mobile money prompt to proceed with the transaction"
	case "1*2":
		// User selected the 50 credits plan
		response = "END You have selected the 50 credits plan for GH₵1.\n You will receive a mobile money prompt to proceed with the transaction"
	case "1*3":
		// User selected the 100 credits plan
		response = "END You have selected the 100 credits plan for GH₵1.\n You will receive a mobile money prompt to proceed with the transaction"
	default:
		// Handle unexpected input
		response = "END Invalid input. Please try again."
	}

	// Send the response back to the API
	c.String(200, response)
}
