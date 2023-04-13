package buffer_write

import (
	"bytes"
	"math/rand"
	"sync"
	"time"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func generateRandomData(size int) []byte {
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = byte(rand.Intn(256))
	}
	return data
}

func processData(data []byte) {
	var buffer bytes.Buffer
	for _, b := range data {
		buffer.WriteByte(b)
	}
}

func processDataOptimized(data []byte) {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)

	for _, b := range data {
		buffer.WriteByte(b)
	}
}
