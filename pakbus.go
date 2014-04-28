package pakbus

import "strings"

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

func CalcSigFor(buffer []uint8, seed uint16) uint16 {
	sig := seed
	for _, value := range buffer {
		sig = CalcSigForByte(value, sig)
	}
	return sig
}

// func CalcSigNullifier(sig uint16) uint16 {
//   var nulb []byte
//   var nullif []byte
//   for i := 0; i < 2; i++ {
//     sig = CalcSigFor(nulb, sig)
//     sig2 := (sig<<1) & 0x01ff
//     if sig2 >= 0x100 {
//       sig2 += 1
//     }
//     nullif[i] = ((0x100 - (sig2 + sig >> 8))) & 0xff
//   }
//   return nullif[0]
// }
