package main

import "fmt"

var (
	coins        = uint32(50)
	users        = []string{"loop", "luna"}
	distribution = make(map[string]uint32, len(users))
)

func calcUserCoinCount(user string) (count uint32) {
	for _, v := range user {
		switch v {
		case 'E':
			fallthrough
		case 'e':
			count++
		case 'I':
			fallthrough
		case 'i':
			count += 2
		case 'O':
			fallthrough
		case 'o':
			count += 3
		case 'U':
			fallthrough
		case 'u':
			count += 4
		}
	}
	return
}

func main() {
	fmt.Println("hello, world")

	for _, user := range users {
		if userCoin := calcUserCoinCount(user); coins > userCoin {
			distribution[user] = userCoin
			coins -= userCoin
		} else {
			distribution[user] = coins
			coins = 0
		}
	}

	fmt.Printf("%v\n%#v\n", distribution, distribution)
	fmt.Println(coins)
}
