package large_file_reading

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"time"
)

func test01(f io.Reader) error {
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 500*1024)
		return lines
	}}

	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	slicePool := sync.Pool{New: func() interface{} {
		lines := make([]string, 100)
		return lines
	}}

	r := bufio.NewReader(f)
	var wg sync.WaitGroup
	for {
		buf := linesPool.Get().([]byte)
		n, err := r.Read(buf)
		buf = buf[:n]
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}

			return err
		}

		nextUntillNewLine, err := r.ReadBytes('\n')
		if err != io.EOF {
			buf = append(buf, nextUntillNewLine...)
		}

		wg.Add(1)
		start := time.Now()
		end := start.Add(time.Second * 1)
		go func() {
			ProcessChunk(buf, &linesPool, &stringPool, &slicePool, start, end)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func ProcessChunk(check []byte, linesPool, stringPool, slicePool *sync.Pool, start time.Time, end time.Time) {
	var wg2 sync.WaitGroup
	logs := stringPool.Get().(string)
	logs = string(check)
	linesPool.Put(check)

	logsSlice := strings.Split(logs, "\n")
	stringPool.Put(logs)
	chunkSize := 100
	n := len(logsSlice)
	noOfThread := n / chunkSize
	if n%chunkSize != 0 {
		noOfThread++
	}
	length := len(logsSlice)

	for i := 0; i < length; i += chunkSize {
		wg2.Add(1)
		go func(s int, e int) {
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}
				logParts := strings.SplitN(text, ",", 2)
				logCreationTimeString := logParts[0]
				logCreationTime, err := time.Parse("2006-01-02T15:04:05.0000Z", logCreationTimeString)
				if err != nil {
					fmt.Printf("\n Could not able to parse the time: %s for log: %v", logCreationTimeString, text)
					return
				}
				if logCreationTime.After(start) && logCreationTime.Before(end) {
					fmt.Println(text)
				}
			}
			logsSlice = nil
			wg2.Done()
		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))

	}
	wg2.Wait()
	logsSlice = nil
}
