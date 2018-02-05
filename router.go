package richRoute

import (
	"net/http"
	"strings"
	"fmt"
)

type route struct {
	root map[string]*node
}

type node struct {
	path      string
	wildChild bool
	children  []*node
	handler   http.Handler
}

func New() *route {
	return &route{
		root:make(map[string]*node),
	}
}

func (r *route) GET(path string, handler http.Handler)  {
	r.addRoutRules("GET", path, handler)
}

func (r *route) addRoutRules(method string, path string, handler http.Handler)  {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}
	segments := strings.Split(path, "/")[1:]

	selectRoot := r.root[method]
	if selectRoot == nil {
		selectRoot = &node{
			path:      segments[0],
			wildChild: segments[0][0] == ':',
		}
		r.root[method] = selectRoot
	}
	selectRoot.insertSubRoot(segments[1:])
}

func (n *node) insertSubRoot(segments []string)  {
	if  len(segments) == 0{
		return
	}
	var pickNode *node

	for _, v := range n.children  {
		if v.path == segments[0] {
			pickNode = v
		}
		if v.wildChild {
			fmt.Printf("it is a wild node, %s will be use as %s", segments[0], n.path)
			return
		}
	}
	if pickNode == nil {
		pickNode = &node{
			path:      segments[0],
			wildChild: segments[0][0] == ':',
		}
		n.children = append(n.children, pickNode)
	}

	pickNode.insertSubRoot(segments[1:])
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

