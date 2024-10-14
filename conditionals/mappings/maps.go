package mappings

import "fmt"

func PrintMails(){
    //define the map
    emails := make(map[string]string)

    //assign the values
    emails["vernon"] = "vernonthedev@gmail.com"
    emails["bum"] = "bum@gmail.com"
    emails["frea"] = "frea@gmail.com"
    //print some spaces
    fmt.Println("################################## \n")

    fmt.Println(emails)
    fmt.Println(emails["bum"])
    fmt.Println(len(emails))

    //delete values from the map
    delete(emails, "frea")
    fmt.Println(len(emails))
    fmt.Println(emails)


    //declaring and using maps at the same time
    names := map[string]string{"Gary":"Mugisha","Donald":"Mukembo","Mariot":"Sheraton"}

    fmt.Println(names)
    fmt.Println(len(names))
}

//looping through ranges
func Rangeloop(){
    //define an array of integers

    fmt.Println("################################## \n")

    ids := []int{12,34,56,67,78,89}

    //loop through the ids
    for i, id := range ids{
        //print the index and the id of the values in the integer array.
        fmt.Printf("%d - ID: %d\n", i, id)
    }
}
