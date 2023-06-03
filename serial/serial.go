package serial

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

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
	n, err := b.Write(cm.Header[:])
	if err != nil {
		return nil, err
	}
	fmt.Println("DEBUG INFO: CommonMessageBytes Header n:", n)

	lenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lenBytes, uint16(cm.Len))
	n, err = b.Write(lenBytes)
	if err != nil {
		return nil, err
	}
	fmt.Println("DEBUG INFO: CommonMessageBytes Len n:", n)
	fmt.Println(cm.Len)
	n, err = b.Write(cm.Data)
	if err != nil {
		return nil, err
	}
	fmt.Println("DEBUG INFO: CommonMessageBytes Command n:", n)

	CRCBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(CRCBytes, cm.CRC)
	n, err = b.Write(CRCBytes)
	if err != nil {
		return nil, err
	}
	fmt.Println("DEBUG INFO: CommonMessageBytes CRC n:", n)
	fmt.Println(cm.CRC)
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
