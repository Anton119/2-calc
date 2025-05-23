package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const AVG = "AVG"
const SUM = "SUM"
const MED = "MED"

func main() {

	fmt.Printf("Выберите операцию: 1)%s 2)%s 3)%s\n", AVG, SUM, MED)
	choice, err := getUserOperation()
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	fmt.Printf("Введите числа через запятую:\n")
	nums, err := getUsersNums()
	if err != nil {
		fmt.Println("ошибка", err)
		return
	}
	res := 0.0
	switch choice {
	case "1":
		res = countAvg(nums)
	case "2":
		res = countSum(nums)
	case "3":
		res = countMed(nums)
	}

	fmt.Println(res)
}

func getUserOperation() (string, error) {
	var input string
	fmt.Scan(&input)

	choice, err := strconv.Atoi(input)
	if err != nil {
		return "", errors.New("can't convert to int")
	}

	if choice > 3 || choice < 1 {
		return "", errors.New("выберите вариант от 1 до 3")
	}
	return input, nil
}

func getUsersNums() ([]int, error) {
	nums := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		for _, el := range parts {
			el = strings.TrimSpace(el)
			num, err := strconv.Atoi(el)
			if err != nil {
				return nil, errors.New("can't convert to int")
			}
			nums = append(nums, num)
		}
	}
	return nums, nil
}

func countAvg(arr []int) float64 {
	if len(arr) == 0 {
		return 0.0
	}
	avg := 0.0
	tmp := 0.0
	for _, el := range arr {
		tmp += float64(el)
	}
	avg = tmp / float64(len(arr))
	return avg
}

func countSum(arr []int) float64 {
	if len(arr) == 0 {
		return 0.0
	}
	sum := 0.0
	for _, el := range arr {
		sum += float64(el)
	}
	return sum
}

func countMed(arr []int) float64 {
	data := []float64{}

	for _, el := range arr {
		data = append(data, float64(el))
	}
	l := len(arr)
	sort.Float64s(data)
	med := 0.0

	if len(data) == 0 {
		return 0.0
	}

	if l%2 == 0 {
		med = (data[l/2] + data[(l/2)-1]) / 2
	} else {
		med = data[l/2]
	}
	return med
}
