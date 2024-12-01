package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	str "strings"
)

func main() {
    part1();
    part2();
}

func part1() {
    input := File_to_array("./input.txt");
    var array1 []int;
    var array2 []int;
    for _, istring := range input {
        number := str.Split(istring, "   ");
        n1, _ := strconv.Atoi(number[0]);
        array1 = append(array1, n1);
        n2, _ := strconv.Atoi(number[1]);
        array2 = append(array2, n2);
    }
    sarray1 := quickSort(array1);
    sarray2 := quickSort(array2);
    sum := 0;
    for i := 0; i < len(sarray1); i++ {
        sum = sum + diff(sarray1[i], sarray2[i]);
    }
    println(sum);
}

func diff(num1 int, num2 int) int {
    if num1 > num2 {
        return num1 - num2
    } else {
        return num2 - num1
    }
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr 
	}

	pivot := arr[0]

	var less, greater []int
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func File_to_array(file_name string) []string {
    input, err := os.Open(file_name)
    Check(err)
    defer input.Close()
    var inputString []string
    scanner := bufio.NewScanner(input)

    for scanner.Scan() {
        inputString = append(inputString, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return inputString
}

func Check(e error) {
    if e != nil {
        panic(e)
    }
}
