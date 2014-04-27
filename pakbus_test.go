package pakbus

import "testing"

func TestQuotePlain(t *testing.T) {
	const in, out = "trim", "trim"
	if x := Quote(in); x != out {
		t.Errorf("Quote(%s) expected %s got %s", in, out, x)
	}
}

func TestQuoteSyncByte(t *testing.T) {
	const in, out = "\xbcsync", "\xbc\xdcsync"
	if x := Quote(in); x != out {
		t.Errorf("Quote(%x) expected %x got %x", in, out, x)
	}
}

func TestQuoteQuoteByte(t *testing.T) {
	const in, out = "\xBDquote", "\xBC\xDDquote"
	if x := Quote(in); x != out {
		t.Errorf("Quote(%x) expected %x got %x", in, out, x)
	}
}

func TestUnquotePlain(t *testing.T) {
	const in, out = "trim", "trim"
	if x := UnQuote(in); x != out {
		t.Errorf("UnQuote(%s) expected %s got %s", in, out, x)
	}
}

func TestUnquoteSyncByte(t *testing.T) {
	const in, out = "\xbc\xdcsync", "\xbcsync"
	if x := UnQuote(in); x != out {
		t.Errorf("UnQuote(%x) expected %x got %x", in, out, x)
	}
}

func TestUnquoteQuoteByte(t *testing.T) {
	const in, out = "\xBC\xDDquote", "\xBDquote"
	if x := UnQuote(in); x != out {
		t.Errorf("UnQuote(%x) expected %x got %x", in, out, x)
	}
}

func TestCalcSigForByte(t *testing.T) {
	const in, out = 0xaa, 0xaaa9
	if x := CalcSigForByte(in, 0xaaaa); x != out {
		t.Errorf("CalcSigForByte(%x) expected %x got %x", in, out, x)
	}
}

func TestCalcFor(t *testing.T) {
	const in, out = "message", 0xbbdc
	if x := CalcSigFor([]byte(in), 0xaaaa); x != out {
		t.Errorf("CalcSigFor(%x) expected %x got %x", in, out, x)
	}
}
