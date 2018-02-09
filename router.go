package richRoute

import (
	"net/http"
	"strings"
	"fmt"
	"errors"
)

type route struct {
	root map[string]*node
}

type node struct {
	path      string
	wildChild bool
	children  []*node
	handler   handler
}

func New() *route {
	return &route{
		root:make(map[string]*node),
	}
}

type handler func(http.ResponseWriter, *http.Request, Params)

type Param struct {
	Key   string
	Value string
}

type Params struct {
	Key   string
	Value string
}

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
		selectRoot = &node{
			path:      segments[0],
			wildChild: segments[0][0] == ':',
			handler: handler,
		}
		r.root[method] = selectRoot
	}
	selectRoot.insertSubRoot(segments[1:], handler)
}

func (r *route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	root := r.root[req.Method]
	segments := strings.Split(req.URL.Path, "/")[1:]
	handler, err:= root.getHandler(segments)
	//pickHandle()
	if err != nil || handler == nil {
		// return not found
	}
	handler(w, req, _)
}

func (n *node) insertSubRoot(segments []string, handler handler)  {
	if  len(segments) == 0{
		return
	}
	var pickNode *node

	for _, v := range n.children  {
		if v.path == segments[0] {
			pickNode = v
		}
		if v.wildChild && segments[0][0] == ':' {
			fmt.Printf("it is another wild node, %s will be use as %s", segments[0], n.path)
			// pickNode = v
			return
		}
	}
	if pickNode == nil {
		pickNode = &node{
			path:      segments[0],
			wildChild: segments[0][0] == ':',
			handler: handler,
		}
		n.children = append(n.children, pickNode)
	}

	pickNode.insertSubRoot(segments[1:], handler)
}

func (n *node) getHandler(segments []string) (handler, error) {
	if segments[0] != n.path {
		return notfoundHandler, errors.New("not found")
	} else if len(segments) == 1 {
		if n.handler == nil {
			return notfoundHandler, nil
		}
		return n.handler, nil
	}
	var pickNode *node
	for _, childNode := range n.children {
		if childNode.path == segments[1] {
			pickNode = childNode
		}
	}
	if pickNode == nil {
		return notfoundHandler, errors.New("not found")
	}

	return notfoundHandler, errors.New("not found")
}

func notfoundHandler(w http.ResponseWriter, r *http.Request, _ Params) {
	fmt.Fprint(w, "not found")
}

func (n *node) ShowNode() {
	fmt.Printf("path is : %s ,wild or not: %t, children: %+v \n", n.path, n.wildChild, n.children)
	for _, v := range n.children {
		v.ShowNode()
	}
}

func (r *route) Show() {
	for k, v := range r.root {
		fmt.Printf("method:%s \n", k)
		v.ShowNode()
	}
}

