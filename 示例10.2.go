package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName) //Go在这么操作时会自动将指针对应于所指的位置
} //传入结构体指针进行修改
//同理，传入结构体本身会进行拷贝，此时无法修改原结构体的属性

func main() {
	var pers1 = Person{"Chris", "Woodward"}
	upPerson(pers1)
	fmt.Print(pers1.firstName, "\n", pers1.lastName)
}
