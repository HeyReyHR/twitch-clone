package amf

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Decoder struct{}

type Object map[string]any

func (d *Decoder) DecodeAmf0(r io.Reader) (any, error) {
	marker, err := ReadMarker(r)
	if err != nil {
		return nil, err
	}
	switch marker {
	case AMF0_NUMBER_MARKER:
		return d.DecodeAmf0Number(r)
	case AMF0_BOOLEAN_MARKER:
		return d.DecodeAmf0Boolean(r)
	case AMF0_STRING_MARKER:
		return d.DecodeAmf0String(r)
	case AMF0_OBJECT_MARKER:
		return d.DecodeAmf0Object(r)
	case AMF0_NULL_MARKER:
		return nil, nil
	case AMF0_ECMA_ARRAY_MARKER:
		return d.DecodeAmf0EcmaArray(r)
	}
	return nil, fmt.Errorf("decode amf0: unsupported type %d", marker)
}

func (d *Decoder) DecodeAmf0Number(r io.Reader) (float64, error) {
	var res float64

	err := binary.Read(r, binary.BigEndian, &res)
	if err != nil {
		return float64(0), err
	}

	return res, nil
}

func (d *Decoder) DecodeAmf0Boolean(r io.Reader) (bool, error) {
	var b byte

	bytes, err := ReadBytes(r, 1)
	if err != nil {
		return false, err
	}

	b = bytes[0]
	switch b {
	case AMF0_BOOLEAN_FALSE:
		return false, nil
	case AMF0_BOOLEAN_TRUE:
		return true, nil
	}

	return false, fmt.Errorf("decode amf0: unexpected value %v for boolean", b)
}

func (d *Decoder) DecodeAmf0String(r io.Reader) (string, error) {
	var length uint16

	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return "", fmt.Errorf("decode amf0: unable to decode string length: %s", err)
	}

	bytes := make([]byte, length)
	if bytes, err = ReadBytes(r, int(length)); err != nil {
		return "", fmt.Errorf("decode amf0: unable to decode string value: %s", err)
	}

	return string(bytes), nil
}

func (d *Decoder) DecodeAmf0Object(r io.Reader) (Object, error) {
	res := make(Object)

	for {
		key, err := d.DecodeAmf0String(r)
		if err != nil {
			return nil, err
		}
		if key == "" {
			marker, err := ReadMarker(r)
			if err != nil {
				return nil, err
			}
			if marker != AMF0_OBJECT_END_MARKER {
				return nil, fmt.Errorf("decode amf0: expected object end marker: %s", err)
			}

			break
		}
		value, err := d.DecodeAmf0(r)
		if err != nil {
			return nil, fmt.Errorf("decode amf0: unable to decode object value: %s", err)
		}

		res[key] = value
	}

	return res, nil
}

func (d *Decoder) DecodeAmf0EcmaArray(r io.Reader) (Object, error) {
	var length uint32
	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	result, err := d.DecodeAmf0Object(r)
	if err != nil {
		return nil, fmt.Errorf("decode amf0: unable to decode ecma array object: %s", err)
	}

	return result, nil
}
