package main

import (
    "fmt"
    "math"
) 


func main(){
    //this is how we initialize the function variable
    name, email := "vernonthedev" , "vernonthedev@gmail.com"
	fmt.Println(name,email)
    fmt.Printf("%T\n",name)
    //get the highest rounded value after the decimal
    fmt.Println(math.Floor(2.7))
    fmt.Println(math.Ceil(2.7))
    fmt.Println(math.Sqrt(16))
}
