package matcher

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"main/matcher/pqueue"
	pb "main/proto"
	"net"
	"strings"
	"sync"
)

type TradeMatcher struct {
	sessions    map[uint32]*Session
	send        chan string
	recv        chan *pb.Packet
	matchQueues map[uint64]*pqueue.MatchQueues
	slab        *pqueue.Slab

	traderId uint32
	r        sync.RWMutex
}

func NewMatcher() *TradeMatcher {
	p := &TradeMatcher{
		matchQueues: make(map[uint64]*pqueue.MatchQueues),
		slab:        pqueue.NewSlab(20),
		send:        make(chan string, 65535),
		recv:        make(chan *pb.Packet, 65535),
	}
	return p
}

func (m *TradeMatcher) Start(potocalType string, ip string) error {
	sock, err := net.Listen(potocalType, ip)
	if err != nil {
		return err
	}

	defer sock.Close()
	log.Println("Wait for clients")

	m.process()

	for {
		conn, err := sock.Accept()
		if err != nil {
			return err
		}
		log.Println(conn.RemoteAddr().String(), potocalType, "connect success")
		m.addToSessions(conn)
		defer conn.Close()
	}
}

func (m *TradeMatcher) process() {
	go func() {
		for {
			select {
			case packet := <-m.recv:
				fmt.Println(packet.String())
				order, err := m.UnPack(packet)
				if err != nil {
					log.Println(err)
				}

				on := m.slab.Malloc()
				on.CopyFrom(order)
				switch order.GetKind() {
				case pb.BUY:
					m.addBuy(on)
				case pb.SELL:
					m.addSell(on)
				case pb.CANCEL:
					m.cancel(on)
				default:
					panic(fmt.Sprintf("MsgKind %v not supported", order))
				}
			}
		}
	}()
}

func (m *TradeMatcher) addBuy(order *pqueue.OrderNode) {
	q := m.getMatchQueues(order.StockId())
	if !m.fillableBuy(order, q) {
		q.PushBuy(order)
	}
}

func (m *TradeMatcher) fillableBuy(b *pqueue.OrderNode, q *pqueue.MatchQueues) bool {
	for {
		s := q.PeekSell()
		if s == nil {
			return false
		}
		if b.Price() >= s.Price() {
			if b.Quantity() > s.Quantity() {
				quantity := s.Quantity()
				price := price(b.Price(), s.Price())
				s.Remove()
				m.slab.Free(s)
				b.ReduceQuantity(quantity)
				m.completeTrade(pb.PARTIAL, pb.FULL, b, s, price, quantity)
				continue // The sell has been used up
			}
			if s.Quantity() > b.Quantity() {
				quantity := b.Quantity()
				price := price(b.Price(), s.Price())
				s.ReduceQuantity(quantity)
				m.completeTrade(pb.PARTIAL, pb.FULL, b, s, price, quantity)
				m.slab.Free(b)
				return true // The buy has been used up
			}
			if s.Quantity() == b.Quantity() {
				quantity := b.Quantity()
				price := price(b.Price(), s.Price())
				m.completeTrade(pb.PARTIAL, pb.FULL, b, s, price, quantity)
				s.Remove()
				m.slab.Free(s)
				m.slab.Free(b)
				return true // The buy and sell have been used up
			}
		} else {
			return false
		}
	}
}

func price(price uint64, price2 uint64) uint64 {
	d := price - price2
	return price2 + (d / 2)
}

func (m *TradeMatcher) getMatchQueues(stockId uint64) *pqueue.MatchQueues {
	q := m.matchQueues[stockId]
	if q == nil {
		q = &pqueue.MatchQueues{}
		m.matchQueues[stockId] = q
	}
	return q
}

func (m *TradeMatcher) addToSessions(conn net.Conn) error {
	if conn == nil {
		return errors.New("conn is null")
	}

	m.r.Lock()
	defer m.r.Unlock()

	sess := &Session{
		conn:        conn,
		messageSend: m.recv,
		messageRecv: make(chan string, 65535),
	}

	traderSession := &pb.TradeSession{
		TradeId: m.traderId,
	}
	m.sessions[traderSession.GetTradeId()] = sess
	sess.Start()
	m.traderId++

	sess.Send(traderSession, pb.Tag1000)
	return nil
}

func (m *TradeMatcher) getSessionIdentity(ip string) string {
	ipinfo := strings.Split(ip, ":")
	ip = ipinfo[0]
	fmt.Println(ip)
	h := sha1.New()
	h.Write([]byte(ip))
	bs := h.Sum(nil)
	identity := hex.EncodeToString(bs)
	return identity
}

func (m *TradeMatcher) UnPack(packet *pb.Packet) (*pb.Order, error) {
	o := &pb.Order{}
	proto.Unmarshal(packet.Data, o) //save data len 20~24
	return o, nil
}

func (m *TradeMatcher) addSell(s *pqueue.OrderNode) {
	q := m.getMatchQueues(s.StockId())
	if !m.fillableSell(s, q) {
		q.PushSell(s)
	}
}

func (m *TradeMatcher) fillableSell(s *pqueue.OrderNode, q *pqueue.MatchQueues) bool {
	for {
		b := q.PeekBuy()
		if b == nil {
			return false
		}
		if b.Price() >= s.Price() {
			if b.Quantity() > s.Quantity() {
				amount := s.Quantity()
				price := price(b.Price(), s.Price())
				b.ReduceQuantity(amount)
				m.completeTrade(pb.PARTIAL, pb.FULL, b, s, price, amount)
				s.Remove()
				m.slab.Free(s)
				return true // The sell has been used up
			}
			if s.Quantity() > b.Quantity() {
				amount := b.Quantity()
				price := price(b.Price(), s.Price())
				s.ReduceQuantity(amount)
				m.completeTrade(pb.PARTIAL, pb.FULL, b, s, price, amount)
				b.Remove()
				m.slab.Free(b) // The buy has been used up
				continue
			}
			if s.Quantity() == b.Quantity() {
				amount := b.Quantity()
				price := price(b.Price(), s.Price())
				m.completeTrade(pb.PARTIAL, pb.FULL, b, s, price, amount)
				b.Remove()
				m.slab.Free(b)
				m.slab.Free(s)
				return true // The sell and buy have been used up
			}
		} else {
			return false
		}
	}
}

func (m *TradeMatcher) cancel(o *pqueue.OrderNode) {
	q := m.getMatchQueues(o.StockId())
	ro := q.Cancel(o)
	if ro != nil {
		m.completeCancelled(ro)
		m.slab.Free(ro)
	} else {
		m.completeNotCancelled(o)
	}
	m.slab.Free(o)
}

func (m *TradeMatcher) completeTrade(partial int32, full int32, b *pqueue.OrderNode, s *pqueue.OrderNode, price uint64, quantity uint64) {
	m.r.RLock()
	defer m.r.RUnlock()

	if buyer, found := m.sessions[b.Uuid()]; found {
		order := &pb.Order{
			Uuid:     b.Uuid(),
			TradeId:  b.TradeId(),
			StockId:  b.StockId(),
			Kind:     partial,
			Price:    price,
			Quantity: quantity,
		}
		buyer.Send(order, pb.Tag1001)
	}

	if seller, found := m.sessions[s.Uuid()]; found {
		order := &pb.Order{
			Uuid:     s.Uuid(),
			TradeId:  s.TradeId(),
			StockId:  s.StockId(),
			Kind:     full,
			Price:    price,
			Quantity: quantity,
		}
		seller.Send(order, pb.Tag1001)
	}
}

func (m *TradeMatcher) completeCancelled(o *pqueue.OrderNode) {
	cm := pb.Order{}
	o.CopyTo(&cm)
	cm.Kind = pb.CANCEL

	m.r.RLock()
	defer m.r.RUnlock()
	if trader, found := m.sessions[o.Uuid()]; found {
		trader.Send(&cm, pb.Tag1002)
	}
}

func (m *TradeMatcher) completeNotCancelled(nc *pqueue.OrderNode) {
	ncm := pb.Order{}
	nc.CopyTo(&ncm)
	ncm.Kind = pb.NOT_CANCELLED

	m.r.RLock()
	defer m.r.RUnlock()
	if trader, found := m.sessions[nc.Uuid()]; found {
		trader.Send(&ncm, pb.Tag1003)
	}
}
