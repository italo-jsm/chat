package domain

import "time"

type Message struct{
	Payload string
	SenderId string
	ReceiverId string
	Timestamp time.Time
}