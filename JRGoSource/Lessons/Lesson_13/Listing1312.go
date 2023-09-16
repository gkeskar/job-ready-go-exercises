package main

import (
   "fmt"
)

func main() {
   months := make(map[string]string)
   fmt.Println("Map value:", months)
   // add key-value pairs to the map
   months["Jan"] = "January"
   months["Feb"] = "February"
   months["Mar"] = "March"
   months["Apr"] = "April"
   months["May"] = "May"
   months["Jun"] = "June"
   months["Jul"] = "July"
   months["Bad"] = "BadMonth"
   months["Aug"] = "August"
   months["Sep"] = "September"
   months["Oct"] = "October"
   months["Nov"] = "November"
   months["Dec"] = "December"

   fmt.Println("Updated map:", months)

   // delete a key-value pair
   delete(months,"Bad")
   fmt.Println("Updated map:", months)
}
