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
}
