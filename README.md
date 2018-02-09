# richRoute
Go practice. This is a route part for a web framework

```
func Index(w http.ResponseWriter, r *http.Request, _ richRoute.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps richRoute.Params) {
	val, _ := ps.GetValByKey("name")
	fmt.Fprintf(w, "hello, %s!\n", val)
}

func main() {
	r := richRoute.New()
	r.GET("/hello", Index)
	r.GET("/welcome/:name", Hello)
	err := http.ListenAndServe(":3000", r) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```
