package agent

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	pb "main/proto"
	"main/utility"
	"net"
	"os"
	"strings"
	"time"
)

type Handler func([]string)

type Agent struct {
	traderId uint32
	tradeId  uint32
	funcMap  map[string]Handler
	send     chan string
	orders   []*pb.Order
}

func NewAgent() *Agent {
	p := &Agent{
		funcMap: make(map[string]Handler),
		send:    make(chan string, 65535),
		tradeId: 1,
	}

	p.funcMap["b"] = p.Buy
	p.funcMap["s"] = p.Sell
	p.funcMap["c"] = p.Cancel
	p.funcMap["l"] = p.OrderList

	return p
}

func (p *Agent) Run() {
	servAddr := "localhost:8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	go func() {
		buffer := make([]byte, 2048)
		result := bytes.NewBuffer(nil)
		for {
			n, err := conn.Read(buffer)
			result.Write(buffer[0:n])
			if err != nil {
				if err == io.EOF {
					continue
				} else {
					fmt.Println("read err:", err)
					break
				}
			} else {
				scanner := bufio.NewScanner(result)
				scanner.Split(packetSlit)
				for scanner.Scan() {
					scannedPack := new(pb.Packet)
					scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
					if bytes.Compare(scannedPack.GetTag(), []byte(pb.TraderID)) == 0 {
						t := &pb.TradeSession{}
						proto.Unmarshal(scannedPack.Data, t)
						p.traderId = t.GetTraderId()
						fmt.Println(t.GetTraderId())
					} else if bytes.Compare(scannedPack.GetTag(), []byte(pb.Buy)) == 0 {
						t := &pb.Order{}
						proto.Unmarshal(scannedPack.Data, t)
						fmt.Println(t)
					} else if bytes.Compare(scannedPack.GetTag(), []byte(pb.Cancel)) == 0 {
						t := &pb.Order{}
						proto.Unmarshal(scannedPack.Data, t)
						fmt.Println(t)
					}
				}
			}
			result.Reset()
		}
	}()

	go func() {
		for {
			select {
			case data := <-p.send:
				_, err := conn.Write([]byte(data))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		//把換行符號去掉
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		args := strings.Split(text, " ")
		if len(args) < 1 {
			fmt.Println("text can't split")
			continue
		}
		fmt.Println("len[%d] - args:%v", len(args), args)
		args2 := args[1:]
		p.DoCommandFunc(args[0], args2)
	}
}

func (p *Agent) DoCommandFunc(cmd string, args []string) {
	if hander, found := p.funcMap[cmd]; found {
		hander(args)
	}
}

func (p *Agent) Buy(args []string) {
	if len(args) < 3 {
		fmt.Println("args not enough.")
		return
	}

	stockId, _ := utility.Interface2uint64(args[0])
	quantity, _ := utility.Interface2uint64(args[1])
	price, _ := utility.Interface2uint64(args[2])

	o := &pb.Order{
		Uuid:     p.traderId,
		StockId:  stockId,
		TradeId:  p.tradeId,
		Kind:     pb.BUY,
		Price:    price,
		Quantity: quantity,
	}
	data, _ := proto.Marshal(o)
	p.send <- p.Pack(data, pb.Buy)
	p.orders = append(p.orders, o)
	p.tradeId++
}

func (p *Agent) Sell(args []string) {
	if len(args) < 3 {
		fmt.Println("args not enough.")
		return
	}

	stockId, _ := utility.Interface2uint64(args[0])
	quantity, _ := utility.Interface2uint64(args[1])
	price, _ := utility.Interface2uint64(args[2])

	o := &pb.Order{
		Uuid:     p.traderId,
		StockId:  stockId,
		Kind:     pb.SELL,
		Price:    price,
		Quantity: quantity,
	}
	data, _ := proto.Marshal(o)
	p.send <- p.Pack(data, pb.Sell)
	p.orders = append(p.orders, o)
	p.tradeId++
}

func (p *Agent) Cancel(args []string) {
	if len(args) < 2 {
		fmt.Println("args not enough.")
		return
	}

	stockId, _ := utility.Interface2uint64(args[0])
	tradeId, _ := utility.Interface2uint32(args[1])

	o := &pb.Order{
		Uuid:    p.traderId,
		TradeId: tradeId,
		StockId: stockId,
		Kind:    pb.CANCEL,
	}
	data, _ := proto.Marshal(o)
	p.send <- p.Pack(data, pb.Cancel)
	p.orders = append(p.orders, o)
	p.tradeId++
}

func (p *Agent) OrderList(args []string) {
	for _, order := range p.orders {
		fmt.Println(order)
	}
}

func packetSlit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && data[4] == 'V' {
		var headerlen uint32 = 6
		tagLen := binary.LittleEndian.Uint32(data[6:10]) // if taglen = 6
		timestampStart := uint32(10) + tagLen            //16
		headerlen = timestampStart + uint32(4)           //20
		offset := headerlen + uint32(4)
		realLen := binary.LittleEndian.Uint32(data[headerlen:offset])
		headerlen += uint32(4)
		if uint32(len(data)) > realLen {
			calSum := int(realLen + headerlen)
			if calSum <= len(data) {
				return calSum, data[:calSum], nil
			}
		}
	}
	return
}

func (p *Agent) Pack(data []byte, tag string) string {
	writeBuf := bytes.NewBuffer(nil)
	resPack := new(pb.Packet)
	resPack.VersionLen = 2
	resPack.Version = []byte("V1")
	resPack.TagLen = 6
	resPack.Tag = []byte(tag)
	resPack.Timestamp = uint32(time.Now().Unix())
	resPack.DataLen = uint32(len(data))
	resPack.Data = data
	resPack.Pack(writeBuf)

	return string(writeBuf.Bytes())
}
