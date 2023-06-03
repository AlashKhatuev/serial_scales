package main

import (
	"fmt"
	"serial_scales/serial"
)

func main() {
	message := serial.NewCommonMessage([]byte("asdasdasdasd"))
	b, err := serial.CommonMessageToBytes(message)
	if err != nil {
		panic(err)
	}
	mc, err := serial.BytesToCommonMessage(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(message)
	fmt.Println(mc)
}
