package race_condition

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	// x diakses oleh go routine secara pararel oleh thread
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// sync between Lock() and Unlock()
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(x)
}
