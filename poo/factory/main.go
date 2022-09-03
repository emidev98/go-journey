package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSender() string
	GetChannel() string
}

type SMSNotification struct{}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification thru SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSender() string {
	return "SMS"
}

func (SMSNotificationSender) GetChannel() string {
	return "WebSocket"
}

type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification thru Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (EmailNotificationSender) GetSender() string {
	return "Email"
}

func (EmailNotificationSender) GetChannel() string {
	return "gRPC"
}

func getNotificationsFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Notification with type '%s' not yet implemented", notificationType)
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func displaySender(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSender())
}

func displayChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetChannel())
}
func main() {
	smsFactory, _ := getNotificationsFactory("SMS")
	emailFactory, _ := getNotificationsFactory("Email")

	sendNotification(smsFactory)
	displaySender(smsFactory)
	displayChannel(smsFactory)

	sendNotification(emailFactory)
	displaySender(emailFactory)
	displayChannel(emailFactory)
}
