package sexpr

import (
	"reflect"
	"testing"
)

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

	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
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
	want := "(\"[]int\" (1 2 3))"

	b, err := Marshal(&x)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(b); got != want {
		t.Errorf("invalid value. got %q, want %q", got, want)
	}
}

func TestUnmarchalBool(t *testing.T) {
	in := "t"
	want := true

	var got bool
	Unmarshal([]byte(in), &got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("invalid value. got %v, want %v", got, want)
	}
}

func TestUnmarchalFloat(t *testing.T) {
	in := "1.5"
	want := 1.5

	var got float64
	Unmarshal([]byte(in), &got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("invalid value. got %v, want %v", got, want)
	}
}

func TestUnmarchalInterface(t *testing.T) {
	in := "(\"[]int\" (1 2 3))"
	want := []int{1, 2, 3}

	var got interface{}
	Unmarshal([]byte(in), &got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("invalid value. got %v, want %v", got, want)
	}
}
