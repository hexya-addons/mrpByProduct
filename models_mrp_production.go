package mrp_byproduct

import (
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MrpProduction().DeclareModel()

	h.MrpProduction().Methods().CreateByproductMove().DeclareMethod(
		`CreateByproductMove`,
		func(rs m.MrpProductionSet, sub_product interface{}) {
			//        Move = self.env['stock.move']
			//        for production in self:
			//            source = production.product_id.property_stock_production.id
			//            product_uom_factor = production.product_uom_id._compute_quantity(
			//                production.product_qty - production.qty_produced, production.bom_id.product_uom_id)
			//            qty1 = sub_product.product_qty
			//            qty1 *= product_uom_factor / production.bom_id.product_qty
			//            data = {
			//                'name': 'PROD:%s' % production.name,
			//                'date': production.date_planned_start,
			//                'product_id': sub_product.product_id.id,
			//                'product_uom_qty': qty1,
			//                'product_uom': sub_product.product_uom_id.id,
			//                'location_id': source,
			//                'location_dest_id': production.location_dest_id.id,
			//                'operation_id': sub_product.operation_id.id,
			//                'production_id': production.id,
			//                'origin': production.name,
			//                'unit_factor': qty1 / (production.product_qty - production.qty_produced),
			//                'subproduct_id': sub_product.id
			//            }
			//            move = Move.create(data)
			//            move.action_confirm()
		})
	h.MrpProduction().Methods().GenerateMoves().DeclareMethod(
		` Generates moves and work orders
        @return: Newly generated picking Id.
        `,
		func(rs m.MrpProductionSet) {
			//        res = super(MrpProduction, self)._generate_moves()
			//        for production in self.filtered(lambda production: production.bom_id):
			//            for sub_product in production.bom_id.sub_products:
			//                production._create_byproduct_move(sub_product)
			//        return res
		})
}
