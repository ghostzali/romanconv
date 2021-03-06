package romanconv

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestConvert(t *testing.T) {
	t.Run("Convert DCLXVIII should be success", func(t *testing.T) {
		got, _ := Convert("DCLXVIII")
		want := 668

		if want != got {
			t.Errorf("Roman to arabic, want %d got %d", want, got)
		}
	})

	t.Run("Convert MMMDCCXXIV should be success", func(t *testing.T) {
		got, _ := Convert("MMMDCCXXIV")
		want := 3724

		if want != got {
			t.Errorf("Roman to arabic, want %d got %d", want, got)
		}
	})

	t.Run("Convert _M_D_C_L_X_VMDCLXVI should be success", func(t *testing.T) {
		got, _ := Convert("_M_D_C_L_X_VMDCLXVI")
		want := 1666666

		if want != got {
			t.Errorf("Roman to arabic, want %d got %d", want, got)
		}
	})

	t.Run("Convert _M_M_M_C_M_X_C_I_XCMXCIX should be success", func(t *testing.T) {
		got, _ := Convert("_M_M_M_C_M_X_C_I_XCMXCIX")
		want := 3999999

		if want != got {
			t.Errorf("Roman to arabic, want %d got %d", want, got)
		}
	})

	t.Run("Convert _M_D_C_L_X_V_IM should be error", func(t *testing.T) {
		input := "_M_D_C_L_X_V_IM"
		_, err := Convert(input)

		if err != FORMAT_ERROR {
			t.Error("it sould return error format")
		}

		if err == nil {
			t.Errorf("expected to be error but it success. input %q", input)
		}
	})

	t.Run("Convert _M_M_M_M_D_C_L_X_VM should be error", func(t *testing.T) {
		input := "_M_M_M_M_D_C_L_X_VM"
		_, err := Convert(input)

		if err != FORMAT_ERROR {
			t.Error("it sould return error format")
		}

		if err == nil {
			t.Errorf("expected to be error but it success. input %q", input)
		}
	})
}

func TestParse(t *testing.T) {
	t.Run("Parse 668 should be success", func(t *testing.T) {
		got, _ := Parse(668)
		want := "DCLXVIII"

		if want != got {
			t.Errorf("Arabic to roman, want %s got %s", want, got)
		}
	})

	t.Run("Parse 3724 should be success", func(t *testing.T) {
		got, _ := Parse(3724)
		want := "MMMDCCXXIV"

		if want != got {
			t.Errorf("Arabic to roman, want %s got %s", want, got)
		}
	})

	t.Run("Parse 1666666 should be success", func(t *testing.T) {
		got, _ := Parse(1666666)
		want := "_M_D_C_L_X_VMDCLXVI"

		if want != got {
			t.Errorf("Arabic to roman, want %s got %s", want, got)
		}
	})

	t.Run("Parse 3999999 should be success", func(t *testing.T) {
		got, _ := Parse(3999999)
		want := "_M_M_M_C_M_X_C_I_XCMXCIX"

		if want != got {
			t.Errorf("Arabic to roman, want %s got %s", want, got)
		}
	})

	t.Run("Parse 4000000 should be error", func(t *testing.T) {
		input := 4000000
		_, err := Parse(input)

		if err != VALUE_ERROR {
			t.Error("it sould return error value")
		}

		if err == nil {
			t.Errorf("expected to be error but it success. input %d", input)
		}
	})
}

func TestPropertiesOfConvertions(t *testing.T) {
	assertion := func(arabic int) bool {
		if arabic < 0 || arabic > 3999999 {
			return true
		}
		roman, _ := Parse(arabic)
		fromRoman, _ := Convert(roman)
		return fromRoman == arabic
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("convertion error", err)
	}
}

func ExampleConvert() {
	roman := "_M_D_C_L_X_VMDCLXVI"
	arabic, _ := Convert(roman)
	fmt.Printf("Roman numeral %s equals to %d\n", roman, arabic)
	// Output: Roman numeral _M_D_C_L_X_VMDCLXVI equals to 1666666
}

func ExampleParse() {
	arabic := 71000
	roman, _ := Parse(arabic)
	fmt.Printf("Arabic number %d equals to %s\n", arabic, roman)
	// Output: Arabic number 71000 equals to _L_X_XM
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert("_M_M_M_C_M_X_C_I_XCMXCIX")
	}
}

// v1.0.0: BenchmarkConvert-4   	   20828	     56474 ns/op	   60810 B/op	     386 allocs/op
// v1.0.1: BenchmarkConvert-4   	   21126	     56476 ns/op	   60803 B/op	     386 allocs/op

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(i)
	}
}

// v1.0.0: BenchmarkParse-4   	 1664641	       723.1 ns/op	     148 B/op	      10 allocs/op
// v1.0.1: BenchmarkParse-4   	48186309	        20.95 ns/op	       4 B/op	       0 allocs/op
