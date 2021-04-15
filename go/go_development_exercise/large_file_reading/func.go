package large_file_reading

import (
	"bufio"
	"fmt"
	"io"
)

func large_file_reading(f io.Reader) error {

	r := bufio.NewReader(f)
	for {
		buf := make([]byte, 4*1024)
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
	}
}
