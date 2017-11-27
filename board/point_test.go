package board

import (
	"reflect"
	"testing"
)

func TestPointVertex(t *testing.T) {
	cases := []struct {
		in   Point
		want int
	}{
		{Point{0, 0, 10}, 0},
		{Point{1, 2, 10}, 12},
		{Point{2, 1, 10}, 21},
		{Point{9, 9, 10}, 99},
	}

	for _, c := range cases {
		got := c.in.Vertex()
		if got != c.want {
			t.Errorf("Vertex(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestPointNeighbors(t *testing.T) {
	cases := []struct {
		in   Point
		want []Point
	}{
		{
			Point{1, 1, 3}, []Point{
				Point{0, 1, 3},
				Point{0, 0, 3},
				Point{0, 2, 3},

				Point{1, 0, 3},
				Point{1, 2, 3},

				Point{2, 1, 3},
				Point{2, 0, 3},
				Point{2, 2, 3},
			},
		},
		{
			Point{0, 0, 3}, []Point{
				Point{0, 1, 3},
				Point{1, 0, 3},
				Point{1, 1, 3},
			},
		},
	}

	for _, c := range cases {
		got := c.in.Neighbors()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Neighbors(%q) == %v, want %v",
				c.in.String(),
				got,
				c.want)
		}
	}
}
