package main

type User struct {
	Name string
	Membership
}

func (Us User) SendMessage(message string, messageLinght int) (resulte string, flag bool) {

	if Us.MessageCharLimit >= messageLinght {
		resulte = message
		flag = true
		return
	}

	return

}

func newUser(name string, membershipType string) User {
	newUser := User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
		},
	}
	if newUser.Type == "premium" {
		newUser.MessageCharLimit = 1000
	} else {
		newUser.MessageCharLimit = 100
	}

	return newUser
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

// func main() {

// 	ziad := newUser("ziad", "premium")

// 	_, flag := ziad.SendMessage("hi man", 8888)
// 	fmt.Println(flag)

// }
