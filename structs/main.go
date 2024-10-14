package main

import "fmt"

//we create our first struct
type Person struct{
    //specify the datatypes for the vars
    firstName string
    lastName string
    city string
    gender string
    age int
}


func main(){
    //initialize the person struct
    firstPerson := Person{firstName: "Masamba", lastName: "Vernon", city: "California", gender: "Male", age: 20}
    fmt.Println(firstPerson)

}
