package t

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Jack", Age: 20})

func unmarsh() {
	stu := &Student{}
	json.Unmarshal(buf, stu)
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}
var data = make([]byte, 100000)

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unmarsh()
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
