package pakbus

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

type PakbusHdr struct {
	LinkState      byte   `gopack:"4"`
	DestPhyAddress uint64 `gopackL"12"`
	ExpectMore     byte   `gopack:"2"`
	Priority       byte   `gopack:"2"`
	SrcPhyAddress  uint64 `gopack:"12"`
	Protocol       byte   `gopack:"4"`
	Dst            uint64 `gopack:"12"`
	HopCount       byte   `gopack:"4"`
	Src            uint64 `gopack:"12"`
}

func (h *PakbusHdr) Encode() [8]byte {
	// hdr := new(big.Int)

	// var buf = make([]byte, 8)

	// gopack.Pack(buf, h)

	var buf [8]byte

	var address = make([]byte, 4)
	binary.PutUvarint(address, h.DestPhyAddress)
	fmt.Print(address)
	buf[0] = h.LinkState<<4 + address[0]
	buf[2] = address[1]

	buf[2] = h.ExpectMore<<4 + h.Priority

	// buf := []byte{160, 147, 232, 0, 34, 23, 1, 64}
	// hdr.SetUint64(123456990812347890)
	// hdr.SetBytes(buf)

	// return hdr.Bytes()
	return buf
}
