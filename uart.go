package beaglebone

import (
	"fmt"
	"io"

	"github.com/SpaceLeap/go-embedded"
	"github.com/ungerik/goserial"
)

const (
	UART_BAUD_115200 = 115200
	UART_BAUD_57600  = 57600
	UART_BAUD_38400  = 38400
	UART_BAUD_19200  = 19200
	UART_BAUD_9600   = 9600
)

type UARTNr int

const (
	UART1 UARTNr = 1
	UART2 UARTNr = 2
	UART4 UARTNr = 4
	UART5 UARTNr = 5
)

type UARTParityMode goserial.ParityMode

const (
	UART_PARITY_NONE = UARTParityMode(goserial.ParityNone)
	UART_PARITY_EVEN = UARTParityMode(goserial.ParityEven)
	UART_PARITY_ODD  = UARTParityMode(goserial.ParityOdd)
)

type UARTByteSize goserial.ByteSize

const (
	UART_BYTESIZE_5 = UARTByteSize(goserial.Byte5)
	UART_BYTESIZE_6 = UARTByteSize(goserial.Byte6)
	UART_BYTESIZE_7 = UARTByteSize(goserial.Byte7)
	UART_BYTESIZE_8 = UARTByteSize(goserial.Byte8)
)

type UARTStopBits goserial.StopBits

const (
	UART_STOPBITS_1 = UARTStopBits(goserial.StopBits1)
	UART_STOPBITS_2 = UARTStopBits(goserial.StopBits2)
)

// var uartTable = map[UARTName]uartInfo{
// 	UART1: {"UART1", "/dev/ttyO1", "ADAFRUIT-UART1", "P9_26", "P9_24"},
// 	UART2: {"UART2", "/dev/ttyO2", "ADAFRUIT-UART2", "P9_22", "P9_21"},
// 	UART4: {"UART4", "/dev/ttyO4", "ADAFRUIT-UART4", "P9_11", "P9_13"},
// 	UART5: {"UART5", "/dev/ttyO5", "ADAFRUIT-UART5", "P8_38", "P8_37"},
// }

// UART wraps "github.com/huin/goserial"
type UART struct {
	nr     UARTNr
	serial io.ReadWriteCloser
}

func NewUART(nr UARTNr, baud int, size UARTByteSize, parity UARTParityMode, stopBits UARTStopBits) (*UART, error) {
	dt := fmt.Sprintf("ADAFRUIT-UART%d", nr)
	err := embedded.LoadDeviceTree(dt)
	if err != nil {
		return nil, err
	}

	uart := &UART{nr: nr}

	config := &goserial.Config{
		Name:     fmt.Sprintf("/dev/ttyO%d", nr),
		Baud:     baud,
		Size:     goserial.ByteSize(size),
		Parity:   goserial.ParityMode(parity),
		StopBits: goserial.StopBits(stopBits),
	}
	uart.serial, err = goserial.OpenPort(config)
	if err != nil {
		return nil, err
	}

	return uart, nil
}

func (uart *UART) Read(p []byte) (n int, err error) {
	return uart.serial.Read(p)
}

func (uart *UART) Write(p []byte) (n int, err error) {
	return uart.serial.Write(p)
}

func (uart *UART) Close() error {
	err := uart.serial.Close()
	if err != nil {
		return err
	}
	return embedded.UnloadDeviceTree(fmt.Sprintf("ADAFRUIT-UART%d", uart.nr))
}
