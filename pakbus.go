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
  output := unquoteSyncByte(pkt)
  return unquoteQuoteByte(output)
}
