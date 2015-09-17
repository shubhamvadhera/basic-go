package sleep

import "time"

//Sleep function without using the time.sleep
func Sleep(d time.Duration) {
	select {
	case <-time.After(d):
	}
}
