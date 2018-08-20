package core

import (
	"testing"
)

func TestReadableSize(t *testing.T) {
	got := ReadableSize(12.123)
	expect := "12.12B"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}

	got = ReadableSize(1124.123)
	expect = "1.10K"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}

	got = ReadableSize(1024 * 1024 * 8)
	expect = "8.00M"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}

	got = ReadableSize(1024 * 1024 * 1024 * 8)
	expect = "8.00G"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}

	got = ReadableSize(1024 * 1024 * 1024 * 1024 * 8)
	expect = "8.00T"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}

	got = ReadableSize(1024 * 1024 * 1024 * 1024 * 1024 * 8)
	expect = "8.00P"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}

	got = ReadableSize(1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 8)
	expect = "TooLarge"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}

	got = ReadableSize(0)
	expect = "0.00B"
	if got != expect {
		t.Errorf("ReadableSize:\n Expect => %v\n Got    => %v\n", expect, got)
	}
}
