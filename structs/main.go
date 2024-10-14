package main

import (
    "fmt"
    "strconv"
)

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
    secondPerson := Person{"Penny","Kans","California","Female",20}
    fmt.Println(secondPerson)
    fmt.Println("##############SPACE##############")

    fmt.Println(secondPerson.firstName)
    fmt.Println(secondPerson.firstName + secondPerson.lastName)


    fmt.Println("##############SPACE##############")
    fmt.Println(secondPerson.greet())
    secondPerson.hasBirthday()
    secondPerson.getMarried("Masamba")
    fmt.Println(secondPerson.greet())
}

// Greeting method - value receiver
func (person Person) greet() string {
    return "Hello my name is " + 
        person.firstName + " "+ 
        person.lastName + " and I am " +    
        strconv.Itoa(person.age)
}

//hasBirthday method - pointer receiver
func (person *Person) hasBirthday(){
    //changing the person's age value in the original struct
    person.age++
}

//getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string){
    if person.gender == "Male"{
        return
    }else{
        person.lastName = spouseLastName
    }
}
