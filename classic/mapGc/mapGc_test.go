package map_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*type playerInfo struct {
	channelID     uint32
	playerID      uint64
	lastWorldTime int64

	level          uint32
	nick           string
	pic            string
	loginGateway   string
	loginTconnd    int32
	loginConnIndex int32
}*/

// Go里边的string 是带指针的,换成定长数组
/*type playerInfo struct {
	channelID     uint32
	playerID      uint64
	lastWorldTime int64

	level          uint32
	nick           [32]byte
	pic            [32]byte
	loginGateway   [64]byte
	loginTconnd    int32
	loginConnIndex int32
}*/

type playerInfo struct {
	playerID uint64

	nick [100]byte
	pic  [10]byte
}

const N = 10000 * 100

var globalMap interface{}

func TestWithPtrGC(t *testing.T) {
	m := make(map[uint64]*playerInfo)
	globalMap = m

	for i := uint64(0); i < N; i++ {
		m[i] = &playerInfo{}
	}

	start := time.Now()
	runtime.GC()
	during := time.Since(start)
	fmt.Printf("\nTestMapWithPtrGC %d  %v\n\n", N, during)
}

func TestWithOutPtrGC(t *testing.T) {
	m := make(map[uint64]playerInfo)
	globalMap = m

	for i := uint64(0); i < N; i++ {
		m[i] = playerInfo{
			playerID: i,
			nick:     [100]byte{},
			pic:      [10]byte{},
		}
	}

	start := time.Now()
	runtime.GC()
	during := time.Since(start)
	fmt.Printf("\nTestMapWithOutPtrGC %d  %v\n\n", N, during)
}
