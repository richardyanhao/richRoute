package richRoute

import (
	"net/http"
	"strings"
	"fmt"
)

type route struct {
	mid  []Middleware
	root map[string]*node
}

func New() *route {
	return &route{
		root:make(map[string]*node),
	}
}

type RichHandler interface {
	Do(http.ResponseWriter, *http.Request, Params)
}

type Handler func(http.ResponseWriter, *http.Request, Params)

func (h Handler) Do(w http.ResponseWriter, r *http.Request, p Params) {
	h(w, r, p)
}

func (r *route) GET(path string, handler RichHandler)  {
	r.addRoutRules("GET", path, handler)
}

func (r *route) addRoutRules(method string, path string, handler RichHandler)  {
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
	segments := strings.Split(req.URL.Path, "/")[1:]
	handler, err, p:= root.getHandler(segments)
	if err != nil {
		fmt.Print("get wrong")
	}
	if handler == nil{
		fmt.Println("handler is empty")
	}
	if r.mid != nil {
		for index := len(r.mid) - 1; index >= 0; index-- {
			handler = r.mid[index](w, req, p, handler)
		}
	}
	handler.Do(w, req, p)
	return
}

func (r *route) AddMiddleWare(m Middleware)  {
	if r.mid == nil {
		r.mid = make([]Middleware, 0)
	}
	r.mid = append(r.mid, m)
}


