package models

import "time"

type NewOTP struct {
	Mobile string
	Ttl    time.Duration
}
