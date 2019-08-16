package main

import (
	"fmt"
)

//main function for int
func main() {
	var x int
	var y string
	var w int32
	var z float32
	var b bool
	var m=4  //daynamic type
	n:=5     // daynamic type



	fmt.Println(x, w, y,b,m,n)
	fmt.Printf("%T %T %T %T %T", x, y, w, z, b)


}
