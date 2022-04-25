package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

//func (ip IPAddr) String() string {
//	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
//}

func (ip IPAddr) String() string {

	result := make([]string, len(ip))
	for i, v := range ip {
		result[i] = strconv.Itoa(int(v))
	}

	return strings.Join(result, ".")
}

//func (ip IPAddr) String() string {
//	var s string
//	for i := range ip {
//		if i == 0 {
//			s += fmt.Sprint(int(ip[i]))
//		} else {
//			s += "." + fmt.Sprint(int(ip[i]))
//		}
//	}
//
//	return s
//}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {129, 0, 0, 1},
		"googledns": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%s: %s\n", name, ip)
	}
}
