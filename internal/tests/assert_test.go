package tests

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"testing"
)

func wrap(t *testing.T) *tWrapper {
	return &tWrapper{T: t}
}

type SomeStruct struct {
	Value int
}

func TestTrueFalse(t *testing.T) {
	wt := wrap(t)
	if !True(wt, true) {
		t.Errorf("True(true) failed")
	}
	if !True(wt, true, "msg %v", 42) {
		t.Errorf("True(true) with extra msg failed")
	}
	if !False(wt, false) {
		t.Errorf("False(false) failed")
	}
	if !False(wt, false, "msg %v", 42) {
		t.Errorf("False(false) with extra msg failed")
	}

	True(wt, false, "msg %v", 42)
	if wt.LastFailure().ExtraMsgStr != "msg 42" {
		t.Errorf("True(false) with extra msg failed")
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		E interface{}
		V interface{}
		O bool
	}{
		{int(42), int(42), true},
		{int(42), uint64(42), false},
		{int(42), int(255), false},
		{"foo", "foo", true},
		{"foo", "bar", false},
		{[]string{}, []string{}, false},
		{nil, nil, true},
		{nil, (*int)(nil), false},
		{(*int)(nil), nil, false},
		{(*int)(nil), (*int)(nil), true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) == %v (%T)",
			tc.E, tc.E, tc.V, tc.V), func(t *testing.T) {

			wt := wrap(t)

			if Equal(wt, tc.E, tc.V) != tc.O {
				t.Errorf("%v (%T) == %v (%T) should be %v",
					tc.E, tc.E, tc.V, tc.V, tc.O)
			}

		})
	}
}

func TestNotEqual(t *testing.T) {
	cases := []struct {
		E interface{}
		V interface{}
		O bool
	}{
		{int(42), int(42), false},
		{int(42), uint64(42), false},
		{int(42), int(255), true},
		{"foo", "foo", false},
		{"foo", "bar", true},
		{[]string{}, []string{}, false},
		{nil, nil, false},
		{nil, (*int)(nil), false},
		{(*int)(nil), nil, false},
		{(*int)(nil), (*int)(nil), false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) != %v (%T)",
			tc.E, tc.E, tc.V, tc.V), func(t *testing.T) {

			wt := wrap(t)

			if NotEqual(wt, tc.E, tc.V) != tc.O {
				t.Errorf("%v (%T) != %v (%T) should be %v",
					tc.E, tc.E, tc.V, tc.V, tc.O)
			}

		})
	}
}

func TestEqualTypes(t *testing.T) {
	cases := []struct {
		E interface{}
		V interface{}
		O bool
	}{
		{int(42), int(42), true},
		{int(42), uint64(42), false},
		{int(42), int(255), true},
		{"foo", "foo", true},
		{"foo", "bar", true},
		{[]string{}, []string{}, true},
		{nil, nil, true},
		{nil, (*int)(nil), false},
		{(*int)(nil), nil, false},
		{(*int)(nil), (*int)(nil), true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) == %v (%T)",
			tc.E, tc.E, tc.V, tc.V), func(t *testing.T) {

			wt := wrap(t)

			if EqualTypes(wt, tc.E, tc.V) != tc.O {
				t.Errorf("%v (%T) == %v (%T) should be %v",
					tc.E, tc.E, tc.V, tc.V, tc.O)
			}

		})
	}
}

func TestNotEqualTypes(t *testing.T) {
	cases := []struct {
		E interface{}
		V interface{}
		O bool
	}{
		{int(42), int(42), false},
		{int(42), uint64(42), true},
		{int(42), int(255), false},
		{"foo", "foo", false},
		{"foo", "bar", false},
		{[]string{}, []string{}, false},
		{nil, nil, false},
		{nil, (*int)(nil), true},
		{(*int)(nil), nil, true},
		{(*int)(nil), (*int)(nil), false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) == %v (%T)",
			tc.E, tc.E, tc.V, tc.V), func(t *testing.T) {

			wt := wrap(t)

			if NotEqualTypes(wt, tc.E, tc.V) != tc.O {
				t.Errorf("%v (%T) == %v (%T) should be %v",
					tc.E, tc.E, tc.V, tc.V, tc.O)
			}

		})
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		E interface{}
		C interface{}
		O bool
	}{
		{int(42), []int{1, 3, 7, 42, 105}, true},
		{int(42), []int{1, 3, 7, 41, 105}, false},
		{int32(42), []int64{1, 3, 7, 42, 105}, false},
		{"needle", []string{"foo", "bar", "needle", "banana"}, true},
		{"needle", []string{"foo", "bar", "twig", "banana"}, false},
		{"bar", map[string]int8{"foo": 1, "bar": 2, "twig": 3, "banana": 4}, true},
		{"bar", map[string]int8{"foo": 1, "beer": 2, "twig": 3, "banana": 4}, false},
		{"nan", "banana", true},
		{"nan", "batman", false},
		{42, map[string]int8{"foo": 1, "beer": 2, "twig": 3, "banana": 4}, false},
		{42, nil, false},
		{42, 42, false},
		{42, "42", false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) in %v (%T)",
			tc.E, tc.E, tc.C, tc.C), func(t *testing.T) {

			wt := wrap(t)

			if Contains(wt, tc.E, tc.C) != tc.O {
				t.Errorf("%v (%T) in %v (%T) should be %v",
					tc.E, tc.E, tc.C, tc.C, tc.O)
			}

		})
	}
}

func TestNotContains(t *testing.T) {
	cases := []struct {
		E interface{}
		C interface{}
		O bool
	}{
		{int(42), []int{1, 3, 7, 42, 105}, false},
		{int(42), []int{1, 3, 7, 41, 105}, true},
		{int32(42), []int64{1, 3, 7, 42, 105}, true},
		{"needle", []string{"foo", "bar", "needle", "banana"}, false},
		{"needle", []string{"foo", "bar", "twig", "banana"}, true},
		{"bar", map[string]int8{"foo": 1, "bar": 2, "twig": 3, "banana": 4}, false},
		{"bar", map[string]int8{"foo": 1, "beer": 2, "twig": 3, "banana": 4}, true},
		{"nan", "banana", false},
		{"nan", "batman", true},
		{42, map[string]int8{"foo": 1, "beer": 2, "twig": 3, "banana": 4}, false},
		{42, nil, false},
		{42, 42, false},
		{42, "42", false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) in %v (%T)",
			tc.E, tc.E, tc.C, tc.C), func(t *testing.T) {

			wt := wrap(t)

			if NotContains(wt, tc.E, tc.C) != tc.O {
				t.Errorf("%v (%T) in %v (%T) should be %v",
					tc.E, tc.E, tc.C, tc.C, tc.O)
			}

		})
	}
}

func TestNil(t *testing.T) {
	cases := []struct {
		V interface{}
		O bool
	}{
		{nil, true},
		{interface{}(nil), true},
		{42, false},
		{"test", false},
		{"", false},
		{(*string)(nil), true},
		{&[]string{"test"}[0], false},
		{[]string{"test"}, false},
		{[]string{}, false},
		{&[]string{}, false},
		{[]string(nil), true},
		{SomeStruct{42}, false},
		{&SomeStruct{42}, false},
		{(*SomeStruct)(nil), true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T)", tc.V, tc.V), func(t *testing.T) {
			wt := wrap(t)
			if Nil(wt, tc.V) != tc.O {
				t.Errorf("Nil(%v (%T)) should be %v", tc.V, tc.V, tc.O)
			}
		})
	}
}

func TestNotNil(t *testing.T) {
	cases := []struct {
		V interface{}
		O bool
	}{
		{nil, false},
		{interface{}(nil), false},
		{42, true},
		{"test", true},
		{"", true},
		{(*string)(nil), false},
		{&[]string{"test"}[0], true},
		{[]string{"test"}, true},
		{[]string{}, true},
		{&[]string{}, true},
		{[]string(nil), false},
		{SomeStruct{42}, true},
		{&SomeStruct{42}, true},
		{(*SomeStruct)(nil), false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T)", tc.V, tc.V), func(t *testing.T) {
			wt := wrap(t)
			if NotNil(wt, tc.V) != tc.O {
				t.Errorf("NotNil(%v (%T)) should be %v", tc.V, tc.V, tc.O)
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	cases := []struct {
		V interface{}
		O bool
	}{
		{nil, false},
		{interface{}(nil), false},
		{42, false},
		{"test", false},
		{"", true},
		{(*string)(nil), false},
		{&[]string{"test"}[0], false},
		{[]string{"test"}, false},
		{[]string{}, true},
		{&[]string{}, false},
		{[]string(nil), true},
		{map[string]int{}, true},
		{map[string]int(nil), true},
		{map[string]int{"foo": 42}, false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T)", tc.V, tc.V), func(t *testing.T) {
			wt := wrap(t)
			if Empty(wt, tc.V) != tc.O {
				t.Errorf("Empty(%v (%T)) should be %v", tc.V, tc.V, tc.O)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	cases := []struct {
		V interface{}
		O bool
	}{
		{nil, false},
		{interface{}(nil), false},
		{42, false},
		{"test", true},
		{"", false},
		{(*string)(nil), false},
		{&[]string{"test"}[0], false},
		{[]string{"test"}, true},
		{[]string{}, false},
		{&[]string{}, false},
		{[]string(nil), false},
		{map[string]int{}, false},
		{map[string]int(nil), false},
		{map[string]int{"foo": 42}, true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T)", tc.V, tc.V), func(t *testing.T) {
			wt := wrap(t)
			if NotEmpty(wt, tc.V) != tc.O {
				t.Errorf("NotEmpty(%v (%T)) should be %v", tc.V, tc.V, tc.O)
			}
		})
	}
}

func TestDeepEqual(t *testing.T) {
	cases := []struct {
		E interface{}
		V interface{}
		O bool
	}{
		{int(42), int(42), true},
		{int(42), uint64(42), false},
		{int(42), int(255), false},
		{"foo", "foo", true},
		{"foo", "bar", false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) == %v (%T)",
			tc.E, tc.E, tc.V, tc.V), func(t *testing.T) {

			wt := wrap(t)

			if DeepEqual(wt, tc.E, tc.V) != tc.O {
				t.Errorf("%v (%T) == %v (%T) should be %v",
					tc.E, tc.E, tc.V, tc.V, tc.O)
			}

		})
	}
}

func TestNotDeepEqual(t *testing.T) {
	cases := []struct {
		E interface{}
		V interface{}
		O bool
	}{
		{int(42), int(42), false},
		{int(42), uint64(42), true},
		{int(42), int(255), true},
		{"foo", "foo", false},
		{"foo", "bar", true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v (%T) != %v (%T)",
			tc.E, tc.E, tc.V, tc.V), func(t *testing.T) {

			wt := wrap(t)

			if NotDeepEqual(wt, tc.E, tc.V) != tc.O {
				t.Errorf("%v (%T) != %v (%T) should be %v",
					tc.E, tc.E, tc.V, tc.V, tc.O)
			}

		})
	}
}

func TestFailed(t *testing.T) {
	wt := wrap(t)
	Failed(wt, "foo")
	if wt.LastFailure().ReasonStr != "foo" {
		t.Errorf("obvious test failed")
	}
	Failed(wt, "foo %v", 42)
	if wt.LastFailure().ReasonStr != "foo 42" {
		t.Errorf("obvious test2 failed")
	}
}

// --- Error assertions ---

var errSentinel = errors.New("sentinel")

type myErr struct{ msg string }

func (e *myErr) Error() string { return e.msg }

func TestNoError(t *testing.T) {
	cases := []struct {
		Name string
		Err  error
		O    bool
	}{
		{"nil", nil, true},
		{"non-nil", errors.New("boom"), false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if NoError(wt, tc.Err) != tc.O {
				t.Errorf("NoError(%v) should be %v", tc.Err, tc.O)
			}
		})
	}
}

func TestError(t *testing.T) {
	cases := []struct {
		Name string
		Err  error
		O    bool
	}{
		{"nil", nil, false},
		{"non-nil", errors.New("boom"), true},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if Error(wt, tc.Err) != tc.O {
				t.Errorf("Error(%v) should be %v", tc.Err, tc.O)
			}
		})
	}
}

func TestErrorIs(t *testing.T) {
	wrapped := fmt.Errorf("context: %w", errSentinel)
	cases := []struct {
		Name   string
		Err    error
		Target error
		O      bool
	}{
		{"identical", errSentinel, errSentinel, true},
		{"wrapped matches", wrapped, errSentinel, true},
		{"unrelated", errors.New("other"), errSentinel, false},
		{"nil err non-nil target", nil, errSentinel, false},
		{"both nil", nil, nil, true},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if ErrorIs(wt, tc.Err, tc.Target) != tc.O {
				t.Errorf("ErrorIs(%v, %v) should be %v", tc.Err, tc.Target, tc.O)
			}
		})
	}
}

func TestNotErrorIs(t *testing.T) {
	wrapped := fmt.Errorf("context: %w", errSentinel)
	cases := []struct {
		Name   string
		Err    error
		Target error
		O      bool
	}{
		{"identical", errSentinel, errSentinel, false},
		{"wrapped matches", wrapped, errSentinel, false},
		{"unrelated", errors.New("other"), errSentinel, true},
		{"nil err non-nil target", nil, errSentinel, true},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if NotErrorIs(wt, tc.Err, tc.Target) != tc.O {
				t.Errorf("NotErrorIs(%v, %v) should be %v", tc.Err, tc.Target, tc.O)
			}
		})
	}
}

func TestErrorAs(t *testing.T) {
	base := &myErr{msg: "boom"}
	wrapped := fmt.Errorf("context: %w", base)

	t.Run("matches in wrapped chain", func(t *testing.T) {
		wt := wrap(t)
		var target *myErr
		if !ErrorAs(wt, wrapped, &target) {
			t.Errorf("ErrorAs should find *myErr in wrapped chain")
		}
		if target == nil || target.msg != "boom" {
			t.Errorf("target should be populated, got %v", target)
		}
	})

	t.Run("type not in chain", func(t *testing.T) {
		wt := wrap(t)
		var target *myErr
		if ErrorAs(wt, errors.New("plain"), &target) {
			t.Errorf("ErrorAs should not find *myErr in plain error")
		}
	})

	t.Run("nil err", func(t *testing.T) {
		wt := wrap(t)
		var target *myErr
		if ErrorAs(wt, nil, &target) {
			t.Errorf("ErrorAs should fail on nil err")
		}
	})
}

func TestErrorContains(t *testing.T) {
	cases := []struct {
		Name   string
		Err    error
		Substr string
		O      bool
	}{
		{"contains substring", errors.New("disk full: out of space"), "out of space", true},
		{"empty substring matches non-nil", errors.New("anything"), "", true},
		{"no match", errors.New("anything"), "missing", false},
		{"nil err", nil, "anything", false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if ErrorContains(wt, tc.Err, tc.Substr) != tc.O {
				t.Errorf("ErrorContains(%v, %q) should be %v", tc.Err, tc.Substr, tc.O)
			}
		})
	}
}

// --- Ordering assertions (one representative per operator) ---

func TestGreaterInt(t *testing.T) {
	cases := []struct {
		L, R int
		O    bool
	}{
		{2, 1, true},
		{1, 1, false},
		{1, 2, false},
		{-1, -2, true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d>%d", tc.L, tc.R), func(t *testing.T) {
			wt := wrap(t)
			if GreaterInt(wt, tc.L, tc.R) != tc.O {
				t.Errorf("GreaterInt(%d, %d) should be %v", tc.L, tc.R, tc.O)
			}
		})
	}
}

func TestGreaterOrEqualFloat64(t *testing.T) {
	cases := []struct {
		L, R float64
		O    bool
	}{
		{2.0, 1.0, true},
		{1.0, 1.0, true},
		{1.0, 2.0, false},
		{math.NaN(), 1.0, false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v>=%v", tc.L, tc.R), func(t *testing.T) {
			wt := wrap(t)
			if GreaterOrEqualFloat64(wt, tc.L, tc.R) != tc.O {
				t.Errorf("GreaterOrEqualFloat64(%v, %v) should be %v", tc.L, tc.R, tc.O)
			}
		})
	}
}

func TestLessString(t *testing.T) {
	cases := []struct {
		L, R string
		O    bool
	}{
		{"apple", "banana", true},
		{"banana", "apple", false},
		{"apple", "apple", false},
		{"", "a", true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%q<%q", tc.L, tc.R), func(t *testing.T) {
			wt := wrap(t)
			if LessString(wt, tc.L, tc.R) != tc.O {
				t.Errorf("LessString(%q, %q) should be %v", tc.L, tc.R, tc.O)
			}
		})
	}
}

func TestLessOrEqualInt(t *testing.T) {
	cases := []struct {
		L, R int
		O    bool
	}{
		{1, 2, true},
		{1, 1, true},
		{2, 1, false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d<=%d", tc.L, tc.R), func(t *testing.T) {
			wt := wrap(t)
			if LessOrEqualInt(wt, tc.L, tc.R) != tc.O {
				t.Errorf("LessOrEqualInt(%d, %d) should be %v", tc.L, tc.R, tc.O)
			}
		})
	}
}

// --- Collection assertions ---

func TestLen(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	cases := []struct {
		Name      string
		Container interface{}
		Expected  int
		O         bool
	}{
		{"slice match", []int{1, 2, 3}, 3, true},
		{"slice mismatch", []int{1, 2, 3}, 2, false},
		{"empty slice", []int{}, 0, true},
		{"nil slice", []int(nil), 0, true},
		{"array", [3]int{1, 2, 3}, 3, true},
		{"map", map[string]int{"a": 1, "b": 2}, 2, true},
		{"string bytes not runes", "héllo", 6, true},
		{"channel buffered", ch, 2, true},
		{"non-container int", 42, 0, false},
		{"nil interface", nil, 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if Len(wt, tc.Container, tc.Expected) != tc.O {
				t.Errorf("Len(%v, %d) should be %v", tc.Container, tc.Expected, tc.O)
			}
		})
	}
}

func TestElementsMatch(t *testing.T) {
	cases := []struct {
		Name string
		A, B interface{}
		O    bool
	}{
		{"identical order", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"different order", []int{1, 2, 3}, []int{3, 1, 2}, true},
		{"both empty", []int{}, []int{}, true},
		{"different lengths", []int{1, 2, 3}, []int{1, 2}, false},
		{"different elements", []int{1, 2, 3}, []int{1, 2, 4}, false},
		{"duplicates different multiplicity", []int{1, 1, 2}, []int{1, 2, 2}, false},
		{"duplicates same multiplicity", []int{1, 2, 1}, []int{2, 1, 1}, true},
		{"strings", []string{"a", "b"}, []string{"b", "a"}, true},
		{"non-slice arg", 42, []int{}, false},
		{"cross-type", []int{1}, []int64{1}, false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if ElementsMatch(wt, tc.A, tc.B) != tc.O {
				t.Errorf("ElementsMatch(%v, %v) should be %v", tc.A, tc.B, tc.O)
			}
		})
	}
}

func TestNotElementsMatch(t *testing.T) {
	cases := []struct {
		Name string
		A, B interface{}
		O    bool
	}{
		{"different order match", []int{1, 2, 3}, []int{3, 1, 2}, false},
		{"different elements", []int{1, 2, 3}, []int{1, 2, 4}, true},
		{"non-slice arg", 42, []int{}, false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if NotElementsMatch(wt, tc.A, tc.B) != tc.O {
				t.Errorf("NotElementsMatch(%v, %v) should be %v", tc.A, tc.B, tc.O)
			}
		})
	}
}

// --- Pointer identity ---

func TestSame(t *testing.T) {
	x := 42
	y := 42
	s := "foo"
	var nilInt *int
	var nilStr *string

	cases := []struct {
		Name string
		E, A interface{}
		O    bool
	}{
		{"same pointer", &x, &x, true},
		{"different pointers same type", &x, &y, false},
		{"different pointer types", &x, &s, false},
		{"both typed nil same type", nilInt, nilInt, true},
		{"both typed nil different types", nilInt, nilStr, false},
		{"non-pointer", 42, 42, false},
		{"one non-pointer", &x, 42, false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if Same(wt, tc.E, tc.A) != tc.O {
				t.Errorf("Same(%v, %v) should be %v", tc.E, tc.A, tc.O)
			}
		})
	}
}

func TestNotSame(t *testing.T) {
	x := 42
	y := 42
	var nilInt *int
	var nilStr *string

	cases := []struct {
		Name string
		E, A interface{}
		O    bool
	}{
		{"same pointer", &x, &x, false},
		{"different pointers same type", &x, &y, true},
		{"both typed nil same type", nilInt, nilInt, false},
		{"both typed nil different types", nilInt, nilStr, true},
		{"non-pointer", 42, 42, false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if NotSame(wt, tc.E, tc.A) != tc.O {
				t.Errorf("NotSame(%v, %v) should be %v", tc.E, tc.A, tc.O)
			}
		})
	}
}

// --- Numeric / value ---

func TestZero(t *testing.T) {
	var nilP *int
	cases := []struct {
		Name string
		V    interface{}
		O    bool
	}{
		{"nil interface", nil, true},
		{"zero int", 0, true},
		{"non-zero int", 42, false},
		{"empty string", "", true},
		{"non-empty string", "hi", false},
		{"zero struct", SomeStruct{}, true},
		{"non-zero struct", SomeStruct{Value: 42}, false},
		{"typed nil pointer", nilP, true},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if Zero(wt, tc.V) != tc.O {
				t.Errorf("Zero(%v) should be %v", tc.V, tc.O)
			}
		})
	}
}

func TestNotZero(t *testing.T) {
	var nilP *int
	cases := []struct {
		Name string
		V    interface{}
		O    bool
	}{
		{"nil interface", nil, false},
		{"zero int", 0, false},
		{"non-zero int", 42, true},
		{"zero struct", SomeStruct{}, false},
		{"non-zero struct", SomeStruct{Value: 42}, true},
		{"typed nil pointer", nilP, false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if NotZero(wt, tc.V) != tc.O {
				t.Errorf("NotZero(%v) should be %v", tc.V, tc.O)
			}
		})
	}
}

func TestInDelta(t *testing.T) {
	cases := []struct {
		Name             string
		Expected, Actual float64
		Delta            float64
		O                bool
	}{
		{"exact", 1.0, 1.0, 0.0, true},
		{"within delta", 1.0, 1.5, 0.5, true},
		{"on boundary", 1.0, 1.5, 0.5, true},
		{"exceeds delta", 1.0, 1.5, 0.4, false},
		{"negative diff within delta", 1.0, 0.5, 0.5, true},
		{"NaN expected", math.NaN(), 1.0, 1.0, false},
		{"NaN actual", 1.0, math.NaN(), 1.0, false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if InDelta(wt, tc.Expected, tc.Actual, tc.Delta) != tc.O {
				t.Errorf("InDelta(%v, %v, %v) should be %v", tc.Expected, tc.Actual, tc.Delta, tc.O)
			}
		})
	}
}

// --- Regex ---

func TestRegexp(t *testing.T) {
	rx := regexp.MustCompile(`\d+`)
	cases := []struct {
		Name    string
		Pattern interface{}
		Str     string
		O       bool
	}{
		{"string match", "foo.*", "foobar", true},
		{"string no match", "^bar", "foobar", false},
		{"compiled match", rx, "abc123", true},
		{"compiled no match", rx, "abcdef", false},
		{"invalid pattern", "[invalid", "x", false},
		{"wrong pattern type", 42, "x", false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if Regexp(wt, tc.Pattern, tc.Str) != tc.O {
				t.Errorf("Regexp(%v, %q) should be %v", tc.Pattern, tc.Str, tc.O)
			}
		})
	}
}

func TestNotRegexp(t *testing.T) {
	cases := []struct {
		Name    string
		Pattern interface{}
		Str     string
		O       bool
	}{
		{"string match", "foo.*", "foobar", false},
		{"string no match", "^bar", "foobar", true},
		{"invalid pattern", "[invalid", "x", false},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wt := wrap(t)
			if NotRegexp(wt, tc.Pattern, tc.Str) != tc.O {
				t.Errorf("NotRegexp(%v, %q) should be %v", tc.Pattern, tc.Str, tc.O)
			}
		})
	}
}

// --- Panic ---

func TestNotPanics(t *testing.T) {
	t.Run("non-panicking", func(t *testing.T) {
		wt := wrap(t)
		if !NotPanics(wt, func() {}) {
			t.Errorf("NotPanics should pass for non-panicking func")
		}
	})
	t.Run("panicking", func(t *testing.T) {
		wt := wrap(t)
		if NotPanics(wt, func() { panic("boom") }) {
			t.Errorf("NotPanics should fail for panicking func")
		}
	})
}
