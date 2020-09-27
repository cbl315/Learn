package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"text/tabwriter"
	"time"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func testSort() {
	sort.Sort(StringSlice{})
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From The Root Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("4m36s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func sortTracks() {
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	// custom sort
	sort.Sort(customSort{tracks, func(x, y *Track) bool {

		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

/*
Practice
7.8
*/
type mulSortTable struct {
	lruTitle []string // Last Recent Used Title
	t        []*Track
	less     func(x, y *Track, lru []string) bool
}

func (x mulSortTable) Len() int           { return len(x.t) }
func (x mulSortTable) Less(i, j int) bool { return x.less(x.t[i], x.t[j], x.lruTitle) }
func (x mulSortTable) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func less(x, y *Track, lruAttr []string) bool {
	// use reflect get value of struct instance by filed name(string)
	valX := reflect.ValueOf(x)
	valY := reflect.ValueOf(y)
	for _, field := range lruAttr {
		kd := valX.Elem().FieldByName(field).Kind()
		fieldX := valX.Elem().FieldByName(field)
		fieldY := valY.Elem().FieldByName(field)
		if fieldX != fieldY {
			// Track Struct only have `int` and `string` two kinds of attributes
			if kd == reflect.Int {
				if fieldX.Int() != fieldY.Int() {
					return fieldX.Int() < fieldY.Int()
				}
			} else if kd == reflect.String {
				if fieldX.String() != fieldY.String() {
					return fieldX.String() < fieldY.String()
				}
			}
		}
	}
	return false
}

func testMultiSortTable() {
	// todo compare with sort.Stable
	lru := []string{"Year", "Title", "Artist", "Album", "Length"}

	sort.Sort(mulSortTable{
		lruTitle: lru,
		t:        tracks,
		less:     less,
	})
	printTracks(tracks)
}

/*
Practice 7.10
*/
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i <= s.Len()/2+1; i++ {
		j := s.Len() - 1 - i
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func testIsPalindrome() {
	tmp := StringSlice{"213", "23", "312", "23421"}
	tmp2 := StringSlice{"1", "2", "1"}
	tmp3 := StringSlice{"1", "2", "2", "1"}
	fmt.Println(IsPalindrome(tmp))
	fmt.Println(IsPalindrome(tmp2))
	fmt.Println(IsPalindrome(tmp3))
}
