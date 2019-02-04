package portparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//GetArrayPort parse string to array port
func GetArrayPort(s string) ([]int, error) {
	result := []int{}

	if len(s) == 0 {
		return result, errors.New("input parameter s is empty")
	}

	//remove empty space
	s = strings.Replace(s, " ", "", -1)

	portByComma := strings.Split(s, ",")
	for _, port := range portByComma {

		ports, err := getArrayPortByHyphen(port)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result = append(result, ports...)
	}

	return result, nil
}

//getArrayportByHyphen parse string to array port
func getArrayPortByHyphen(ports string) ([]int, error) {
	result := []int{}
	portByHyphen := strings.Split(ports, "-")

	if len(portByHyphen) == 0 {
		return result, errors.New("array len == 0")
	}

	if len(portByHyphen) == 1 {
		port, err := getPort(portByHyphen[0])
		if err != nil {
			fmt.Println(err)
			return result, err
		}

		return addToResult(result, port), nil
	}

	return getArrayWithIncremental(portByHyphen)
}

//getArrayWithIncremental fill port between min and max port
func getArrayWithIncremental(ports []string) ([]int, error) {
	result := []int{}

	if len(ports) < 2 {
		return result, errors.New("array ports len < 2")
	}

	first, err := getPort(ports[0])
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	last, err := getPort(ports[len(ports)-1])
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	if first == last {
		result = addToResult(result, first)
		return result, nil
	}

	result = addToResult(result, first)

	for {
		if first >= last {
			result = addToResult(result, first)
			return result, nil
		}

		first++
		if first > 0 {
			result = append(result, first)
		}
	}
}

//getPort convert string to int and validate
func getPort(s string) (int, error) {

	port, err := toInt(s)
	if err != nil {
		fmt.Println(err)
		return port, err
	}

	err = validate(port)
	if err != nil {
		fmt.Println(err)
		return port, err
	}

	return port, nil
}

//toInt string to int
func toInt(s string) (int, error) {

	port, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return port, nil
}

//validate port
func validate(port int) error {
	if port <= 0 {
		mes := fmt.Sprint("port %v <= 0", port)
		return errors.New(mes)
	}

	return nil
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
