package main

import (
    "fmt"
    "math"
    "github.com/vernonthedev/goLearning/helloPrograms/strutil"
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
    //called the function from my newly created package to reverse
    //passed in string
    fmt.Println(strutil.Reverse("vernonthedev"))
    
    //format the new name using the newly created function
    fmt.Println(formatName("codezilla"))

    fmt.Println(getSum(12,34))
}


//function to return a formatted name
func formatName(name string) string {
    return "Welcome to the world of code: " + name
}

func getSum(num1, num2 int) int {
    return num1 + num2
}
