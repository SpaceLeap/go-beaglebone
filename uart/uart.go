package uart

import (
	"fmt"
	"io"

	"github.com/SpaceLeap/go-embedded"
	"github.com/ungerik/goserial"
)

const (
	BAUD_115200 = 115200
	BAUD_57600  = 57600
	BAUD_38400  = 38400
	BAUD_19200  = 19200
	BAUD_9600   = 9600
)

type UARTNr int

const (
	UART1 UARTNr = 1
	UART2 UARTNr = 2
	UART4 UARTNr = 4
	UART5 UARTNr = 5
)

func (nr UARTNr) Open(baud int, size ByteSize, parity ParityMode, stopBits StopBits) (*UART, error) {
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

type ParityMode goserial.ParityMode

const (
	PARITY_NONE = ParityMode(goserial.ParityNone)
	PARITY_EVEN = ParityMode(goserial.ParityEven)
	PARITY_ODD  = ParityMode(goserial.ParityOdd)
)

type ByteSize goserial.ByteSize

const (
	BYTESIZE_5 = ByteSize(goserial.Byte5)
	BYTESIZE_6 = ByteSize(goserial.Byte6)
	BYTESIZE_7 = ByteSize(goserial.Byte7)
	BYTESIZE_8 = ByteSize(goserial.Byte8)
)

type StopBits goserial.StopBits

const (
	STOPBITS_1 = StopBits(goserial.StopBits1)
	STOPBITS_2 = StopBits(goserial.StopBits2)
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
