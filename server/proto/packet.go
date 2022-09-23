package proto

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	BUY = iota
	SELL
	CANCEL
	NOT_CANCELLED
	PARTIAL
	FULL
)

const (
	Tag1000 = "t_1000"
	Tag1001 = "t_1001"
	Tag1002 = "t_1002"
	Tag1003 = "t_1003"
)

type Packet struct {
	VersionLen uint32
	Version    []byte
	TagLen     uint32
	Tag        []byte
	Timestamp  uint32
	DataLen    uint32
	Data       []byte
}

func (p *Packet) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.LittleEndian, &p.VersionLen)
	err = binary.Write(writer, binary.LittleEndian, &p.Version)
	err = binary.Write(writer, binary.LittleEndian, &p.TagLen)
	err = binary.Write(writer, binary.LittleEndian, &p.Tag)
	err = binary.Write(writer, binary.LittleEndian, &p.Timestamp)
	err = binary.Write(writer, binary.LittleEndian, &p.DataLen)
	err = binary.Write(writer, binary.LittleEndian, &p.Data)
	return err
}

func (p *Packet) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.LittleEndian, &p.VersionLen)
	p.Version = make([]byte, p.VersionLen)
	err = binary.Read(reader, binary.LittleEndian, &p.Version)
	err = binary.Read(reader, binary.LittleEndian, &p.TagLen)
	p.Tag = make([]byte, p.TagLen)
	err = binary.Read(reader, binary.LittleEndian, &p.Tag)
	err = binary.Read(reader, binary.LittleEndian, &p.Timestamp)
	err = binary.Read(reader, binary.LittleEndian, &p.DataLen)
	p.Data = make([]byte, p.DataLen)
	err = binary.Read(reader, binary.LittleEndian, &p.Data)
	return err
}

func (p *Packet) GetTag() []byte {
	return p.Tag
}

func (p *Packet) String() string {
	return fmt.Sprintf("version:%s dataLen:%d timestamp:%d tag:%s data:%s",
		p.Version,
		p.DataLen,
		p.Timestamp,
		p.Tag,
		p.Data,
	)
}

func Int32(value int32) *int32 {
	v := new(int32)
	*v = value
	return v
}

func Float32(value float32) *float32 {
	v := new(float32)
	*v = value
	return v
}

func String(value string) *string {
	return &value
}
