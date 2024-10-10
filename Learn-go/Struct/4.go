package main

type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + auth.username + ":" + auth.password

}

// func main() {

// 	userauth := authenticationInfo{
// 		username: "ziad",
// 		password: "zizoshata2003",
// 	}

// 	fmt.Println(userauth.getBasicAuth())

// }
