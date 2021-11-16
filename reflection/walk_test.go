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

	t.Run("with maps", func(t *testing.T) {
		myMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(myMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		myChannel := make(chan Profile)
		go func() {
			myChannel <- Profile{45, "Dallas"}
			myChannel <- Profile{50, "Monterrey"}
			close(myChannel)
		}()

		var got []string
		want := []string{"Dallas", "Monterrey"}

		walk(myChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		myFunction := func() (Profile, Profile) {
			return Profile{45, "Dallas"}, Profile{50, "Monterrey"}
		}

		var got []string
		want := []string{"Dallas", "Monterrey"}

		walk(myFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but alas", haystack, needle)
	}
}
