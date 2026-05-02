package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {

	zeroFound := func(u user) bool {
		return u.name == "" || u.number == 0
	}

	return !(zeroFound(mToSend.recipient) || zeroFound(mToSend.sender))
}
