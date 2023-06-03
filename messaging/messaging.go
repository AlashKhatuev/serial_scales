package messaging

import (
	"bytes"
	"encoding/binary"
	"errors"
	"serial_scales/utils"
	"time"

	"github.com/tarm/serial"
)

var (
	// Нагрузка на весовом устройстве превышает НПВ
	ErrTooBigLoad = errors.New("Too big load on scales")
	// Весовое устройство не в режиме взвешивания
	ErrScalesNotInWheigtingMode = errors.New("Scales not in weighting mode")
	// Нет связи с модулем взвешивания
	ErrNoWeightingModuleConnection = errors.New("No connection with weighting module")
	// Установлена нагрузка на платформу при включении весового устройства
	ErrLoadWhenTurningOn = errors.New("Scales has load when turned on")
	// Весовое устройство не исправно
	ErrWeightingModuleDefective = errors.New("Weighting module out of order")
	// Установка нуля невозможна
	ErrSetZeroImpossible = errors.New("Set zero is impossible")
	// Ошибка установки тары
	ErrSetTare = errors.New("Can't set tare")
	// Неизвестная команда
	ErrNack = errors.New("Given command not supported")
)

type SerialConnection struct {
	Stream *serial.Port
}

func NewSerialConnection(name string, baud, readTimeout, size int) *SerialConnection {
	stream, err := serial.OpenPort(&serial.Config{
		Name:        name,
		Baud:        baud,
		ReadTimeout: time.Duration(readTimeout),
		Size:        byte(size),
	})
	if err != nil {
		panic(err)
	}
	return &SerialConnection{
		Stream: stream,
	}

}

func (s *SerialConnection) SetZero() error {
	message := utils.NewCommonMessage([]byte{0x72})
	messageBytes := utils.CommonMessageToBytes(message)
	_, err := s.Stream.Write(messageBytes)
	if err != nil {
		return err
	}
	responseBytes := make([]byte, 16)
	_, err = s.Stream.Read(responseBytes)
	if err != nil {
		return err
	}
	response, err := utils.BytesToCommonMessage(responseBytes)
	if err != nil {
		return err
	}
	if response.Data[0] == 0x28 {
		if response.Data[1] == 0x15 {
			return ErrSetZeroImpossible
		}
	}
	return nil
}

func (s *SerialConnection) SetTare(tare int32) error {
	var data bytes.Buffer
	data.Write([]byte{0xA3})
	tareBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(tareBytes, uint32(tare))
	data.Write(tareBytes)

	message := utils.NewCommonMessage(data.Bytes())
	messageBytes := utils.CommonMessageToBytes(message)
	_, err := s.Stream.Write(messageBytes)
	if err != nil {
		return err
	}
	responseBytes := make([]byte, 16)
	_, err = s.Stream.Read(responseBytes)
	if err != nil {
		return err
	}
	response, err := utils.BytesToCommonMessage(responseBytes)
	if err != nil {
		return err
	}
	if response.Data[0] == 0x15 {
		return ErrSetTare
	}
	if response.Data[0] == 0xF0 {
		return ErrNack
	}
	return nil
}
