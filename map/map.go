package main
import ("fmt")
var(
	m map[string]string
)
func main() {
	m = make(map[string]string)
	m["namne"] = "sajjad"
	m["family"] = "rezaei"
	m["lang"] = "Go"
	m["Ab"] = "df"

	c := map[string]string{
		"name1": "sajjad",
		"age":   "22",
	}

	delete(c, "name")

	fmt.Println(c)
	delete(m, "Ab")

	fmt.Println(m, "\r\n", m["family"])

	//empty map

	myMap := map[bool]bool{}
	fmt.Println(myMap)
	if v, ok := c["name1"]; ok {
		println("exist ", c["name"])
		println(v)
		println(ok)
	} else {
		println("not found")
	}
	for key, value := range m {
		fmt.Printf("Key : %v Value : %v\r\n", key, value)
	}
	//for key, value := range c {
	//	fmt.Println(key, "  ", value)
	//}
	//


//nested ma

mapc:=map[string]map[string]string{
	"A":map[string]string{"prefix":"Ali","suffix":"baba"},
	"B":map[string]string{"prefix":"bad","suffix":"bob"},
}


println(mapc["A"])




}