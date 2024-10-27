package internal

import (
	"time"
)

type User struct {
	userId   int
	username string

	firstName string
	lastName  string

	dateOfBirth time.Time
}
