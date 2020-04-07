package main

import (
	"TGPL/ch12/display"
	"TGPL/ch12/encode"
	"TGPL/ch12/format"
	"fmt"
	"os"
	"reflect"
	"time"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func updateByReflect() {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // convert `value` to int pointer
	*px = 3
	fmt.Println(x)
	d.Set(reflect.ValueOf(4))
	fmt.Println(x)
	d.SetInt(5)                                 // if x's type is interface, this line will panic
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fd := stdout.FieldByName("name")
	fmt.Println(fd.String()) // "1"
	//fd.SetInt(2) // panic
	fmt.Println(fd.CanAddr(), fd.CanSet())
}

func main() {
	// 12.1
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	fmt.Println(v.CanAddr())

	// 12.2
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{x}))
	fmt.Println(format.Any([]time.Duration{d}))
	// 12.3
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb", Year: 1964,
		Color: false,
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
			"Best Adapted Screenplay (Nomin.)", "Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	display.Display("strangelove", strangelove)
	display.Display("os.Stderr", os.Stderr)
	// 12.4
	bytes, err := encode.Marshal(strangelove)
	if err == nil {
		fmt.Println((string)(bytes))
	}
	// 12.5
	updateByReflect()
}
