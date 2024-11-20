package main

import "fmt"

func main() {
	accounts := []string{"X", "Facebook"}

	if len(accounts) >= 3 {
		fmt.Println("Daha fazla hesaba sahip olamazsın.")
	} else {
		fmt.Println("Devam et")
	}

	// Ranks
	admin := 4
	moderator := 3
	premium := 2
	user := 1

	yourRank := 4

	if yourRank == admin {
		fmt.Println("Admin")
	} else if yourRank == moderator {
		fmt.Println("Moderatör")
	} else if yourRank == premium {
		fmt.Println("Premium kullanıcı")
	} else if yourRank == user {
		fmt.Println("Kullanıcı")
	}
}
