package ipparse

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
)

//GetArrayIP parse string to array ip
func GetArrayIP(s string) ([]string, error) {
	result := []string{}

	if len(s) == 0 {
		return result, errors.New("input parameter s is empty")
	}

	//remove empty space
	s = strings.Replace(s, " ", "", -1)

	ipByComma := strings.Split(s, ",")
	for _, ip := range ipByComma {
		ips, err := getArray(ip)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, res := range ips {
			result = addToResult(result, res)
		}
	}

	return result, nil
}

//getArray parse string to array ip
func getArray(s string) ([]string, error) {
	result := []string{}
	ips := strings.Split(s, "-")

	if len(ips) == 0 {
		return result, errors.New("array len == 0")
	}

	if len(ips) == 1 {
		result = addToResult(result, ips[0])
		return result, nil
	}

	return getArrayWithIncremental(ips)
}

//getArrayWithIncremental fill ip between min and max ip
func getArrayWithIncremental(ips []string) ([]string, error) {
	result := []string{}

	if len(ips) < 2 {
		return result, errors.New("array ips len < 2")
	}

	first := ips[0]
	last := ips[len(ips)-1]

	if first == last {
		result = addToResult(result, first)
		return result, nil
	}

	result = addToResult(result, first)

	for {
		res, err := equalOrGreater(first, last)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//exit from loop
		if res {
			result = addToResult(result, first)
			return result, nil
		}

		first, err = incrementalIP(first)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result = addToResult(result, first)
	}
}

//incrementalIP incremental ip, eg 127.0.0.1 -> 127.0.0.2
// or 127.0.0.255 -> 127.0.1.0
func incrementalIP(x string) (string, error) {
	ip := net.ParseIP(x)

	ip = ip.To4()
	if ip == nil {
		return x, errors.New(fmt.Sprintln("%s non ipv4 address", x))
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
		return x, errors.New(fmt.Sprintln("%s non ipv4 address", ip.String()))
	}

	return ip.String(), nil
}

//equalOrGreater compare ip where x > y
func equalOrGreater(x string, y string) (bool, error) {
	ip1 := net.ParseIP(x)
	ip1 = ip1.To4()
	if ip1 == nil {
		return false, errors.New("x non ipv4 address")
	}

	ip2 := net.ParseIP(y)
	ip2 = ip2.To4()
	if ip2 == nil {
		return false, errors.New("y non ipv4 address")
	}

	res := bytes.Compare(ip1, ip2) >= 0
	return res, nil
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
