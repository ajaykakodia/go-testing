//go:build integration
// +build integration

package service_test

import (
	"fmt"
	"testing"

	"github.com/ajaykakodia/go-testing/service"
	"github.com/stretchr/testify/mock"
)

type SMSServiceMock struct {
	mock.Mock
}

func (sms *SMSServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge Notification")
	args := sms.Called(value)

	return args.Bool(0)
}

func TestChangeCustomer(t *testing.T) {
	smsService := new(SMSServiceMock)
	smsService.On("SendChargeNotification", 100).Return(true)

	myService := service.MyService{
		MessageSevice: smsService,
	}

	myService.ChargeCustomer(100)

	smsService.AssertExpectations(t)
}
