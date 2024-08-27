package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func (user User) SendMessage(message string, messageLength int) (string, bool) {
	if user.MessageCharLimit >= messageLength {
		return message, true
	}
	return "", false
}

func newUser(name string, membershipType membershipType) User {
	newUser := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 100,
		},
	}

	if newUser.Type == TypePremium {
		newUser.MessageCharLimit = 1000
	}
	return newUser
}
