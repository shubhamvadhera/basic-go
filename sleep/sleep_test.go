package sleep

import (
	"fmt"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	type testpair struct {
		act, exp int
	}

	//number of seconds to test the sleep function
	var tests = []testpair{
		{2, 2},
		{0, 0},
		{5, 5},
		{1, 1},
	}

	for _, pair := range tests {
		fmt.Println(pair.act, "second(s) test")
		//time test - do not insert any statement after this
		tStart := time.Now()
		Sleep(time.Second * time.Duration(pair.act))
		tEnd := time.Now()
		//end of time test
		tTotal := tEnd.Sub(tStart)
		tVal := int(tTotal.Seconds())
		if tVal != pair.exp {
			t.Error(
				"For", pair.act, "second(s) test",
				"expected", pair.exp,
				"got", tVal,
			)
		}
	}

}
