package main

import (
	"fmt"
)

// Interfaces

type INotificationFactory interface {
	SendNotification()
	GetSender() Isender
}

type Isender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS Sender

type SMSNotificationSender struct{}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// SMS Notification

type SMSNotification struct{}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS.")
}

func (SMSNotification) GetSender() Isender {
	return SMSNotificationSender{}
}

// EMAIL Notification

type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email.")
}

func (EmailNotification) GetSender() Isender {
	return EmailNotificationSender{}
}

// EMAIL Sender

type EmailNotificationSender struct{}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

// Interfaces Methods

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "EMAIL":
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf(`"%s" is not a valid notification type.`, notificationType)
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
