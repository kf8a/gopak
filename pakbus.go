package pakbus

import (
	"bytes"
	"encoding/binary"
	// "github.com/synful/gopack"
	"log"
	// "math/big"
	"strings"
)

type Packet struct {
	start_sync byte
	header     [8]byte
	message    Message
	nullifier  [2]byte
	end_sync   byte
}

type Header struct {
	link_state          byte //4 bits
	destination_address byte //[12]bits MSB first
	expect_more         byte // 2 bits
	priority            byte // 2 bits
	source_address      byte //12 bits MSB first
	proto_code          byte // 4 bits
	destination_node_id byte // 12 bits MSB first
	hop_count           byte //4 bits
	source_node_id      byte // 12 bits MSB first
}

type Message struct {
	message_type   byte
	transaction_id byte
	body           []byte
}

func quoteSyncByte(pkt string) string {
	return strings.Replace(pkt, "\xbc", "\xbc\xdc", -1)
}

func quoteQuoteByte(pkt string) string {
	return strings.Replace(pkt, "\xbd", "\xbc\xdd", -1)
}

func Quote(pkt string) string {
	output := quoteSyncByte(pkt)
	return quoteQuoteByte(output)
}

func unquoteSyncByte(pkt string) string {
	return strings.Replace(pkt, "\xbc\xdc", "\xbc", -1)
}

func unquoteQuoteByte(pkt string) string {
	return strings.Replace(pkt, "\xbc\xdd", "\xbd", -1)
}

func UnQuote(pkt string) string {
	output := unquoteQuoteByte(pkt)
	return unquoteSyncByte(output)
}

func CalcSigForByte(value byte, seed uint16) uint16 {
	sig := (seed << 1) & 0x1ff
	if sig >= 0x100 {
		sig += 1
	}
	return (((sig + (seed >> 8) + uint16(value)) & 0xff) | (seed << 8)) & 0xffff
}

func CalcSigFor(buffer []byte, seed uint16) uint16 {
	sig := seed

	for _, value := range buffer {
		sig = CalcSigForByte(value, sig)
	}
	return sig
}

func convertNullifierToBuffer(nulb uint16, size int) []byte {
	buf := new(bytes.Buffer)
	rbuf := make([]byte, size)

	if err := binary.Write(buf, binary.LittleEndian, nulb); err != nil {
		log.Fatal("write nulb failed %v", nulb)
	}
	if _, err := buf.Read(rbuf); err != nil {
		log.Fatal("read of nulb to buffer failed")
	}
	return rbuf
}

func CalcSigNullifier(sig uint16) uint16 {
	var nullif, nulb uint16

	for i := 0; i < 2; i++ {
		rbuf := convertNullifierToBuffer(nulb, i)
		sig := CalcSigFor(rbuf, sig)

		sig2 := (sig << 1) & 0x1FF
		if sig2 >= 0x100 {
			sig2 += 1
		}

		nulb = ((0x100 - (sig2 + (sig >> 8))) & 0xFF)

		nullif = nullif << 8
		nullif += nulb
	}
	return nullif
}

type SerPkt struct {
	LinkState  byte
	ExpectMore byte
	Priority   byte
	Dest       uint16
	Src        uint16
}

func (s *SerPkt) Encode() [4]byte {
	var buf [4]byte

	buf[0] = s.LinkState<<4 | uint8(s.Dest>>12)
	buf[1] = uint8(s.Dest)
	buf[2] = s.ExpectMore<<6 | s.Priority<<4 | uint8(s.Src>>12)
	buf[3] = uint8(s.Src)
	return buf
}

// func (s *SerPkt) Decode([4]byte) SerPkt {
// }

type PakbusHdr struct {
	LinkState  byte
	ExpectMore byte
	Priority   byte
	Protocol   byte
	Dest       uint16
	HopCount   byte
	Src        uint16
}

func (h *PakbusHdr) Encode() [8]byte {
	var buf [8]byte

	buf[0] = h.LinkState<<4 | uint8(h.Dest>>12)
	buf[1] = uint8(h.Dest)
	buf[2] = h.ExpectMore<<6 | h.Priority<<4 | uint8(h.Src>>12)
	buf[3] = uint8(h.Src)
	buf[4] = h.Protocol<<4 | uint8(h.Dest>>12)
	buf[5] = uint8(h.Dest)
	buf[6] = h.HopCount<<4 | uint8(h.Src>>12)
	buf[7] = uint8(h.Src)

	return buf
}
