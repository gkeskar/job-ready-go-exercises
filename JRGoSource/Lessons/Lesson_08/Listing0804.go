package main

import "fmt" 

func rectStuff(length int, width int) (int, int) {
   a := length * width
   c := length + length + width + width 
   return a, c 
}

func main() {
   area, perimeter := rectStuff(3, 5)
   fmt.Println("area:", area)
   fmt.Println("perimeter:", perimeter)
}