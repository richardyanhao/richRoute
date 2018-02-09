package richRoute

import (
	"net/http"
	"strings"
	"fmt"
)

type route struct {
	root map[string]*node
}

func New() *route {
	return &route{
		root:make(map[string]*node),
	}
}

type handler func(http.ResponseWriter, *http.Request, Params)

func (r *route) GET(path string, handler handler)  {
	r.addRoutRules("GET", path, handler)
}

func (r *route) handle(method, path string, handler http.Handler) {
	r.addRoutRules(method, path, func(w http.ResponseWriter, r *http.Request, _ Params) {
		handler.ServeHTTP(w, r)
	})
}

func (r *route) addRoutRules(method string, path string, handler handler)  {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}
	segments := strings.Split(path, "/")[1:]

	selectRoot := r.root[method]
	if selectRoot == nil {
		selectRoot = &node{}
		r.root[method] = selectRoot
	}
	selectRoot.insertSubRoot(segments[0:], handler)
}

func (r *route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	root := r.root[req.Method]
	fmt.Printf("find %s \n", req.URL.Path)
	segments := strings.Split(req.URL.Path, "/")[1:]
	handler, err, p:= root.getHandler(segments)
	if err != nil {
		fmt.Print("get wrong")
	}
	if handler == nil{
		fmt.Println("richard handler is empty")
	}
	if p == nil{
		fmt.Println("richard params is empty")
	}
	handler(w, req, p)
	return
}


