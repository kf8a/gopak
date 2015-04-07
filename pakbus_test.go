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
	var testdata = []struct {
		in  string
		out uint16
	}{
		{"message", 0x1a17},
		{"testing", 0x1ef5},
	}
	for _, tt := range testdata {
		if x := CalcSigFor([]byte(tt.in), 0xaaaa); x != tt.out {
			t.Errorf("CalcSigFor(%x) expected %x got %x", tt.in, tt.out, x)
		}
	}
}

func TestCalcSigFor(t *testing.T) {
	var testdata = []struct {
		buff string
		seed uint16
		out  uint16
	}{
		{"", 0x0a, 0x0a},
		{"", 0x1C, 0x1c},
		{"", 0x134C, 0x134c},
	}
	for _, tt := range testdata {
		x := CalcSigFor([]byte(tt.buff), tt.seed)
		if x != tt.out {
			t.Errorf("CalcSigFor(%x, %x) expected %x got %x", tt.buff, tt.seed, tt.out, x)
		}
	}
}

func TestCalcSigNullifier(t *testing.T) {
	const in, out = 0x1a17, 0xb8e9
	if x := CalcSigNullifier(in); x != out {
		t.Errorf("CalcSigNullifer(%x) expected %x got %x", in, out, x)
	}
}
