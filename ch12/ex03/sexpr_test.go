// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestMarshalComplex(t *testing.T) {
	z := 1 + 2i
	want := "#C(1 2)" // #C(1.0 2.0)

	b, err := Marshal(z)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(b); got != want {
		t.Errorf("invalid value. got %q, want %q", got, want)
	}
}
func TestMarshalFloat(t *testing.T) {
	f := 1.5
	want := "1.5"

	b, err := Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(b); got != want {
		t.Errorf("invalid value. got %q, want %q", got, want)
	}

}

func TestMarshalBool(t *testing.T) {
	ts := []struct {
		b    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}

	for _, tc := range ts {
		b, err := Marshal(tc.b)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); got != tc.want {
			t.Errorf("invalid value. got %q, want %q", got, tc.want)
		}
	}
}
func TestMarshalInterface(t *testing.T) {
	var x interface{} = []int{1, 2, 3}
	want := `("[]int" (1 2 3))`

	b, err := Marshal(&x)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(b); got != want {
		t.Errorf("invalid value. got %q, want %q", got, want)
	}
}
