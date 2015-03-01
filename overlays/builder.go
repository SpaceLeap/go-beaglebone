package overlays

import (
	"os/exec"
	"path"
)

func Compile(basePath string) {
	// SPI Overlays
	exec.Command("dtc", "-O", "dtb", "-o", path.Join(basePath, "overlays/ADAFRUIT-SPI0-00A0.dtbo"), "-b", "o", "-@", path.Join(basePath, "overlays/ADAFRUIT-SPI0-00A0.dts")).Run()
	exec.Command("dtc", "-O", "dtb", "-o", path.Join(basePath, "overlays/ADAFRUIT-SPI1-00A0.dtbo"), "-b", "o", "-@", path.Join(basePath, "overlays/ADAFRUIT-SPI1-00A0.dts")).Run()
	// // UART Overlayss
	exec.Command("dtc", "-O", "dtb", "-o", path.Join(basePath, "overlays/ADAFRUIT-UART1-00A0.dtbo"), "-b", "o", "-@", path.Join(basePath, "overlays/ADAFRUIT-UART1-00A0.dts")).Run()
	exec.Command("dtc", "-O", "dtb", "-o", path.Join(basePath, "overlays/ADAFRUIT-UART2-00A0.dtbo"), "-b", "o", "-@", path.Join(basePath, "overlays/ADAFRUIT-UART2-00A0.dts")).Run()
	exec.Command("dtc", "-O", "dtb", "-o", path.Join(basePath, "overlays/ADAFRUIT-UART4-00A0.dtbo"), "-b", "o", "-@", path.Join(basePath, "overlays/ADAFRUIT-UART4-00A0.dts")).Run()
	exec.Command("dtc", "-O", "dtb", "-o", path.Join(basePath, "overlays/ADAFRUIT-UART5-00A0.dtbo"), "-b", "o", "-@", path.Join(basePath, "overlays/ADAFRUIT-UART5-00A0.dts")).Run()
}

func Copy(basePath string) {
	// SPI Overlays
	exec.Command("mv", "-f", path.Join(basePath, "overlays/ADAFRUIT-SPI0-00A0.dtbo"), "/lib/firmware/ADAFRUIT-SPI0-00A0.dtbo").Run()
	exec.Command("mv", "-f", path.Join(basePath, "overlays/ADAFRUIT-SPI1-00A0.dtbo"), "/lib/firmware/ADAFRUIT-SPI1-00A0.dtbo").Run()
	// UART Overlays
	exec.Command("mv", "-f", path.Join(basePath, "overlays/ADAFRUIT-UART1-00A0.dtbo"), "/lib/firmware/ADAFRUIT-UART1-00A0.dtbo").Run()
	exec.Command("mv", "-f", path.Join(basePath, "overlays/ADAFRUIT-UART2-00A0.dtbo"), "/lib/firmware/ADAFRUIT-UART2-00A0.dtbo").Run()
	exec.Command("mv", "-f", path.Join(basePath, "overlays/ADAFRUIT-UART4-00A0.dtbo"), "/lib/firmware/ADAFRUIT-UART4-00A0.dtbo").Run()
	exec.Command("mv", "-f", path.Join(basePath, "overlays/ADAFRUIT-UART5-00A0.dtbo"), "/lib/firmware/ADAFRUIT-UART5-00A0.dtbo").Run()
}
