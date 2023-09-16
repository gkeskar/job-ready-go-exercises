package main

import (
   "fmt"
)

func main() {

   // add key:value pairs to the map
   months := map[string]string{
      "Jan": "January",
      "Feb": "February",
      "Mar": "March",
      "Apr": "April",
      "May": "May",
      "Jun": "June",
      "Jul": "July",
      "Bad": "BadMonth",
      "Aug": "August",
      "Sep": "September",
      "Oct": "October",
      "Nov": "November",
      "Dec": "December"}

   fmt.Println("Map value:", months)

   // delete a key-value pair
   delete(months,"Bad")
   fmt.Println("Updated map:", months)
}