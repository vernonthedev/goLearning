package main

import (
    "fmt"
    "github.com/vernonthedev/goLearning/conditionals/mappings"
)

func main(){
   // name := "vernon"
    age := 12


    if age > 10 {
        fmt.Println("He is old enough to become a developer.")
    }else{
        fmt.Println("He is not eligible for this program.")
    }

    //creating a for loop
    for i:=1; i <=10; i++ {
        fmt.Printf("Number: %d\n",i)
    }

    //run the imported pack file
    mappings.PrintMails()

    //get the integer range ids
    mappings.Rangeloop()


    //more about maps
    var myMap = map[string]uint8{"Adam":28, "Sarah":45}
    fmt.Println(myMap["Adam"])
    //return true if value exists in the map and false if it doesn't
    var years, ok = myMap["Jason"]
    if ok{
        fmt.Printf("The age is %v", years)
    }else{
        fmt.Println("INVALID NAME")
    }

   // this 
   // is some
   // gibberish code
   // am supposed to comment out
}
