package main

import ("fmt"
)



// var global 
var x int


 
func main() {


   // var local 
   var y int = 2

    
    x = 1
    
    fmt.Printf("\nx = %d, y = %d \n", x, y)


    // declara e inicializa
    z := 3
    fmt.Printf("\nz = %d \n\n", z)


}
