package common

import (
	"time"
	"fmt"
)

func GetTimesFromUnix(unixs []int64) (tfs []string) {
	tfs = make([]string, len(unixs))
	var tf string
	for i, unix := range unixs {
		tf = GetTimeFromUnix(unix)
		tfs[i] = tf
	}
	return
}

func GetTimeFromUnix(unix int64) (tf string) {
	t := time.Unix(unix/1e7, 0)
	tf = t.Format("2006-01-02 15:04:05")
	return
}

func transferTime() {

	unixs := []int64{
		15088446750013028,
		9466884000000000,
		14762688284636972,
	}
	times := GetTimesFromUnix(unixs)
	for _, time := range times {
		fmt.Println(time)
	}
}