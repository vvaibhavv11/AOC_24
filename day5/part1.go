package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	part1();
	part2();
}

func part1() {
	input, update := file_to_array("./input.txt");
	sum := 0;
	inputprod := make(map[int][]int);
	var inputupdate [][]int;
	for _, i := range input {
		s := strings.Split(i, "|");
		key, _ := strconv.Atoi(s[0]);
		value, _ := strconv.Atoi(s[1]);
		_, ok := inputprod[key];
		if ok == true {
			temp := inputprod[key];
			temp = append(temp, value);
			inputprod[key] = temp;
		} else {
			temp := []int{value};
			inputprod[key] = temp;
		}
	}
	for _, i := range update {
		if len(i) > 1 {
			var s []int;
			for _, snum := range strings.Split(i, ",") {
				n1, _ := strconv.Atoi(snum);
				s = append(s, n1);
			}
			inputupdate = append(inputupdate, s);
		} else {
			n2, _ := strconv.Atoi(i);
			inum := []int{n2};
			inputupdate = append(inputupdate, inum);
		}
	}
	valide_update := valide_numarr(inputprod, inputupdate);
	for _, k := range valide_update {
		middle_idx := math.Floor(float64(len(inputupdate[k])/2));
		final := uint(middle_idx);
		sum += inputupdate[k][final];
	}
	println(sum);
}

func valide_numarr(inputprod map[int][]int, inputupdate [][]int) []int {
	var valide_num []int
	for i, _ := range inputupdate {
		if len(inputupdate[i]) < 2 {
			continue;
		}
		valide := is_valide(inputupdate[i], inputprod);
		if valide == true {
			valide_num = append(valide_num, i);
		}
	}
	return valide_num;
}

func is_valide(arr []int, rule map[int][]int) bool {
	valide := true;
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			value := rule[arr[j]];
			for _, prod := range value {
				if prod == arr[i] {
					valide = false;
					continue;
				}
			}
		}
	}
	return valide;
}

func file_to_array(file_name string) ([]string, []string) {
    input, err := os.Open(file_name)
    check(err)
    defer input.Close()
    var inputnumber []string
    var inputupdate []string
	sig := 0;
    scanner := bufio.NewScanner(input)

    for scanner.Scan() {
		if scanner.Text() == "" {
			sig = 1;
			continue;
		}
		if sig == 0 {
			inputnumber = append(inputnumber, scanner.Text())
		}
		if sig == 1 {
			inputupdate = append(inputupdate, scanner.Text())
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return inputnumber, inputupdate;
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
