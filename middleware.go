package richRoute

import (
	"net/http"
	"fmt"
)

func AddMiddleWare(next RichHandler) RichHandler {
	return handler(func(w http.ResponseWriter, req *http.Request, p Params){
		fmt.Fprint(w, "pre controller")
		next.Do(w, req, p)
	})
}