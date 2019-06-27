package mrp_byproduct

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MrpBom().DeclareModel()

	h.MrpBom().AddFields(map[string]models.FieldDefinition{
		"SubProducts": models.One2ManyField{
			RelationModel: h.MrpSubproduct(),
			ReverseFK:     "",
			String:        "Byproducts",
			NoCopy:        false,
		},
	})
}
