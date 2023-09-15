package commu

import (
	"errors"
	"fmt"
	"github.com/tarm/serial"
)

type SerialCommunicator struct {
	portName string
	baudRate int
	dataBits byte
	stopBits serial.StopBits
	parity   serial.Parity

	io *serial.Port
}

func (s *SerialCommunicator) Connect() error {
	config := &serial.Config{
		Name:     s.portName, // 串口设备名称
		Baud:     s.baudRate, // 波特率
		Size:     s.dataBits, // 数据位
		StopBits: s.stopBits,
		Parity:   s.parity,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println("fail to open serial port with err: ", err)
		return errors.New(fmt.Sprintf("fail to open serial port with err: %v", err))
	}

	fmt.Println("success to open serial port: ", s.portName)
	s.io = port
	return nil
}

func (s *SerialCommunicator) Send(data []byte) (int, error) {
	if s.io == nil {
		fmt.Println("serial port is not opened")
		return 0, errors.New("serial port is not opened")
	}

	byteLen, err := s.io.Write(data)
	if err != nil {
		fmt.Println("fail to write serial port with err: ", err)
		return byteLen, errors.New(fmt.Sprintf("fail to write serial port with err: %v", err))
	} else {
		return byteLen, nil
	}
}

func (s *SerialCommunicator) Receive(data []byte) (int, error) {
	if s.io == nil {
		fmt.Println("serial port is not opened")
		return 0, errors.New("serial port is not opened")
	}

	byteLen, err := s.io.Read(data)
	if err != nil {
		fmt.Println("fail to read serial port with err: ", err)
		return byteLen, errors.New(fmt.Sprintf("fail to read serial port with err: %v", err))
	} else {
		return byteLen, nil
	}
}

func (s *SerialCommunicator) Close() error {
	if s.io != nil {
		err := s.io.Close()
		if err != nil {
			fmt.Println("fail to close serial port with err: ", err)
			return err
		}
	}
	return nil
}
