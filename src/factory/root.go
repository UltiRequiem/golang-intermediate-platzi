package main

import "fmt"

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
