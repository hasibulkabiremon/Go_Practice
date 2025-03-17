package main

type contact struct {
	userID       string
	
	sendingLimit int64
	
	age          int64

	
	
}

type perms struct {
	
	canSend         bool
	permissionLevel float64
	canReceive      bool
	canManage       bool
	
	
}
