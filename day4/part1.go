package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	part1();
	part2();
}

func part1() {
	input := file_to_array("./input.txt");
	var input_matrix [][]rune = make([][]rune, len(input));
	count := 0;
	for idx, istring := range input {
		for _, w := range istring {
			input_matrix[idx] = append(input_matrix[idx], w);
		}
	}
	for i := 0; i < len(input_matrix); i++ {
		for j := 0; j < len(input_matrix[i]); j++ {
			if input_matrix[i][j] == 'X' {
				count += find_all_xmas(i, j, input_matrix);
			}
		}
	}
	println(count);
}

func find_all_xmas(i int, j int, arr [][]rune) int {
	right := j + 3 < len(arr[0]);
	left := j - 3 >= 0;
	bottom := i + 3 < len(arr);
	top := i - 3 >= 0;
	var all_str []string;
	count := 0;
	if right {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i][j+k]);
		}
		all_str = append(all_str, sb.String())
	}
	if left {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i][j-k]);
		}
		all_str = append(all_str, sb.String())
	}
	if bottom {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i+k][j]);
		}
		all_str = append(all_str, sb.String())
	}
	if top {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i-k][j]);
		}
		all_str = append(all_str, sb.String())
	}
	if bottom && right {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i+k][j+k]);
		}
		all_str = append(all_str, sb.String())
	}
	if bottom && left {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i+k][j-k]);
		}
		all_str = append(all_str, sb.String())
	}
	if top && left {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i-k][j-k]);
		}
		all_str = append(all_str, sb.String())
	}
	if top && right {
		var sb strings.Builder;
		for k := 0; k < 4; k++ {
			sb.WriteRune(arr[i-k][j+k]);
		}
		all_str = append(all_str, sb.String())
	}
	for _, xmas := range all_str {
		if xmas == "XMAS" {
			count++
		}
	}
	return count;
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
