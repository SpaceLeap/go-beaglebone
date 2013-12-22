package beaglebone

import (
	"fmt"

	"github.com/SpaceLeap/go-embedded"
	"github.com/SpaceLeap/go-embedded/adc"
	"github.com/SpaceLeap/go-embedded/pwm"
	"github.com/SpaceLeap/go-embedded/spi"
)

func init() {
	err := embedded.Init("bone_capemgr")
	if err != nil {
		panic(err)
	}

	spi.Init("ADAFRUIT-SPI")

	err = pwm.Init("am33xx_pwm", "bone_pwm_")
	if err != nil {
		panic(err)
	}

	err = adc.Init("cape-bone-iio")
	if err != nil {
		panic(err)
	}
}

func PWMKey(nameOrKey string) string {
	pin, ok := PinByNameOrKey(nameOrKey)
	if !ok || pin.PWMMuxMode == -1 {
		panic(fmt.Errorf("No PWM with name or key '%s'", nameOrKey))
	}
	return pin.Key
}

func AInNameByPin(pinKey string) (adc.Name, bool) {
	pin, ok := PinByKey(pinKey)
	if !ok {
		return "", false
	}
	return adc.Name(fmt.Sprintf("AIN%d", pin.AInNr)), true
}

func GPIONr(nameOrKey string) (int, error) {
	pin, ok := PinByNameOrKey(nameOrKey)
	if !ok {
		return 0, fmt.Errorf("No GPIO with name or key '%s' found", nameOrKey)
	}
	return pin.GPIO, nil
}

type Pin struct {
	Name       string
	Key        string
	GPIO       int
	PWMMuxMode int
	AInNr      int
}

//Table generated based on https://raw.github.com/jadonk/bonescript/master/node_modules/bonescript/bone.js
var pinsTable = []Pin{
	{"USR0", "USR0", 53, -1, -1},
	{"USR1", "USR1", 54, -1, -1},
	{"USR2", "USR2", 55, -1, -1},
	{"USR3", "USR3", 56, -1, -1},
	{"DGND", "P8_1", 0, -1, -1},
	{"DGND", "P8_2", 0, -1, -1},
	{"GPIO1_6", "P8_3", 38, -1, -1},
	{"GPIO1_7", "P8_4", 39, -1, -1},
	{"GPIO1_2", "P8_5", 34, -1, -1},
	{"GPIO1_3", "P8_6", 35, -1, -1},
	{"TIMER4", "P8_7", 66, -1, -1},
	{"TIMER7", "P8_8", 67, -1, -1},
	{"TIMER5", "P8_9", 69, -1, -1},
	{"TIMER6", "P8_10", 68, -1, -1},
	{"GPIO1_13", "P8_11", 45, -1, -1},
	{"GPIO1_12", "P8_12", 44, -1, -1},
	{"EHRPWM2B", "P8_13", 23, 4, -1},
	{"GPIO0_26", "P8_14", 26, -1, -1},
	{"GPIO1_15", "P8_15", 47, -1, -1},
	{"GPIO1_14", "P8_16", 46, -1, -1},
	{"GPIO0_27", "P8_17", 27, -1, -1},
	{"GPIO2_1", "P8_18", 65, -1, -1},
	{"EHRPWM2A", "P8_19", 22, 4, -1},
	{"GPIO1_31", "P8_20", 63, -1, -1},
	{"GPIO1_30", "P8_21", 62, -1, -1},
	{"GPIO1_5", "P8_22", 37, -1, -1},
	{"GPIO1_4", "P8_23", 36, -1, -1},
	{"GPIO1_1", "P8_24", 33, -1, -1},
	{"GPIO1_0", "P8_25", 32, -1, -1},
	{"GPIO1_29", "P8_26", 61, -1, -1},
	{"GPIO2_22", "P8_27", 86, -1, -1},
	{"GPIO2_24", "P8_28", 88, -1, -1},
	{"GPIO2_23", "P8_29", 87, -1, -1},
	{"GPIO2_25", "P8_30", 89, -1, -1},
	{"UART5_CTSN", "P8_31", 10, -1, -1},
	{"UART5_RTSN", "P8_32", 11, -1, -1},
	{"UART4_RTSN", "P8_33", 9, -1, -1},
	{"UART3_RTSN", "P8_34", 81, 2, -1},
	{"UART4_CTSN", "P8_35", 8, -1, -1},
	{"UART3_CTSN", "P8_36", 80, 2, -1},
	{"UART5_TXD", "P8_37", 78, -1, -1},
	{"UART5_RXD", "P8_38", 79, -1, -1},
	{"GPIO2_12", "P8_39", 76, -1, -1},
	{"GPIO2_13", "P8_40", 77, -1, -1},
	{"GPIO2_10", "P8_41", 74, -1, -1},
	{"GPIO2_11", "P8_42", 75, -1, -1},
	{"GPIO2_8", "P8_43", 72, -1, -1},
	{"GPIO2_9", "P8_44", 73, -1, -1},
	{"GPIO2_6", "P8_45", 70, 3, -1},
	{"GPIO2_7", "P8_46", 71, 3, -1},
	{"DGND", "P9_1", 0, -1, -1},
	{"DGND", "P9_2", 0, -1, -1},
	{"VDD_3V3", "P9_3", 0, -1, -1},
	{"VDD_3V3", "P9_4", 0, -1, -1},
	{"VDD_5V", "P9_5", 0, -1, -1},
	{"VDD_5V", "P9_6", 0, -1, -1},
	{"SYS_5V", "P9_7", 0, -1, -1},
	{"SYS_5V", "P9_8", 0, -1, -1},
	{"PWR_BUT", "P9_9", 0, -1, -1},
	{"SYS_RESETn", "P9_10", 0, -1, -1},
	{"UART4_RXD", "P9_11", 30, -1, -1},
	{"GPIO1_28", "P9_12", 60, -1, -1},
	{"UART4_TXD", "P9_13", 31, -1, -1},
	{"EHRPWM1A", "P9_14", 50, 6, -1},
	{"GPIO1_16", "P9_15", 48, -1, -1},
	{"EHRPWM1B", "P9_16", 51, 6, -1},
	{"I2C1_SCL", "P9_17", 5, -1, -1},
	{"I2C1_SDA", "P9_18", 4, -1, -1},
	{"I2C2_SCL", "P9_19", 13, -1, -1},
	{"I2C2_SDA", "P9_20", 12, -1, -1},
	{"UART2_TXD", "P9_21", 3, 3, -1},
	{"UART2_RXD", "P9_22", 2, 3, -1},
	{"GPIO1_17", "P9_23", 49, -1, -1},
	{"UART1_TXD", "P9_24", 15, -1, -1},
	{"GPIO3_21", "P9_25", 117, -1, -1},
	{"UART1_RXD", "P9_26", 14, -1, -1},
	{"GPIO3_19", "P9_27", 115, -1, -1},
	{"SPI1_CS0", "P9_28", 113, 4, -1},
	{"SPI1_D0", "P9_29", 111, 1, -1},
	{"SPI1_D1", "P9_30", 112, -1, -1},
	{"SPI1_SCLK", "P9_31", 110, 1, -1},
	{"VDD_ADC", "P9_32", 0, -1, -1},
	{"AIN4", "P9_33", 0, -1, 4},
	{"GNDA_ADC", "P9_34", 0, -1, -1},
	{"AIN6", "P9_35", 0, -1, 6},
	{"AIN5", "P9_36", 0, -1, 5},
	{"AIN2", "P9_37", 0, -1, 2},
	{"AIN3", "P9_38", 0, -1, 3},
	{"AIN0", "P9_39", 0, -1, 0},
	{"AIN1", "P9_40", 0, -1, 1},
	{"CLKOUT2", "P9_41", 20, -1, -1},
	{"GPIO0_7", "P9_42", 7, 0, -1},
	{"DGND", "P9_43", 0, -1, -1},
	{"DGND", "P9_44", 0, -1, -1},
	{"DGND", "P9_45", 0, -1, -1},
	{"DGND", "P9_46", 0, -1, -1},
}

func PinByName(name string) (pins Pin, found bool) {
	for i := range pinsTable {
		if pinsTable[i].Name == name {
			return pinsTable[i], true
		}
	}
	return Pin{}, false
}

func PinByKey(key string) (pins Pin, found bool) {
	for i := range pinsTable {
		if pinsTable[i].Key == key {
			return pinsTable[i], true
		}
	}
	return Pin{}, false
}

func PinByNameOrKey(nameOrKey string) (pins Pin, found bool) {
	for i := range pinsTable {
		if pinsTable[i].Name == nameOrKey || pinsTable[i].Key == nameOrKey {
			return pinsTable[i], true
		}
	}
	return Pin{}, false
}
