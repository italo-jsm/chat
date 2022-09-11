package domain

import "time"

type Message struct{
	Payload string
	SenderId string
	ReceiverId string
	Consumed bool
	ConsumeMoment time.Time
	Timestamp time.Time
}