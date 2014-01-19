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

func PinByName(name string) (pinInfo *PinInfo, found bool) {
	for _, pinInfo = range pinsTable {
		if pinInfo.Name == name {
			return pinInfo, true
		}
	}
	return nil, false
}

func PinByKey(key string) (pinInfo *PinInfo, found bool) {
	for _, pinInfo = range pinsTable {
		if pinInfo.Key == key {
			return pinInfo, true
		}
	}
	return nil, false
}

func PinByNameOrKey(nameOrKey string) (pinInfo *PinInfo, found bool) {
	for _, pinInfo = range pinsTable {
		if pinInfo.Name == nameOrKey || pinInfo.Key == nameOrKey {
			return pinInfo, true
		}
	}
	return nil, false
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
