package richRoute

import (
	"net/http"
	"fmt"
)

type Middleware func(http.ResponseWriter, *http.Request, Params, RichHandler) RichHandler

func PreMiddleWare(next RichHandler) RichHandler {
	return Handler(func(w http.ResponseWriter, req *http.Request, p Params){
		fmt.Fprint(w, "pre controller")
		next.Do(w, req, p)
	})
}

