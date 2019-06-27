package mrp_byproduct

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MrpSubproduct().DeclareModel()

	h.MrpSubproduct().AddFields(map[string]models.FieldDefinition{
		"ProductId": models.Many2OneField{
			RelationModel: h.ProductProduct(),
			String:        "Product",
			Required:      true,
		},
		"ProductQty": models.FloatField{
			String:  "Product Qty",
			Default: models.DefaultValue(1),
			//digits=dp.get_precision('Product Unit of Measure')
			Required: true,
		},
		"ProductUomId": models.Many2OneField{
			RelationModel: h.ProductUom(),
			String:        "Unit of Measure",
			Required:      true,
			//oldname='product_uom'
		},
		"BomId": models.Many2OneField{
			RelationModel: h.MrpBom(),
			String:        "BoM",
			OnDelete:      `cascade`,
		},
		"OperationId": models.Many2OneField{
			RelationModel: h.MrpRoutingWorkcenter(),
			String:        "Produced at Operation",
		},
	})
	h.MrpSubproduct().Methods().OnchangeProductId().DeclareMethod(
		` Changes UoM if product_id changes. `,
		func(rs m.MrpSubproductSet) {
			//        if self.product_id:
			//            self.product_uom_id = self.product_id.uom_id.id
		})
	h.MrpSubproduct().Methods().OnchangeUom().DeclareMethod(
		`OnchangeUom`,
		func(rs m.MrpSubproductSet) {
			//        res = {}
			//        if self.product_uom_id and self.product_id and self.product_uom_id.category_id != self.product_id.uom_id.category_id:
			//            res['warning'] = {
			//                'title': _('Warning'),
			//                'message': _('The Product Unit of Measure you chose has a different category than in the product form.')
			//            }
			//            self.product_uom_id = self.product_id.uom_id.id
			//        return res
		})
}
