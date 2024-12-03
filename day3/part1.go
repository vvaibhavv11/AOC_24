package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// part1();
	part2();
}

func part1() {
	input := file_to_array("./input.txt");
	// var mul []string;
	mult := 0;
	for _, istring := range input {
		mul := find_mul(istring);
		for _, mul := range mul {
			multipication := give_mul(mul);
			mult += multipication;
		}
	}
	println(mult);
	// for _, m := range mul {
	// 	println(m);
	// }
}

func find_mul(arr string) []string {
	pattern := `mul\(\d{1,3},\d{1,3}\)`;
	re := regexp.MustCompile(pattern);
	mul := re.FindAllString(arr, -1);
	return mul;
}

func give_mul(str string) int {
	pattern := `\d{1,3},\d{1,3}`;
	re := regexp.MustCompile(pattern);
	arr := strings.Split(re.FindString(str), ",");
	n1, _ := strconv.Atoi(arr[0]);
	n2, _ := strconv.Atoi(arr[1]);
	return n1 * n2;
}

func file_to_array(file_name string) []string {
    input, err := os.Open(file_name)
    check(err)
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

func check(e error) {
    if e != nil {
        panic(e)
    }
}

