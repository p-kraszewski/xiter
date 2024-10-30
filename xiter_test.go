package xiter

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestMapSame(t *testing.T) {
	in := slices.Values([]int{1, 2, 3, 4, 5})
	out := Map(
		in,
		func(v int) int {
			return 2 * v
		},
	)
	outSlice := slices.Collect(out)

	if !slices.Equal(outSlice, []int{2, 4, 6, 8, 10}) {
		t.Fail()
	}
}

func TestMapOther(t *testing.T) {
	in := slices.Values([]int{1, 2, 3, 4, 5})
	out := Map(
		in,
		func(v int) string {
			return fmt.Sprintf("*%d", v)
		},
	)
	outSlice := slices.Collect(out)

	if !slices.Equal(outSlice, []string{"*1", "*2", "*3", "*4", "*5"}) {
		t.Fail()
	}
}

func TestMap2Same(t *testing.T) {
	in := slices.All([]int{1, 2, 3, 4, 5})
	out := Map2(
		in,
		func(k, v int) int {
			return k + 10*v
		},
	)
	outIter := Seq2ToSeqValues(out)
	outSlice := slices.Collect(outIter)

	if !slices.Equal(outSlice, []int{10, 21, 32, 43, 54}) {
		t.Fail()
	}
}

func TestMap2Other(t *testing.T) {
	in := slices.All([]int{1, 2, 3, 4, 5})
	out := Map2(
		in,
		func(k, v int) string {
			return fmt.Sprintf("[%d]=%d", k, v)
		},
	)
	outIter := Seq2ToSeqValues(out)
	outSlice := slices.Collect(outIter)

	if !slices.Equal(outSlice, []string{"[0]=1", "[1]=2", "[2]=3", "[3]=4", "[4]=5"}) {
		t.Fail()
	}
}

func TestMap2K(t *testing.T) {
	in := maps.All(map[int]string{1: "one", 2: "two", 3: "three"})
	out := Map2K(
		in,
		func(k int, v string) (string, int) {
			return fmt.Sprintf("[%d]", k), len(v)
		},
	)

	outMap := maps.Collect(out)

	if !maps.Equal(outMap, map[string]int{"[1]": 3, "[2]": 3, "[3]": 5}) {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	in := slices.Values([]int{1, 2, 3, 4, 5})
	out := Filter(
		in,
		func(v int) bool {
			return v%2 == 0
		},
	)
	outSlice := slices.Collect(out)

	if !slices.Equal(outSlice, []int{2, 4}) {
		t.Fail()
	}
}

func TestFilter2(t *testing.T) {
	in := slices.All([]int{1, 2, 3, 4, 5})
	out := Filter2(
		in,
		func(k, v int) bool {
			return k < 4 && v%2 != 0
		},
	)
	outIter := Seq2ToSeqValues(out)
	outSlice := slices.Collect(outIter)

	if !slices.Equal(outSlice, []int{1, 3}) {
		t.Fail()
	}
}

func TestFilterMap(t *testing.T) {
	in := slices.Values([]int{1, 2, 3, 4, 5})
	out := FilterMap(
		in,
		func(v int) (int, bool) {
			return v * 2, v%2 != 0
		},
	)
	outSlice := slices.Collect(out)

	if !slices.Equal(outSlice, []int{2, 6, 10}) {
		t.Fail()
	}
}

func TestFilterMap2(t *testing.T) {
	in := slices.All([]int{1, 2, 3, 4, 5})
	out := FilterMap2(
		in,
		func(k, v int) (int, bool) {
			return v * 2, k < 4 && v%2 != 0
		},
	)
	outIter := Seq2ToSeqValues(out)
	outSlice := slices.Collect(outIter)

	if !slices.Equal(outSlice, []int{2, 6}) {
		t.Fail()
	}
}

func TestFilterMap2K(t *testing.T) {
	in := maps.All(map[int]string{1: "one", 2: "two", 3: "three"})
	out := FilterMap2K(
		in,
		func(k int, v string) (string, int, bool) {
			return fmt.Sprintf("[%d]", k), len(v), k > 1
		},
	)

	outMap := maps.Collect(out)

	if !maps.Equal(outMap, map[string]int{"[2]": 3, "[3]": 5}) {
		t.Fail()
	}
}

func TestFold(t *testing.T) {
	in := slices.Values([]int{1, 2, 3, 4, 5})

	sum := Fold(
		in,
		0,
		func(acc int, v int) int {
			return acc + v
		},
	)

	if sum != 15 {
		t.Fail()
	}
}

func TestFold2(t *testing.T) {
	in := maps.All(map[int]bool{1: true, 2: false, 3: true, 4: false, 5: true})

	sum := Fold2(
		in,
		0,
		func(acc int, k int, v bool) int {
			if v {
				return acc + k
			} else {
				return acc
			}
		},
	)

	if sum != 9 {
		t.Fail()
	}
}
