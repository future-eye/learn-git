package main

import (
	"fmt"
)

func sushu(n int){
	var a int
	var b=1
	for i:=2;i<n;i++{
		a=n%i
		if a==0{
			b=0
			break
		}
	}
	if b==1{
		fmt.Printf("%d\n",n)
	}
}

func main(){
	for n:=2;n<=10000;n++{
		go sushu(n)
	}
}