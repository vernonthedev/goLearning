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
