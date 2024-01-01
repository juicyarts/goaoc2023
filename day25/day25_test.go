package day25

import "testing"

var testInput = []string{
	"jqt: rhn xhk nvd",
	"rsh: frs pzl lsr",
	"xhk: hfx",
	"cmg: qnr nvd lhk bvb",
	"rhn: xhk bvb hfx",
	"bvb: xhk hfx",
	"pzl: lsr hfx nvd",
	"qnr: nvd",
	"ntq: jqt hfx bvb xhk",
	"nvd: lhk",
	"lsr: lhk",
	"rzs: qnr cmg lsr rsh",
	"frs: qnr lhk lsr",
}

func TestEz(t *testing.T) {
	expected := 54

	actual := ReadInput(testInput)

	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
