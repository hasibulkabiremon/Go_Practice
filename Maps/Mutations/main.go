package main

import "fmt"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
	if _, ok := users[name]; !ok {
		return false, fmt.Errorf("not found")
	} else {
		if users[name].scheduledForDeletion {
			delete(users, name)
			return true, nil
		} else {
			return false, nil
		}
	}
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
