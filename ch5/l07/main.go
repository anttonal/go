package main

type contact struct {
	userID       string
	age          int32
	sendingLimit int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}
