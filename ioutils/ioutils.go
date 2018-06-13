package ioutils

import (
	"io"
	"encoding/binary"
	"bytes"
	"hash/adler32"
	"errors"
)

type Packet struct {
	TotalSize uint32
	Magic     [4]byte
	Payload   []byte
	Checksum  uint32
}

var RPC_MAGIC = [4]byte{'p', 'y', 'x', 'i'}

/*
为了计算checksum，我们使用了一个内存buffer来缓存数据，最后把所有的数据一次性读出来算checksum，
考虑到计算checksum是一个不断update地过程，我们应该有方法直接略过内存buffer而计算checksum。
*/
func EncodePacket1(w io.Writer, payload []byte) error {
	totalsize := uint32(len(payload) + 8)
	binary.Write(w, binary.BigEndian, totalsize)
	binary.Write(w, binary.BigEndian, RPC_MAGIC)
	w.Write(payload)

	var buf bytes.Buffer
	buf.Write(RPC_MAGIC[:])
	buf.Write(payload)
	checksum := adler32.Checksum(buf.Bytes())

	return binary.Write(w, binary.BigEndian, checksum)
}

/*
这是一个通用的计算hash的接口，标准库里面所有计算hash的对象都实现了这个接口，比如md5, crc32等。
由于Hash实现了io.Writer接口，因此我们可以把所有要计算的数据像写入文件一样写入到这个对象中，
最后调用Sum(nil)就可以得到最终的hash的byte数组。利用这个思路，第二版可以这样写:

注意这次的变化，前面写入TotalSize，Magic，Payload部分没有变化，
在计算checksum的时候去掉了bytes.Buffer，减少了一次内存申请和拷贝。
*/
func EncodePacket2(w io.Writer, payload []byte) error {
	totalsize := uint32(len(payload) + 8)
	binary.Write(w, binary.BigEndian, totalsize)
	binary.Write(w, binary.BigEndian, RPC_MAGIC)
	w.Write(payload)

	sum := adler32.New()
	sum.Write(RPC_MAGIC[:])
	sum.Write(payload)
	checksum := sum.Sum32()

	return binary.Write(w, binary.BigEndian, checksum)
}

/*
考虑到sum和w都是io.Writer，利用神奇的io.MultiWriter，我们可以这样写
*/
func EncodePacket3(w io.Writer, payload []byte) error {
	totalsize := uint32(len(payload) + 8)
	binary.Write(w, binary.BigEndian, totalsize)

	sum := adler32.New()
	ww := io.MultiWriter(w, sum)

	// write magic bytes
	binary.Write(ww, binary.BigEndian, RPC_MAGIC)

	// write payload
	ww.Write(payload)

	// calculate checksum
	checksum := sum.Sum32()

	// write checksum
	return binary.Write(w, binary.BigEndian, checksum)
}

/*
我们使用了io.TeeReader，这个函数的原型为func TeeReader(r Reader, w Writer) Reader，
它返回一个Reader，这个Reader是参数r的代理，读取的数据还是来自r，不过同时把读取的数据写入到w里面
*/
func DecodePacket(r io.Reader) ([]byte, error) {
	var totalsize uint32
	err := binary.Read(r, binary.BigEndian, &totalsize)
	if err != nil {
		return nil, err
	}

	if totalsize < 8 {
		return nil, errors.New("bad packet. header")
	}

	sum := adler32.New()
	rr := io.TeeReader(r, sum)

	var magic [4]byte
	err = binary.Read(rr, binary.BigEndian, &magic)

	if err != nil {
		return nil, err
	}

	if magic != RPC_MAGIC {
		return nil, errors.New("bad rpc magic")
	}

	payload := make([]byte, totalsize-8)
	_, err = io.ReadFull(rr, payload)
	if err != nil {
		return nil, errors.New("read payload")
	}

	var checksum uint32
	err = binary.Read(r, binary.BigEndian, &checksum)
	if err != nil {
		return nil, errors.New("read checksum")
	}

	if checksum != sum.Sum32() {
		return nil, errors.New("checkSum error")
	}

	return payload, nil
}


