package main

func main() {
	sms := smsNotification{
		destinationId: "some-destination",
		msg:           "this is a message",
	}

	slack := slackNotification{
		notification: sms,
	}

	slack.send()
}

func simulateResponse(dst string) notificationResponse {
	return notificationResponse{
		nil,
		dst,
	}
}
