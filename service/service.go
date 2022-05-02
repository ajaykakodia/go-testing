package service

import "fmt"

type MessageService interface {
	SendChargeNotification(int) bool
}

type SMSService struct {
}

type MyService struct {
	MessageSevice MessageService
}

func (sms SMSService) SendChargeNotification(val int) bool {
	fmt.Println("Sending Production charge Notification")

	return true
}

func (ms MyService) ChargeCustomer(value int) error {
	ms.MessageSevice.SendChargeNotification(value)
	fmt.Printf("Charging the customer for value %d\n", value)
	return nil
}
