package main

import (
	"fmt"
)

func main()  {

 //method 1  static type variables
var x int32
var y float32
var b bool


//method 2
var userId,showId int32
var userName,showTitle string



//method 3
var name string="sajjad"
var deletede bool=true;
var json string ="JSON obj"




//method 3
var venue,status,countSeat,price="golestan",true,30,1542.4545445


//method 4
var name1,name2,name3 string ="sajjad","oppps","ohMyGash"


//method 5
var (
	xx=10
	yy string="sajjad rezaei"
	zz string
	bb bool=true
	cc=4.5

)



// method 6
x1:=10
y1:="sds,d.,"
nameN,nameM:="majjid","vahid"




//method 7

userName1:=func(){
	fmt.Println("salam")
}
userName1()



//method 8
_, xxx:=10,50


//method 9

var ddd=10
var ooo="jklsaaks"

//method 11
//arrays,slice,etc


println(x,y,b,userName,showTitle,userId,showId,name,deletede,json,venue,status,countSeat,price,name1,name2,name3,xx,yy,zz,bb,cc,x1,y1,nameN,nameM,xxx,ddd,ooo)
fmt.Printf("%T %T %T",x,y,b);






}