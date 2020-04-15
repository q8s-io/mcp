package filters

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
)

func AddFilters(container *restful.Container) {
	container.Filter(func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		fmt.Println("just test filter")
		chain.ProcessFilter(req, resp)
	})
}
