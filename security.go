package mrp_byproduct

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/pool/h"
)

//vars

//rights
func init() {
	h.MrpSubproduct().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpSubproduct().Methods().AllowAllToGroup(GroupMrpManager)
	h.MrpSubproduct().Methods().Load().AllowGroup(base.GroupUser)
}
