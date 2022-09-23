package pqueue

import (
	"github.com/fmstephe/flib/fmath"
	"main/proto"
)

type OrderNode struct {
	priceNode node
	guidNode  node
	quantity  uint64
	stockId   uint64
	kind      int32
	nextFree  *OrderNode
}

func (o *OrderNode) CopyFrom(from *proto.Order) {
	o.quantity = from.GetQuantity()
	o.stockId = from.StockId
	o.kind = from.GetKind()
	o.setup(from.Price, uint64(fmath.CombineInt32(int32(from.GetUuid()), int32(from.TradeId))))
}

func (o *OrderNode) CopyTo(to *proto.Order) {
	to.Kind = o.Kind()
	to.Price = o.Price()
	to.Quantity = o.Quantity()
	to.Uuid = o.Uuid()
	to.TradeId = o.TradeId()
	to.StockId = o.StockId()
}

func (o *OrderNode) setup(price, guid uint64) {
	initNode(o, price, &o.priceNode, &o.guidNode)
	initNode(o, guid, &o.guidNode, &o.priceNode)
}

func (o *OrderNode) Price() uint64 {
	return o.priceNode.val
}

func (o *OrderNode) Guid() uint64 {
	return o.guidNode.val
}

func (o *OrderNode) Uuid() uint32 {
	return uint32(fmath.HighInt32(int64(o.guidNode.val)))
}

func (o *OrderNode) TradeId() uint32 {
	return uint32(fmath.LowInt32(int64(o.guidNode.val)))
}

func (o *OrderNode) Quantity() uint64 {
	return o.quantity
}

func (o *OrderNode) ReduceQuantity(s uint64) {
	o.quantity -= s
}

func (o *OrderNode) StockId() uint64 {
	return o.stockId
}

func (o *OrderNode) Kind() int32 {
	return o.kind
}

func (o *OrderNode) Remove() {
	o.priceNode.pop()
	o.guidNode.pop()
}

//func (o *OrderNode) String() string {
//	if o == nil {
//		return "<nil>"
//	}
//	price := fstrconv.ItoaDelim(int64(o.Price()), ',')
//	quantity := fstrconv.ItoaDelim(int64(o.Amount()), ',')
//	traderId := fstrconv.ItoaDelim(int64(o.TraderId()), '-')
//	tradeId := fstrconv.ItoaDelim(int64(o.TradeId()), '-')
//	stockId := fstrconv.ItoaDelim(int64(o.StockId()), '-')
//	kind := o.kind
//	return fmt.Sprintf("%v, price %s, quantity %s, trader %s, trade %s, stock %s", kind, price, quantity, traderId, tradeId, stockId)
//}
