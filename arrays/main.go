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
    

    //declaring and assigning an array at once.
    sauces := [4]string{"Beans", "Gnuts","Meat","Greens"}
    fmt.Println(sauces)
    fmt.Println(sauces[1])

    //a sliced array is an array that is dynamic in terms of number of items
    slicedArray := []string{"BMW","Audi","RangeRover"}
    fmt.Println(len(slicedArray))
    fmt.Println(slicedArray)

    //more about slices
    var intSlice []int32 = []int32{4,5,6}
    fmt.Println(intSlice)
    intSlice = append(intSlice, 7)
    fmt.Println(intSlice)

}
