package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro = "pro"
)

func getMessageWithRetriesForPlan(plan string, message [3]string) ([]string, error){
	newMsg := message[:]
	var err error
	if plan == planFree {
		newMsg = message[:2]
	}else if plan != planPro {
		newMsg = nil
		err = errors.New("unsupported plan")
	}

	return newMsg, err
}