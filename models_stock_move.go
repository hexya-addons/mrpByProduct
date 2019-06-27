package mrp_byproduct

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.StockMove().DeclareModel()

	h.StockMove().AddFields(map[string]models.FieldDefinition{
		"SubproductId": models.Many2OneField{
			RelationModel: h.MrpSubproduct(),
			String:        "Subproduct",
			Help:          "Subproduct line that generated the move in a manufacturing order",
		},
	})
}
