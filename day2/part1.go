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
    input := file_to_array("./input.txt");
    count := 0;

    for _, istring := range input {
        sreport := str.Split(istring, " ");
        report := to_int(sreport);
        if is_sorted(report) {
            if is_diff_right(report) {
                count += 1;
            }
        }
    }

    println(count);
}

func is_diff_right(arr []int) bool {
    di := true;
    for i := 1; i < len(arr); i++ {
        if diff(arr[i], arr[i-1]) > 3 {
            di = false;
        }
    }

    return di;
}

func diff(num1 int, num2 int) int {
    if num1 > num2 {
        return num1 - num2
    } else {
        return num2 - num1
    }
}

func is_sorted(arr []int) bool {
    ace := true;
    dec := true;
    for i := 1; i < len(arr); i++ {
        if arr[i] <= arr[i-1] {
            ace = false;
        }
    }

    for i := 1; i < len(arr); i++ {
        if arr[i] >= arr[i-1] {
            dec = false;
        }
    }

    return ace || dec;
}

func to_int(arr []string) []int {
    var rarr []int;
    for _, num := range arr {
        n, _ := strconv.Atoi(num);
        rarr = append(rarr, n); 
    }
    return rarr;
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

