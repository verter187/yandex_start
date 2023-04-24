package main

import "fmt"

type MyType int

func (m MyType) String() string {
	return fmt.Sprintf("MyType: %d", m)
}

type DeliveryState string

const (
	DeliveryStatePending   DeliveryState = "pending"      // сообщение отправлено
	DeliveryStateAck       DeliveryState = "acknowledged" // сообщение получено
	DeliveryStateProcessed DeliveryState = "processed"    // сообщение обработано успешно
	DeliveryStateCanceled  DeliveryState = "canceled"     // обработка сообщения прервана

)

func (s DeliveryState) IsValid() bool {
	switch s {
	case DeliveryStatePending, DeliveryStateAck, DeliveryStateProcessed, DeliveryStateCanceled:
		return true
	default:
		return false
	}
}

func (s DeliveryState) String() string {
	return string(s)
}

func HandleMsgDeliveryStatus(status DeliveryState) error {
	if !status.IsValid() {
		return fmt.Errorf("status: invalid")
	}
	return nil
}

func main() {
	var m MyType = 5

	s := m.String()
	fmt.Println(s)

	// приводим строку "fake" к типу DeliveryState
	if err := HandleMsgDeliveryStatus(DeliveryState("fake")); err != nil {
		panic(err)
	}
}
