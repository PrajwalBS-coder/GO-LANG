package main

import f "fmt"



func main(){
	x:=5
	f.Println(x)//Here we can say That pass by value so x won't change here
	f.Println(inc(x))// here It's returning function value
	f.Println(x)//here it won't change so we can say pass by value only sends its copy 

}
func inc(x int)int{
	return x+2
}