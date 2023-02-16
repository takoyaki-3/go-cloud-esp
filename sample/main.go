package main

import (
	"fmt"
	"log"

	. "github.com/takoyaki-3/go-cloud-esp"
)

func main(){
	fmt.Println("start")

	esp1,err := NewESP("conf.json")
	if err!=nil{
		log.Fatalln(err)
	}

	esp1.SetPinMode(5,"OUTPUT")
	esp1.SetPinMode(35,"INPUT")

	for {
		fmt.Println("loop")
		esp1.WriteDigital(5,"HIGH")
		err := esp1.WriteDigital(5,"LOW")
		esp1.ReadAnalog(35)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
