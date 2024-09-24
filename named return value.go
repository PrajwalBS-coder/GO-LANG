package main

import f "fmt"

func main(){
	a,b:=fun()
	f.Println(a,b)
	c,d:=fun2()
	f.Println(c,d)


}
func fun()(x, y int){
	// x=0
	// y=0
	// return
	return 5,7
}
func fun2()(int,int){
	var x int
	var y int
	return x ,y
}