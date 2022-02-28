package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func tutorial() {
	fmt.Println("[~] First argument should be range in 40 and 100")
	fmt.Println("[~] Example:")
	fmt.Println(os.Args[0], "40")
	fmt.Println(os.Args[0], "60")
	fmt.Println(os.Args[0], "100")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			tutorial()
		}
	}()

	param := os.Args[1]

	match, err := regexp.MatchString("^([0-9]{2,3})$", param)
	if err != nil {
		fmt.Println("[!] Limit should be between 40-100")
		os.Exit(1)
	}

	if match {
		control, _ := strconv.Atoi(param)
		if control >= 40 && control <= 100 {
			arg := []byte("40 " + param)
			err := os.WriteFile("/sys/devices/platform/huawei-wmi/charge_control_thresholds", arg, 0644)
			if err != nil {
				fmt.Println("[!] Write error :", err)
				os.Exit(127)
			} else {
				fmt.Println("[+] New charge limit :", param)
			}
		} else {
			fmt.Println("[!] Limit should be between 40-100")
			os.Exit(1)
		}
	} else {
		tutorial()
		os.Exit(1)
	}

}
