package ipparse

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

//GetArrayIP parse string to array ip
func GetArrayIP(s string) []string {
	result := []string{}

	//remove empty space
	s = strings.Replace(s, " ", "", -1)

	ipByComma := strings.Split(s, ",")
	for _, ip := range ipByComma {
		ips := getArrayIPByHyphen(ip)

		for _, res := range ips {
			result = addToResult(result, res)
		}
	}

	return result
}

//getArrayIPByHyphen parse string to array ip
func getArrayIPByHyphen(ips string) []string {
	result := []string{}
	ipByHyphen := strings.Split(ips, "-")

	if len(ipByHyphen) == 0 {
		return result
	}

	if len(ipByHyphen) == 1 {
		result = addToResult(result, ipByHyphen[0])
		return result
	}

	first := ipByHyphen[0]
	last := ipByHyphen[len(ipByHyphen)-1]

	if first == last {
		result = addToResult(result, first)
		return result
	}

	result = addToResult(result, first)

	for {
		if equalOrGreater(first, last) {
			result = addToResult(result, first)
			return result
		}

		first = incrementalIP(first)
		result = addToResult(result, first)
	}
}

//incrementalIP incremental ip, eg 127.0.0.1 -> 127.0.0.2
// or 127.0.0.255 -> 127.0.1.0
func incrementalIP(x string) string {
	ip := net.ParseIP(x)

	ip = ip.To4()
	if ip == nil {
		log.Println("non ipv4 address")
	}

	if ip[3] < 255 {
		ip[3]++
	} else if ip[2] < 255 {
		ip[2]++
		ip[3] = 0
	} else if ip[1] < 255 {
		ip[1]++
		ip[2] = 0
	} else if ip[0] == 255 {
		log.Fatalln("non ipv4 address")
	}

	fmt.Println(ip)

	return ip.String()
}

//equalOrGreater compare ip where x > y
func equalOrGreater(x string, y string) bool {
	ip1 := net.ParseIP(x)
	ip1 = ip1.To4()
	if ip1 == nil {
		log.Println("x non ipv4 address")
	}

	ip2 := net.ParseIP(y)
	ip2 = ip2.To4()
	if ip2 == nil {
		log.Println("y non ipv4 address")
	}

	return ip1.Equal(ip2) || bytes.Compare(ip1, ip2) >= 0

}

//addToResult add ip to result, exclude alike
func addToResult(result []string, ip string) []string {
	if Contains(result, ip) {
		return result
	}

	return append(result, ip)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
