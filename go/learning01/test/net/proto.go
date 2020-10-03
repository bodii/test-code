package net

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode - encode net sticky packet message
func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)

	// write message header
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// write message entity
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode - decode net sticky packet message
func Decode(reader *bufio.Reader) (string, error) {
	// read message length
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32

	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// return buffer read byte nums
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// return real message data
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	return string(pack[4:]), nil
}
