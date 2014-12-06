package uart

import (
	"fmt"
	"time"

	"github.com/SpaceLeap/go-embedded"
	"github.com/ungerik/goserial"
)

type UARTNr int

const (
	UART1 UARTNr = 1
	UART2 UARTNr = 2
	UART4 UARTNr = 4
	UART5 UARTNr = 5
)

func (nr UARTNr) Open(baud serial.Baud, byteSize serial.ByteSize, parity serial.ParityMode, stopBits serial.StopBits, readTimeout time.Duration) (*UART, error) {
	dt := fmt.Sprintf("ADAFRUIT-UART%d", nr)
	err := embedded.LoadDeviceTree(dt)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("/dev/ttyO%d", nr)

	uart := &UART{Nr: nr}
	uart.Connection, err = serial.Open(name, baud, byteSize, parity, stopBits, readTimeout)
	if err != nil {
		return nil, err
	}

	return uart, nil
}

type UART struct {
	*serial.Connection
	Nr UARTNr
}

func (uart *UART) Close() error {
	err := uart.Connection.Close()
	if err != nil {
		return err
	}
	return embedded.UnloadDeviceTree(fmt.Sprintf("ADAFRUIT-UART%d", uart.Nr))
}

// var uartTable = map[UARTName]uartInfo{
// 	UART1: {"UART1", "/dev/ttyO1", "ADAFRUIT-UART1", "P9_26", "P9_24"},
// 	UART2: {"UART2", "/dev/ttyO2", "ADAFRUIT-UART2", "P9_22", "P9_21"},
// 	UART4: {"UART4", "/dev/ttyO4", "ADAFRUIT-UART4", "P9_11", "P9_13"},
// 	UART5: {"UART5", "/dev/ttyO5", "ADAFRUIT-UART5", "P8_38", "P8_37"},
// }
