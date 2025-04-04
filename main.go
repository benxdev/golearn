package main
import (
	"fmt"
)

type messageToSend struct {
	message string
	sender user
	recepient user
}
type user struct {
	name string
	number int
}
func canSendmessage(mToSend, messageToSend) bool {
	return true
}

func main() {
	fmt.Print()
}