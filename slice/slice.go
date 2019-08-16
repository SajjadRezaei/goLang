package main

import (
	"fmt"
)

var (
	i []int
)
func main()  {

i=make([]int,3)
i[0]=10
i[1]=20
i[2]=30
fmt.Println("slice:===>",i)
j:=[]string{"Java","C#","Php","Ruby","Go"}
fmt.Println(j)

v:=[]bool{}

fmt.Println(v)


fmt.Println("I:",i,len(i))
fmt.Println("J:",j,len(j))
fmt.Println("V:",v,len(v))


fmt.Println("j[:0]==>",j[:0])
fmt.Println("j[1:3]==>",j[1:3])
fmt.Println("j[:2]==>",j[:2])
fmt.Println("j[2:]==>",j[2:])


fmt.Println("Before split :",j)

j=j[:4]


fmt.Println("After Split ",j,len(j))

j=j[2:]
fmt.Println(j,len(j))




k:=make([]bool,4)
k[0]=true
k[1]=true
k[2]=false
k[3]=true
//k[4]=true  error out of range

fmt.Println(k)

  k=append(k,true)
fmt.Println(k)



p:=[]float32{3.14,15.22,21251.45,455,414,1424}
fmt.Println(p)

p= append(p, 24124,1414,4445)

fmt.Println(p)



empty:=make([]float32 ,len(p))

copy(empty,p)

fmt.Println(empty)



	for i:=0;i<len(p);i++  {
		fmt.Printf("%f \t",p[i])
	}

      printSlice(p)






}



func printSlice(s []float32){
	for i:=0;i<len(s);i++  {
		fmt.Printf("%f \t",s[i])
	}
}
