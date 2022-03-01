package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func tutorial() {
	fmt.Println("[~] First argument should be range in 4 and 100")
	fmt.Println("[~] Example:")
	fmt.Println(os.Args[0], "50")
	fmt.Println(os.Args[0], "80")
	fmt.Println(os.Args[0], "100")
}

func currentLimit() string {
	defer rec()
	currentLimitByte, err := ioutil.ReadFile("/sys/devices/platform/huawei-wmi/charge_control_thresholds")
	if err != nil {
		fmt.Println("[!] Read error :", err)
		os.Exit(127)
	}
	currentLimit := fmt.Sprintf("%s", currentLimitByte)
	currentLimit = strings.Trim(currentLimit, "\n")
	currentLimit = strings.Split(currentLimit, " ")[1]
	return currentLimit
}

func rec() {
	if r := recover(); r != nil {
		tutorial()
	}
}

func main() {

	defer rec()

	param := os.Args[1]

	match, err := regexp.MatchString("^([0-9]{2,3})$", param)
	if err != nil {
		fmt.Println("[!] Limit should be between 41-100")
		os.Exit(1)
	}

	if match {
		control, _ := strconv.Atoi(param)
		if control > 40 && control <= 100 {
			arg := []byte("40 " + param)

			current := currentLimit()

			err := os.WriteFile("/sys/devices/platform/huawei-wmi/charge_control_thresholds", arg, 0644)

			if err != nil {
				fmt.Println("[!] Write error :", err)
				os.Exit(127)

			} else {
				fmt.Printf("[+] Limit changed from %s to %s\n", current, param)
			}
		} else {
			fmt.Println("[!] Limit should be between 41-100")
			os.Exit(1)
		}
	} else {
		tutorial()
		os.Exit(1)
	}

}
