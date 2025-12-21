package amf

import (
	"io"
)

func ReadMarker(r io.Reader) (byte, error) {
	bytes, err := ReadBytes(r, 1)
	if err != nil {
		return 0x00, err
	}

	return bytes[0], nil
}

func ReadBytes(r io.Reader, n int) ([]byte, error) {
	buf := make([]byte, n)

	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
