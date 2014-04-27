package pakbus

import "testing"

func TestQuotePlain(t *testing.T) {
  const in, out = "trim", "trim"
  if x := Quote(in); x != out{
		t.Errorf("Quote(%s) expected %s got %s", in, out, x )
	}
}

func TestQuoteSyncByte(t *testing.T) {
  const in, out = "\xbcsync", "\xbc\xdcsync"
  if x := Quote(in); x != out {
    t.Errorf("Quote(%x) expected %x got %x", in, out, x )
  }
}

func TestQuoteQuoteByte(t * testing.T) {
  const in, out = "\xBDquote", "\xBC\xDDquote"
  if x := Quote(in); x != out {
    t.Errorf("Quote(%x) expected %x got %x", in, out, x )
  }
}

func testUnquotePlain(t *testing.T) {
  const in, out = "trim", "trim"
  if x := UnQuote(in); x != out{
		t.Errorf("UnQuote(%s) expected %s got %s", in, out, x )
	}
}

func testUnquoteSyncByte(t *testing.T) {
  const in, out = "\xbc\xdcsync", "\xbcsync"
  if x := UnQuote(in); x != out {
    t.Errorf("Quote(%x) expected %x got %x", in, out, x )
  }
}

func testUnquoteQuoteByte(t *testing.T) {
  const in, out = "\xBC\xDDquote", "\xBDquote"
  if x := UnQuote(in); x != out {
    t.Errorf("Quote(%x) expected %x got %x", in, out, x )
  }
}
