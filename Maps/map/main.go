package main

import "fmt"


func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, fmt.Errorf("invalid sizes")
	}
	userMap := make(map[string]user)
	for i, name := range names {
		userMap[name] = user{name: name, phoneNumber: phoneNumbers[i]}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}




