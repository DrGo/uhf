package uhf

import (
	"bytes"
	"compress/lzw"
	"io"
	"log"
)

func Compress(val []byte) ([]byte, error) {
	var buff bytes.Buffer
	wtr := lzw.NewWriter(&buff, lzw.LSB, 8)
	_, err := wtr.Write(val)
	if err != nil {
		log.Printf("failed to write to compressor: %s\n", err)
		return []byte{}, err
	}
	wtr.Close()
	return buff.Bytes(), nil
}

func Decompress(val []byte) ([]byte, error) {
	b := bytes.NewBuffer(val)
	var out bytes.Buffer

	rd := lzw.NewReader(b, lzw.LSB, 8)
	defer rd.Close()

	_, err := io.Copy(&out, rd)
	if err != nil {
		log.Printf("failed to read compressed data: %s\n", err)
		return []byte{}, err
	}

	return out.Bytes(), nil
}
