package esp32fw

import (
	"bufio"
	"bytes"
	"io"
)

func padWith(buffer io.Writer, item byte, length uint) error {
	bytes := make([]byte, length)
	var i uint

	for i = 0; i < length; i++ {
		bytes[i] = item
	}

	_, err := buffer.Write(bytes)

	return err
}

func writeToBuffer(buffer *bytes.Buffer, reader *bufio.Reader) (uint, error) {
	var length uint
	for {
		byteRead, err := reader.ReadByte()
		if err == io.EOF {
			return length, nil
		} else if err != nil && err != io.EOF {
			return length, err
		}

		err = buffer.WriteByte(byteRead)
		if err != nil {
			return length, err
		}

		length++
	}
}
