package portparse

import (
	"fmt"
	"strconv"
	"strings"
)

//GetArrayPort parse string to array port
func GetArrayPort(s string) []int {
	result := []int{}

	//remove empty space
	s = strings.Replace(s, " ", "", -1)

	portByComma := strings.Split(s, ",")
	for _, port := range portByComma {
		ports := getArrayPortByHyphen(port)
		result = append(result, ports...)
	}

	return result
}

//getArrayportByHyphen parse string to array port
func getArrayPortByHyphen(ports string) []int {
	result := []int{}
	portByHyphen := strings.Split(ports, "-")

	if len(portByHyphen) == 0 {
		return result
	}

	if len(portByHyphen) == 1 {
		result = addToResult(result, toInt(portByHyphen[0]))
		return result
	}

	first := toInt(portByHyphen[0])
	last := toInt(portByHyphen[len(portByHyphen)-1])

	if first == last {
		result = addToResult(result, first)
		return result
	}

	result = addToResult(result, first)

	for {
		if first >= last {
			result = addToResult(result, first)
			return result
		}

		first++
		if first > 0 {
			result = append(result, first)
		}
	}
}

//toInt string to int
func toInt(s string) int {

	port, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return port
}

//addToResult add port to result, exclude alike
func addToResult(result []int, port int) []int {
	if port <= 0 || Contains(result, port) {
		return result
	}

	return append(result, port)
}

// Contains tells whether a contains x.
func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
