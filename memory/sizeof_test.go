package memory

import (
	"reflect"
	"testing"
)

type S struct {
	a  int
	s  string
	p  *S
	m  map[int32]uint32
	u  []uint64
	ua [8]uint64
	ch chan int
	i  interface{}
}

func rSizeof(o interface{}) uint64 {
	return uint64(reflect.TypeOf(o).Size())
}
func TestSizeOf(t *testing.T) {
	esz := Sizeof(S{})
	if rsz := rSizeof(S{}); esz != rsz {
		t.Fatalf("rSizeof(S{}) != Sizeof(S{}), expected %d, got %d", rsz, esz)
	}
	s := S{s: "test"}
	if sz := Sizeof(s); sz != esz+4 {
		t.Fatalf(`Sizeof(S{s: "test"}) != Sizeof(S{}) + 4, expected %d, got %d`, esz+4, sz)
	}

	s = S{m: map[int32]uint32{1: 1}}
	if sz := Sizeof(s); sz != esz+8 /*sizeof(uint32) * 2*/ {
		t.Fatalf(`Sizeof(S{m: map[int32]uint32{1: 1}}) != Sizeof(S{}) + 8, expected %d, got %d`, esz+8, sz)
	}

	s = S{p: &s}
	if sz := Sizeof(&s); sz != esz+8 /*sizeof(uint32) * sizeof(ptr)*/ {
		t.Fatalf(`Sizeof(S{p: &s}) != Sizeof(S{}), expected %d, got %d`, esz+8, sz)
	}

	m := map[int32]S{1: S{}}
	if sz := Sizeof(m); sz != esz+12 /*sizeof(uint32) + sizeof(mapHeader)*/ {
		t.Fatalf(`Sizeof(map[int32]S{1: S{}}) != Sizeof(S{}) + 12, expected %d, got %d`, esz+12, sz)
	}

	if sz := Sizeof(S{p: &S{}}); sz != esz*2 {
		t.Fatalf(`Sizeof(S{p:&S{}}) != Sizeof(S{}) *2, expected %d, got %d`, esz*2, sz)
	}

	if sz := Sizeof([...]S{S{}}); sz != esz {
		t.Fatalf(`Sizeof([...]S{S{}}) != Sizeof(S{}), expected %d, got %d`, esz, sz)
	}

	if sz := Sizeof([]S{S{}}); sz != esz+24 {
		t.Fatalf(`Sizeof([...]S{S{}}) != Sizeof(S{}), expected %d, got %d`, esz, sz)
	}

	if sz := Sizeof("test"); sz != stringSize+4 {
		t.Fatalf(`Sizeof("test") != stringSize + 4, expected %d, got %d`, stringSize+4, sz)
	}
}
