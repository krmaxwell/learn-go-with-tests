package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Kyle"},
			[]string{"Kyle"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Kyle", "Dallas"},
			[]string{"Kyle", "Dallas"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Kyle", 45},
			[]string{"Kyle"},
		},
		{
			"Nested fields",
			Person{
				"Kyle",
				Profile{45, "Dallas"},
			},
			[]string{"Kyle", "Dallas"},
		},
		{
			"Pointers to things",
			&Person{
				"Kyle",
				Profile{45, "Dallas"},
			},
			[]string{"Kyle", "Dallas"},
		},
		{
			"Slices",
			[]Profile{
				{45, "Dallas"},
				{50, "Monterrey"},
			},
			[]string{"Dallas", "Monterrey"},
		},
		{
			"Arrays",
			[2]Profile{
				{45, "Dallas"},
				{50, "Monterrey"},
			},
			[]string{"Dallas", "Monterrey"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
