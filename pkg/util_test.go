package pkg

import (
	"errors"
	"testing"
)

func Test_ParseNumbers_EmptyString(t *testing.T) {

}
func Test_ParseNumbers_All(t *testing.T) {
	//s := "l12.5.43-2,5 10.3.3.4.5"
	//s := "l10.3.3.4.5"
	correctResults := [][]float64{
		{-3, -15, 2, -.39},

		{3},

		{-.3},

		{-1.3},

		{-.3, .3},

		{.3, -.4, .55, .5, .5, .5, -3, -3.2, -.5}}

	test_data := [][]string{
		{"l-3-15,2-.39"},
		{"l3",
			"l 3"},

		{"l -.3",
			"l-.3"},

		{"l -1.3",
			"l-1.3"},

		{"l -.3.3",
			"l -.3,.3",
			"l -.3 .3",
			"l-.3.3"},

		{"l .3 -.4 .55 .5 .5 .5,-3 -3.2 -.5",
			"l .3 -.4 .55 .5 .5 .5,-3,-3.2 -.5",
			"l .3 -.4 .55 .5 .5 .5,-3, -3.2 -.5",
			"l.3-.4.55.5.5.5,-3-3.2-.5",
			"l.3 -.4.55.5.5.5,-3-3.2-.5",
			"l.3,-.4.55.5.5.5,-3-3.2-.5",
			"l.3,-.4.55.5.5.5,-3-3.2 -.5",
			"l.3,-.4.55.5.5.5,-3,-3.2 -.5",
			"l .3-.4.55.5.5.5,-3-3.2-.5"}}

	for i := 0; i < len(correctResults); i++ {
		for j, s := range test_data[i] {
			nums, err := ParseNumbers(s)
			if err != nil {
				t.Fatal(err)
			}

			if len(nums) != len(correctResults[i]) {
				t.Log("!!!!!!!!!!!!!!!!")
				t.Log(nums)
				t.Log(test_data[i])
				t.Fatal(errors.New("Number mismatch"))
			}
			for k := 0; k < len(correctResults[i]); k++ {
				if nums[k] != correctResults[i][k] {
					t.Fatal(errors.New("Number mismatch"))
				}
			}
			t.Log("------------------------------------")
			t.Log(j, s)
			t.Log("results:", nums)
		}
	}

}

func TestPersist_RegexpNum(t *testing.T) {
	s := "mmmm12mmm"
	if !numRE.Match([]byte(s)) {
		err := errors.New("Bad")
		t.Fatal(err)
	}

}
