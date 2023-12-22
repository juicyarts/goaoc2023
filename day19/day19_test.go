package day19

import (
	"testing"
)

var testInput = []string{
	"px{a<2006:qkq,m>2090:A,rfg}",
	"pv{a>1716:R,A}",
	"lnx{m>1548:A,A}",
	"rfg{s<537:gd,x>2440:R,A}",
	"qs{s>3448:A,lnx}",
	"qkq{x<1416:A,crn}",
	"crn{x>2662:A,R}",
	"in{s<1351:px,qqz}",
	"qqz{s>2770:qs,m<1801:hdj,R}",
	"gd{a>3333:R,R}",
	"hdj{m>838:A,pv}",
	"",
	"{x=787,m=2655,a=1222,s=2876}",
	"{x=1679,m=44,a=2067,s=496}",
	"{x=2036,m=264,a=79,s=2244}",
	"{x=2461,m=1339,a=466,s=291}",
	"{x=2127,m=1623,a=2188,s=1013}",
}

// func TestMain(t *testing.T) {
// 	expected := 19114
// 	_, _, actual := Main(testInput)

// 	if actual != expected {
// 		t.Errorf("Expected %+v, got %+v", expected, actual)
// 	}
// }

func TestMainDistinct(t *testing.T) {
	expected := 167409079868000
	_, _, actual := Main(testInput)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

// func TestMainWithInput(t *testing.T) {
// 	godotenv.Load()
// 	expected, _ := strconv.Atoi(os.Getenv("result_1"))

// 	Input, _ := utils.ReadInputFile("input.txt")
// 	_, _, result := Main(Input)

// 	if result != expected {
// 		t.Errorf("Expected to equal %+v, got %+v", expected, result)
// 	}
// }
