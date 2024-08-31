package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) ussdController(c *gin.Context) {
	sessionID := c.PostForm("sessionId")
	phoneNumber := c.PostForm("phoneNumber")
	networkCode := c.PostForm("networkCode")
	serviceCode := c.PostForm("serviceCode")
	text := c.PostForm("text")

	fmt.Printf("Received USSD request - SessionID: %s, PhoneNumber: %s, NetworkCode: %s, ServiceCode: %s, Text: %s\n",
		sessionID, phoneNumber, networkCode, serviceCode, text)

	var response string

	// Simulate a simple USSD menu
	switch text {
	case "":
		response = "CON Welcome to our service.\n1. Check Balance\n2. Buy Airtime"
	case "1":
		response = "END Your balance is $10"
	case "2":
		response = "CON Enter amount to buy airtime:"
	default:
		response = "END Invalid option"
	}

	c.String(http.StatusOK, response)
}
