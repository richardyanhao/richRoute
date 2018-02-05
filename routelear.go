package richRoute

import (
	"fmt"
	"httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	route := httprouter.New()
	fmt.Print(route)
	route.GET("/man", Hello)
	fmt.Println(route)
	fmt.Println((route.GetTree())["GET"])
	route.GET("/man/:id", Hello)
	route.GET("/man/:id/x", Hello)
	route.GET("/manxx", Hello)
	value := (route.GetTree())["GET"]
	fmt.Println(value)
	fmt.Println("the children of the node")
	for _, v := range value.GetChild() {
		fmt.Println(v)
		fmt.Println("the children of the node deep")
		for _, v2 := range v.GetChild() {
			fmt.Println(v2)
			fmt.Println("the children of the node deep")
			for _, v := range v2.GetChild() {
				fmt.Println(v)
			}
		}
	}
}
