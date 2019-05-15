package message

import (
	"fmt"
)

type Message struct {
  text string
}

func (m Message) showText() string {
	fmt.Println(m.text)
	return m.text
}