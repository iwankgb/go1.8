package mutexes

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

//const SLEEP_TIME = 1 * time.Second

var wg sync.WaitGroup

//var m sync.Mutex

func TestMutexesProfiling(t *testing.T) {
	//wg.Add(30)
	wg.Add(1)
	go blocking()
	//go waiting()
	//go sleeping()
	wg.Wait()
}
func blocking() {
	//var wg sync.WaitGroup
	var m sync.Mutex

	//wg.Add(10)
	for i := 0; i < 1; i++ {
		duration := time.Duration(rand.Intn(10)) * time.Second
		m.Lock()
		go func() {
			time.Sleep(duration)
			m.Unlock()
			wg.Done()
		}()
	}
	//wg.Wait()
}

func waiting() {
	//var wg sync.WaitGroup
	var m sync.Mutex

	//wg.Add(10)
	for i := 0; i < 10; i++ {
		duration := time.Duration(rand.Intn(10)) * time.Second
		m.Lock()
		go func() {
			time.Sleep(duration)
			m.Unlock()
			wg.Done()
		}()
	}
	//wg.Wait()
}

func sleeping() {
	//var wg sync.WaitGroup
	var m sync.Mutex

	//wg.Add(10)
	for i := 0; i < 10; i++ {
		duration := time.Duration(rand.Intn(10)) * time.Second
		m.Lock()
		go func() {
			time.Sleep(duration)
			m.Unlock()
			wg.Done()
		}()
	}
	//wg.Wait()
}

//func lockRandom() {
//	duration := time.Duration(rand.Intn(10)) * time.Second
//	m.Lock()
//	go func() {
//		time.Sleep(duration)
//		m.Unlock()
//		wg.Done()
//	}()
//}
