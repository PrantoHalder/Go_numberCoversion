package main

import (
	"fmt"
)

func valueConversion(){
	var myNumber int
	fmt.Println("enter a number ")
	fmt.Scanf("%d",&myNumber)
	fmt.Printf("This is the binary %b\n",myNumber)
	fmt.Printf("this is the octal %o\n",myNumber)
	fmt.Printf("this is the hexa decimal %x\n",myNumber)
}