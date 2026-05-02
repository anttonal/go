package main

type User struct {
	Membership
	Name string
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	u := User{
		Membership: Membership{Type: membershipType},
		Name:       name}

	switch u.Type {
	case "premium":
		u.MessageCharLimit = 1000
	default:
		u.MessageCharLimit = 100
	}

	return u
}
