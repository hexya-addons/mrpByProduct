package mrp_byproduct

import (
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ChangeProductionQty().DeclareModel()

	h.ChangeProductionQty().Methods().UpdateProductToProduce().DeclareMethod(
		`UpdateProductToProduce`,
		func(rs m.ChangeProductionQtySet, prod interface{}, qty interface{}) {
			//        super(ChangeProductionQty, self)._update_product_to_produce(prod, qty)
			//        Production = self.env['mrp.production']
			//        UoM = self.env['product.uom']
			//        for sub_product_line in prod.bom_id.sub_products:
			//            move = prod.move_finished_ids.filtered(
			//                lambda x: x.subproduct_id == sub_product_line and x.state not in ('done', 'cancel'))
			//            if move:
			//                product_uom_factor = prod.product_uom_id._compute_quantity(
			//                    prod.product_qty - prod.qty_produced, prod.bom_id.product_uom_id)
			//                qty1 = sub_product_line.product_qty
			//                qty1 *= product_uom_factor / prod.bom_id.product_qty
			//                move[0].write({'product_uom_qty': qty1})
			//            else:
			//                prod._create_byproduct_move(sub_product_line)
		})
}
