package uhf

import (
	"bufio"
	"bytes"
	"compress/lzw"
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
	out := []byte{}
	rd := lzw.NewReader(b, lzw.LSB, 8)
	defer rd.Close()
	scn := bufio.NewScanner(rd)
	for scn.Scan() {
		out = append(out, scn.Bytes()...)
	}
	if scn.Err() != nil {
		log.Printf("failed to read compressed data: %s\n", scn.Err())
		return []byte{}, scn.Err()
	}
	return out, nil
}
