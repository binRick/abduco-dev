package main

import (
	"strings"

	"github.com/google/uuid"
)

func SessionNameString() string {
	if len(session_name) == 0 {
		return strings.Split(uuid.NewString(), `-`)[0]
	}
	return session_name
}

func NewNameString() string {
	return uuid.NewString()
}
