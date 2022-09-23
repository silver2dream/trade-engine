package matcher

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	pb "main/proto"
	"net"
	"time"
)

// Connect is net.Conn rename
type Connect net.Conn

// Session is save the client connect info
type Session struct {
	conn        Connect
	messageSend chan *pb.Packet
	messageRecv chan string
}

// Start if client connected
func (s *Session) Start() {
	go func() {
		buffer := make([]byte, 2048)
		result := bytes.NewBuffer(nil)
		for {
			n, err := s.conn.Read(buffer)
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
				scanner.Split(s.packetSlit)
				for scanner.Scan() {
					scannedPack := new(pb.Packet)
					scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
					fmt.Println(len(string(scanner.Bytes())))
					s.messageSend <- scannedPack
				}
			}
			result.Reset()
		}
	}()

	go func() {
		for {
			select {
			case msg := <-s.messageRecv:
				_, err := s.conn.Write([]byte(msg))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
}

// Stop if client disconnected
func (s *Session) Stop() {
	close(s.messageRecv)
	s.conn.Close()
	fmt.Printf("%s is disconnected", s.conn.RemoteAddr().String())
}

func (s *Session) packetSlit(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

func (s *Session) Send(msg proto.Message, tag string) {
	data, _ := proto.Marshal(msg)
	s.messageRecv <- s.Pack(data, tag)
}

func (s *Session) Pack(data []byte, tag string) string {
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
