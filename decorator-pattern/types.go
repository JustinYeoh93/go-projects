package main

import "fmt"

type notificationResponse struct {
	err           error
	destinationId string
}

type notification interface {
	send() notificationResponse
}

type smsNotification struct {
	destinationId string
	msg           string
}

func (sn smsNotification) send() notificationResponse {
	fmt.Println("Sending sms notification")
	fmt.Println(sn.msg)

	if len(sn.destinationId) > 0 {
		fmt.Printf("delivered to %v\n", sn.destinationId)
	}

	return simulateResponse(sn.destinationId)
}

type slackNotification struct {
	notification notification
}

func (s slackNotification) send() notificationResponse {
	resp := s.notification.send()

	if len(resp.destinationId) > 0 {
		slackUser := "slack-user-test"
		fmt.Printf("sending to slack as %s", slackUser)
	}

	return simulateResponse(resp.destinationId)
}
