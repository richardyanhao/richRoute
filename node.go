package richRoute

import (
	"fmt"
	"net/http"
)

type node struct {
	path      string
	wildChild bool
	children  []*node
	handler   handler
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

func (n *node) getHandler(segments []string) (handler, error, Params) {
	var ps Params
	var pickNode *node
	for _, v := range n.children {
		if v.path == segments[0] || v.wildChild {
			pickNode = v
			break
		}
	}
	if pickNode == nil {
		return notfoundHandler, nil, ps
	}
	if pickNode.wildChild {
		p := Param{
			Key: pickNode.path[1:],
			Value: segments[0],
		}
		ps = make(Params,1)
		ps = append(ps, p)
	}

	if len(segments) == 1 {
		return pickNode.handler, nil, ps
	}

	th, te, tp := pickNode.getHandler(segments[1:])
	if n.wildChild {
		return th, te, MergeSlice(ps, tp)
	}
	return th, te, tp
}

func MergeSlice(s1 Params, s2 Params) Params {
	slice := make(Params, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
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