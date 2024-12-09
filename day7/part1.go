package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1();
}

func part1() {
	input := file_to_array("./test.txt");
	input_map := make(map[int][]int);
	total := 0;
	for _, element := range input {
		sp := strings.Split(element, ":");
		l_n, _ := strconv.Atoi(sp[0]);
		r_ns := strings.Split(sp[1], " ");
		var r_nnn []int;
		for i := 1; i < len(r_ns); i++ {
			r_n, _ := strconv.Atoi(r_ns[i]);
			r_nnn = append(r_nnn, r_n);
		}
		input_map[l_n] = r_nnn;
	}
	for key, value := range input_map {
		fmt.Println(key, value);
		if is_valid(int(math.Pow(2, float64(len(value)))), key, value) {
			total += key;
		}
	}
	fmt.Println(total);
}

func is_valid(posi int, key int, input_map_arr[]int) bool {
	if len(input_map_arr) < 3 {
		if key == input_map_arr[0]*input_map_arr[1] || key == input_map_arr[0]+input_map_arr[1] {
			fmt.Println("");
			return true;
		} else {
			fmt.Println("");
			return false;
		}
	}
	if posi > 0 {
		mul := input_map_arr[0]*input_map_arr[1];
		add := input_map_arr[0]+input_map_arr[1];
		mul_sub_arr := append([]int{mul}, input_map_arr[2:]...);
		fmt.Println(mul_sub_arr);
		add_sub_arr := append([]int{add}, input_map_arr[2:]...);
		return is_valid(posi/2, key, mul_sub_arr) || is_valid(posi/2, key, add_sub_arr);
	} else {
		println("why");
		return false;
	}
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
