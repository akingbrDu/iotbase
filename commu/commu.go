package commu

import (
	"errors"
	"fmt"
	"github.com/tarm/serial"
	"sync"
)

var (
	cuFactoryInstance *CommunicatorFactory
	recLock           sync.Mutex
)

//const (
//	Serial    = "serial"
//	TcpClient = "tcpclient"
//	TcpServer = "tcpserver"
//)

//type ICommunicator interface {
//	Start() error
//	Close() error
//}

const (
	SerialPortName = "port"
	SeralBaudRate  = "baudrate"
	SerialDataBit  = "databit"
	SerialStopBit  = "stopbit"
	SerialParity   = "parity"
)

type CommunicatorFactory struct {
}

func (factory CommunicatorFactory) CreateTcpServerCommunicator(param map[string]interface{}) (*TcpServerCommunicator, error) {
	return nil, nil
}

func (factory CommunicatorFactory) CreateTcpClientCommunicator(param map[string]interface{}) (*TcpClientCommunicator, error) {
	return nil, nil
}

func (factory CommunicatorFactory) CreateSerialCommunicator(param map[string]interface{}) (*SerialCommunicator, error) {
	portName, ok := param[SerialPortName].(string)
	if !ok {
		fmt.Println("can't parse port field in param information")
		return nil, errors.New("can't parse port field in param information")
	}

	baudRate, ok := param[SeralBaudRate].(int)
	if !ok {
		fmt.Println("can't parse baudrate field in param information")
		return nil, errors.New("can't parse baudrate field in param information")
	}

	dataBits, ok := param[SerialDataBit].(byte)
	if !ok {
		fmt.Println("can't parse databit field in param information")
		return nil, errors.New("can't parse databit field in param information")
	}

	stopBits, ok := param[SerialStopBit].(int)
	if !ok {
		fmt.Println("can't parse stopbit field in param information")
		return nil, errors.New("can't parse stopbit field in param information")
	}

	sb, err := factory.ConvertStopBits(stopBits)
	if err != nil {
		return nil, err
	}

	parity, ok := param[SerialParity].(int)
	if !ok {
		fmt.Println("can't parse parity field in param information")
		return nil, errors.New("can't parse parity field in param information")
	}

	pr, err := factory.ConvertParity(parity)
	if err != nil {
		return nil, err
	}

	return &SerialCommunicator{
		portName: portName,
		baudRate: baudRate,
		dataBits: dataBits,
		stopBits: sb,
		parity:   pr,
	}, nil
}

func (factory CommunicatorFactory) ConvertStopBits(stopBits int) (serial.StopBits, error) {
	switch stopBits {
	case 1:
		return serial.Stop1, nil
	case 2:
		return serial.Stop2, nil
	default:
		return serial.Stop1, errors.New("unsupported stop bit setting")
	}
}

func (factory CommunicatorFactory) ConvertParity(parity int) (serial.Parity, error) {
	switch parity {
	case 0:
		return serial.ParityNone, nil
	case 1:
		return serial.ParityOdd, nil
	case 2:
		return serial.ParityEven, nil
	case 3:
		return serial.ParityMark, nil
	case 4:
		return serial.ParitySpace, nil
	default:
		return serial.ParityNone, errors.New("unsupported parity setting")
	}
}

func GetCommunicatorFactory() *CommunicatorFactory {
	recLock.Lock()
	defer recLock.Unlock()

	if cuFactoryInstance == nil {
		cuFactoryInstance = &CommunicatorFactory{}
	}
	return cuFactoryInstance
}
