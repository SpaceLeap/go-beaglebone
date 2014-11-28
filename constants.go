package beaglebone

type PinInfo struct {
	Name       string
	Key        string
	GPIO       int
	PWMMuxMode int
	AInNr      int
	// PinMuxFilename string
}

var (
	USR0  = PinInfo{"USR0", "USR0", 53, -1, -1}
	USR1  = PinInfo{"USR1", "USR1", 54, -1, -1}
	USR2  = PinInfo{"USR2", "USR2", 55, -1, -1}
	USR3  = PinInfo{"USR3", "USR3", 56, -1, -1}
	P8_1  = PinInfo{"DGND", "P8_1", 0, -1, -1}
	P8_2  = PinInfo{"DGND", "P8_2", 0, -1, -1}
	P8_3  = PinInfo{"GPIO1_6", "P8_3", 38, -1, -1}
	P8_4  = PinInfo{"GPIO1_7", "P8_4", 39, -1, -1}
	P8_5  = PinInfo{"GPIO1_2", "P8_5", 34, -1, -1}
	P8_6  = PinInfo{"GPIO1_3", "P8_6", 35, -1, -1}
	P8_7  = PinInfo{"TIMER4", "P8_7", 66, -1, -1}
	P8_8  = PinInfo{"TIMER7", "P8_8", 67, -1, -1}
	P8_9  = PinInfo{"TIMER5", "P8_9", 69, -1, -1}
	P8_10 = PinInfo{"TIMER6", "P8_10", 68, -1, -1}
	P8_11 = PinInfo{"GPIO1_13", "P8_11", 45, -1, -1}
	P8_12 = PinInfo{"GPIO1_12", "P8_12", 44, -1, -1}
	P8_13 = PinInfo{"EHRPWM2B", "P8_13", 23, 4, -1}
	P8_14 = PinInfo{"GPIO0_26", "P8_14", 26, -1, -1}
	P8_15 = PinInfo{"GPIO1_15", "P8_15", 47, -1, -1}
	P8_16 = PinInfo{"GPIO1_14", "P8_16", 46, -1, -1}
	P8_17 = PinInfo{"GPIO0_27", "P8_17", 27, -1, -1}
	P8_18 = PinInfo{"GPIO2_1", "P8_18", 65, -1, -1}
	P8_19 = PinInfo{"EHRPWM2A", "P8_19", 22, 4, -1}
	P8_20 = PinInfo{"GPIO1_31", "P8_20", 63, -1, -1}
	P8_21 = PinInfo{"GPIO1_30", "P8_21", 62, -1, -1}
	P8_22 = PinInfo{"GPIO1_5", "P8_22", 37, -1, -1}
	P8_23 = PinInfo{"GPIO1_4", "P8_23", 36, -1, -1}
	P8_24 = PinInfo{"GPIO1_1", "P8_24", 33, -1, -1}
	P8_25 = PinInfo{"GPIO1_0", "P8_25", 32, -1, -1}
	P8_26 = PinInfo{"GPIO1_29", "P8_26", 61, -1, -1}
	P8_27 = PinInfo{"GPIO2_22", "P8_27", 86, -1, -1}
	P8_28 = PinInfo{"GPIO2_24", "P8_28", 88, -1, -1}
	P8_29 = PinInfo{"GPIO2_23", "P8_29", 87, -1, -1}
	P8_30 = PinInfo{"GPIO2_25", "P8_30", 89, -1, -1}
	P8_31 = PinInfo{"UART5_CTSN", "P8_31", 10, -1, -1}
	P8_32 = PinInfo{"UART5_RTSN", "P8_32", 11, -1, -1}
	P8_33 = PinInfo{"UART4_RTSN", "P8_33", 9, -1, -1}
	P8_34 = PinInfo{"UART3_RTSN", "P8_34", 81, 2, -1}
	P8_35 = PinInfo{"UART4_CTSN", "P8_35", 8, -1, -1}
	P8_36 = PinInfo{"UART3_CTSN", "P8_36", 80, 2, -1}
	P8_37 = PinInfo{"UART5_TXD", "P8_37", 78, -1, -1}
	P8_38 = PinInfo{"UART5_RXD", "P8_38", 79, -1, -1}
	P8_39 = PinInfo{"GPIO2_12", "P8_39", 76, -1, -1}
	P8_40 = PinInfo{"GPIO2_13", "P8_40", 77, -1, -1}
	P8_41 = PinInfo{"GPIO2_10", "P8_41", 74, -1, -1}
	P8_42 = PinInfo{"GPIO2_11", "P8_42", 75, -1, -1}
	P8_43 = PinInfo{"GPIO2_8", "P8_43", 72, -1, -1}
	P8_44 = PinInfo{"GPIO2_9", "P8_44", 73, -1, -1}
	P8_45 = PinInfo{"GPIO2_6", "P8_45", 70, 3, -1}
	P8_46 = PinInfo{"GPIO2_7", "P8_46", 71, 3, -1}
	P9_1  = PinInfo{"DGND", "P9_1", 0, -1, -1}
	P9_2  = PinInfo{"DGND", "P9_2", 0, -1, -1}
	P9_3  = PinInfo{"VDD_3V3", "P9_3", 0, -1, -1}
	P9_4  = PinInfo{"VDD_3V3", "P9_4", 0, -1, -1}
	P9_5  = PinInfo{"VDD_5V", "P9_5", 0, -1, -1}
	P9_6  = PinInfo{"VDD_5V", "P9_6", 0, -1, -1}
	P9_7  = PinInfo{"SYS_5V", "P9_7", 0, -1, -1}
	P9_8  = PinInfo{"SYS_5V", "P9_8", 0, -1, -1}
	P9_9  = PinInfo{"PWR_BUT", "P9_9", 0, -1, -1}
	P9_10 = PinInfo{"SYS_RESETn", "P9_10", 0, -1, -1}
	P9_11 = PinInfo{"UART4_RXD", "P9_11", 30, -1, -1}
	P9_12 = PinInfo{"GPIO1_28", "P9_12", 60, -1, -1}
	P9_13 = PinInfo{"UART4_TXD", "P9_13", 31, -1, -1}
	P9_14 = PinInfo{"EHRPWM1A", "P9_14", 50, 6, -1}
	P9_15 = PinInfo{"GPIO1_16", "P9_15", 48, -1, -1}
	P9_16 = PinInfo{"EHRPWM1B", "P9_16", 51, 6, -1}
	P9_17 = PinInfo{"I2C1_SCL", "P9_17", 5, -1, -1}
	P9_18 = PinInfo{"I2C1_SDA", "P9_18", 4, -1, -1}
	P9_19 = PinInfo{"I2C2_SCL", "P9_19", 13, -1, -1}
	P9_20 = PinInfo{"I2C2_SDA", "P9_20", 12, -1, -1}
	P9_21 = PinInfo{"UART2_TXD", "P9_21", 3, 3, -1}
	P9_22 = PinInfo{"UART2_RXD", "P9_22", 2, 3, -1}
	P9_23 = PinInfo{"GPIO1_17", "P9_23", 49, -1, -1}
	P9_24 = PinInfo{"UART1_TXD", "P9_24", 15, -1, -1}
	P9_25 = PinInfo{"GPIO3_21", "P9_25", 117, -1, -1}
	P9_26 = PinInfo{"UART1_RXD", "P9_26", 14, -1, -1}
	P9_27 = PinInfo{"GPIO3_19", "P9_27", 115, -1, -1}
	P9_28 = PinInfo{"SPI1_CS0", "P9_28", 113, 4, -1}
	P9_29 = PinInfo{"SPI1_D0", "P9_29", 111, 1, -1}
	P9_30 = PinInfo{"SPI1_D1", "P9_30", 112, -1, -1}
	P9_31 = PinInfo{"SPI1_SCLK", "P9_31", 110, 1, -1}
	P9_32 = PinInfo{"VDD_ADC", "P9_32", 0, -1, -1}
	P9_33 = PinInfo{"AIN4", "P9_33", 0, -1, 4}
	P9_34 = PinInfo{"GNDA_ADC", "P9_34", 0, -1, -1}
	P9_35 = PinInfo{"AIN6", "P9_35", 0, -1, 6}
	P9_36 = PinInfo{"AIN5", "P9_36", 0, -1, 5}
	P9_37 = PinInfo{"AIN2", "P9_37", 0, -1, 2}
	P9_38 = PinInfo{"AIN3", "P9_38", 0, -1, 3}
	P9_39 = PinInfo{"AIN0", "P9_39", 0, -1, 0}
	P9_40 = PinInfo{"AIN1", "P9_40", 0, -1, 1}
	P9_41 = PinInfo{"CLKOUT2", "P9_41", 20, -1, -1}
	P9_42 = PinInfo{"GPIO0_7", "P9_42", 7, 0, -1}
	P9_43 = PinInfo{"DGND", "P9_43", 0, -1, -1}
	P9_44 = PinInfo{"DGND", "P9_44", 0, -1, -1}
	P9_45 = PinInfo{"DGND", "P9_45", 0, -1, -1}
	P9_46 = PinInfo{"DGND", "P9_46", 0, -1, -1}
)

var (
	GPIO1_6    = &P8_3
	GPIO1_7    = &P8_4
	GPIO1_2    = &P8_5
	GPIO1_3    = &P8_6
	TIMER4     = &P8_7
	TIMER7     = &P8_8
	TIMER5     = &P8_9
	TIMER6     = &P8_10
	GPIO1_13   = &P8_11
	GPIO1_12   = &P8_12
	EHRPWM2B   = &P8_13
	GPIO0_26   = &P8_14
	GPIO1_15   = &P8_15
	GPIO1_14   = &P8_16
	GPIO0_27   = &P8_17
	GPIO2_1    = &P8_18
	EHRPWM2A   = &P8_19
	GPIO1_31   = &P8_20
	GPIO1_30   = &P8_21
	GPIO1_5    = &P8_22
	GPIO1_4    = &P8_23
	GPIO1_1    = &P8_24
	GPIO1_0    = &P8_25
	GPIO1_29   = &P8_26
	GPIO2_22   = &P8_27
	GPIO2_24   = &P8_28
	GPIO2_23   = &P8_29
	GPIO2_25   = &P8_30
	UART5_CTSN = &P8_31
	UART5_RTSN = &P8_32
	UART4_RTSN = &P8_33
	UART3_RTSN = &P8_34
	UART4_CTSN = &P8_35
	UART3_CTSN = &P8_36
	UART5_TXD  = &P8_37
	UART5_RXD  = &P8_38
	GPIO2_12   = &P8_39
	GPIO2_13   = &P8_40
	GPIO2_10   = &P8_41
	GPIO2_11   = &P8_42
	GPIO2_8    = &P8_43
	GPIO2_9    = &P8_44
	GPIO2_6    = &P8_45
	GPIO2_7    = &P8_46
	PWR_BUT    = &P9_9
	SYS_RESETn = &P9_10
	UART4_RXD  = &P9_11
	GPIO1_28   = &P9_12
	UART4_TXD  = &P9_13
	EHRPWM1A   = &P9_14
	GPIO1_16   = &P9_15
	EHRPWM1B   = &P9_16
	I2C1_SCL   = &P9_17
	I2C1_SDA   = &P9_18
	I2C2_SCL   = &P9_19
	I2C2_SDA   = &P9_20
	UART2_TXD  = &P9_21
	UART2_RXD  = &P9_22
	GPIO1_17   = &P9_23
	UART1_TXD  = &P9_24
	GPIO3_21   = &P9_25
	UART1_RXD  = &P9_26
	GPIO3_19   = &P9_27
	SPI1_CS0   = &P9_28
	SPI1_D0    = &P9_29
	SPI1_D1    = &P9_30
	SPI1_SCLK  = &P9_31
	VDD_ADC    = &P9_32
	AIN4       = &P9_33
	GNDA_ADC   = &P9_34
	AIN6       = &P9_35
	AIN5       = &P9_36
	AIN2       = &P9_37
	AIN3       = &P9_38
	AIN0       = &P9_39
	AIN1       = &P9_40
	CLKOUT2    = &P9_41
	GPIO0_7    = &P9_42
)

//Table generated based on https://raw.github.com/jadonk/bonescript/master/node_modules/bonescript/bone.js
var pinsTable = []*PinInfo{
	&USR0,
	&USR1,
	&USR2,
	&USR3,
	&P8_1,
	&P8_2,
	&P8_3,
	&P8_4,
	&P8_5,
	&P8_6,
	&P8_7,
	&P8_8,
	&P8_9,
	&P8_10,
	&P8_11,
	&P8_12,
	&P8_13,
	&P8_14,
	&P8_15,
	&P8_16,
	&P8_17,
	&P8_18,
	&P8_19,
	&P8_20,
	&P8_21,
	&P8_22,
	&P8_23,
	&P8_24,
	&P8_25,
	&P8_26,
	&P8_27,
	&P8_28,
	&P8_29,
	&P8_30,
	&P8_31,
	&P8_32,
	&P8_33,
	&P8_34,
	&P8_35,
	&P8_36,
	&P8_37,
	&P8_38,
	&P8_39,
	&P8_40,
	&P8_41,
	&P8_42,
	&P8_43,
	&P8_44,
	&P8_45,
	&P8_46,
	&P9_1,
	&P9_2,
	&P9_3,
	&P9_4,
	&P9_5,
	&P9_6,
	&P9_7,
	&P9_8,
	&P9_9,
	&P9_10,
	&P9_11,
	&P9_12,
	&P9_13,
	&P9_14,
	&P9_15,
	&P9_16,
	&P9_17,
	&P9_18,
	&P9_19,
	&P9_20,
	&P9_21,
	&P9_22,
	&P9_23,
	&P9_24,
	&P9_25,
	&P9_26,
	&P9_27,
	&P9_28,
	&P9_29,
	&P9_30,
	&P9_31,
	&P9_32,
	&P9_33,
	&P9_34,
	&P9_35,
	&P9_36,
	&P9_37,
	&P9_38,
	&P9_39,
	&P9_40,
	&P9_41,
	&P9_42,
	&P9_43,
	&P9_44,
	&P9_45,
	&P9_46,
}
