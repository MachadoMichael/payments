package schema

import "time"

type Payment struct {
	Id       string
	Date     time.Time
	Method   string
	Status   string
	Sender   Payer
	Reciever Store
	Bank     string
}
