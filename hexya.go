package mrp_byproduct

import (
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "mrp_byproduct"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
