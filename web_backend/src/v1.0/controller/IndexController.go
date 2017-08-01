package controller

import (
	"fmt"

	"v1.0/vendor"
)

// IndexController struct
type IndexController struct {
	vendor.Controller
}

func (c *IndexController) Index() {
	sess := globalSessions.SessionStart(c.GetResponseWriter(), c.GetRequest())

	fmt.Println(sess.Get("id"))
}
