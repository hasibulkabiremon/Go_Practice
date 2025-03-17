package main

type Membership struct{
	Type string
	MessageCharLimit int
}

type User struct {
	Membership
	Name string
}

func newUser(name string, membershipType string) User {
	// ?
	new_user := User{}
	new_user.Name = name
	new_user.Type = membershipType
	if membershipType == "premium"{
		new_user.MessageCharLimit = 1000
	}else{
		new_user.MessageCharLimit = 100
	}

	return new_user
}