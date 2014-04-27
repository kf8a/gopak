package pakbus

import "strings"

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
