package main

import "fmt"

func main(){
    //declare the array
    //we specify the size as well as the datatype
    var fruits [3]string
    
    //assign the array values
    fruits[0] = "Apple"
    fruits[1] = "Mango"
    fruits[2] = "Pineapple"

    fmt.Println(fruits)
    fmt.Println(fruits[2])


}
