package main

import (
	"flag"
	"fmt"
	"github.com/Karl-Gustav/flexitModbus"
	"strconv"
	"time"
)

func main() {
	registerPtr := flag.String("register", "", "A flexit register, like \"SetAirSpeed\"")
	valuePtr := flag.String("value", "", "A value to set a register to, like \"2\"")
	flag.Parse()

	if *registerPtr != "" && *valuePtr != "" {
		value, err := strconv.ParseInt(*valuePtr, 10, 64)
		if err != nil {
			panic(fmt.Errorf("value needs to be an integer"))
		}
		writeValue(*registerPtr, value)
	} else {
		inputRegisters, _ := flexitModbus.ReadAllInputRegisters()
		fmt.Println("Input registers:")
		for _, r := range inputRegisters {
			fmt.Printf("%v %v value %v min/max %v/%v bytearray %v\n", r.Address, r.Name, r.Value, r.Min, r.Max, r.ByteValue)
		}

		holdingRegisters, _ := flexitModbus.ReadAllHoldingRegisters()
		fmt.Println("Holding registers:")
		for _, r := range holdingRegisters {
			fmt.Printf("%v %v value %v min/max %v/%v bytearray %v\n", r.Address, r.Name, r.Value, r.Min, r.Max, r.ByteValue)
		}
	}
}

func writeValue(key string, value int64) {
	register, err := flexitModbus.ReadHoldingRegister(key)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	register.Value = value
	fmt.Println("Writing", register.Name, "value", register.Value)
	err = flexitModbus.WriteHoldingRegister(*register)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)
	register, err = flexitModbus.ReadHoldingRegister(key)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	if register.Value != value {
		fmt.Println("Value was not stored/written, expected", value, "got", register.Value)
	} else {
		fmt.Println("Success, expected", value, "got", register.Value)
	}
}
