package main

import "fmt"

// Sender

type SMSNotificationSender struct{}

type Isender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Notification

type SMSNotification struct{}

type INotificationFactory interface {
	SendNotification()
	GetSender() Isender
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS.")
}

func (SMSNotification) GetSender() Isender {
	return SMSNotificationSender{}
}
