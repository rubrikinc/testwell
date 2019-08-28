package tests

import (
	"fmt"
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
		{42, false},
		{"test", false},
		{"", false},
		{(*string)(nil), false},
		{&[]string{"test"}[0], true},
		{[]string{"test"}, true},
		{[]string{}, true},
		{&[]string{}, true},
		{[]string(nil), false},
		{SomeStruct{42}, false},
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
