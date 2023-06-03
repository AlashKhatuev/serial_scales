package utils

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/howeyc/crc16"
)

var (
	ErrWrongHeader = errors.New("Wrong header")
	ErrWrongCRC    = errors.New("Wrong CRC")
)

type CommonMessage struct {
	Header [3]byte
	Len    int16
	Data   []byte
	CRC    uint16
}

func NewCommonMessage(data []byte) *CommonMessage {
	return &CommonMessage{
		Header: [3]byte{0xF8, 0x55, 0xCE},
		Len:    int16(len(data)),
		Data:   data,
		CRC:    crc16.ChecksumCCITT(data),
	}
}

func CommonMessageToBytes(cm *CommonMessage) ([]byte, error) {
	var b bytes.Buffer
	b.Write(cm.Header[:])

	lenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lenBytes, uint16(cm.Len))
	b.Write(lenBytes)

	b.Write(cm.Data)
	CRCBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(CRCBytes, cm.CRC)
	b.Write(CRCBytes)

	return b.Bytes(), nil
}

func BytesToCommonMessage(message []byte) (*CommonMessage, error) {
	cm := new(CommonMessage)
	z := bytes.NewBuffer(message)

	lenBytes := make([]byte, 2)
	CRCBytes := make([]byte, 2)

	z.Read(cm.Header[:])
	if cm.Header != [3]byte{0xF8, 0x55, 0xCE} {
		return nil, ErrWrongHeader
	}
	z.Read(lenBytes)

	cm.Len = int16(binary.BigEndian.Uint16(lenBytes))
	cm.Data = make([]byte, cm.Len)

	z.Read(cm.Data)
	z.Read(CRCBytes)

	cm.CRC = binary.BigEndian.Uint16(CRCBytes)
	if crc16.ChecksumCCITT(cm.Data) != cm.CRC {
		return nil, ErrWrongCRC
	}
	return cm, nil
}
