package main

import (
	"bufio"
	"os"
	"log"
)

func main() {
	// part1();
	part2();
}

func part1() {
	finput := file_to_array("./input.txt");
	var input[][]rune = make([][]rune, len(finput));
	var py int;
	var px int;
	count := 0;
	for idx, istring := range finput {
		for x, w := range istring {
			input[idx] = append(input[idx], w);
			if w == '^' {
				py = idx;
				px = x;
			}
		}
	}
	replae(px, py, &input);
	for _, ele := range input {
		for _, el := range ele {
			if el == '*' {
				count += 1;
			}
		}
	}
	println(count);
}

func replae(px int, py int, arr *[][]rune) {
	directions := "top";
	(*arr)[py][px] = '*';
	for {
		if directions == "top" {
			for py >= 0 {
				if py == 0 {
					directions = "none";
					(*arr)[py][px] = '*';
					break;
				}
				if (*arr)[py-1][px] == '#'{
					(*arr)[py][px] = '*';
					directions = "right"
					break;
				}
				(*arr)[py][px] = '*';
				py--;
			}
		}
		if directions == "right" {
			for px < len((*arr)[0]) {
				if px == len((*arr)[0]) - 1 {
					directions = "none";
					(*arr)[py][px] = '*';
					break
				}
				if (*arr)[py][px+1] == '#' {
					(*arr)[py][px] = '*';
					directions = "bottom"
					break;
				}
				(*arr)[py][px] = '*';
				px++;
			}
		}
		if directions == "bottom" {
			for py < len(*arr) {
				if py == len(*arr)-1 {
					directions = "none";
					(*arr)[py][px] = '*';
					break;
				}
				if (*arr)[py+1][px] == '#' {
					// println(len((*arr)[0]));
					directions = "left"
					(*arr)[py][px] = '*';
					break;
				}
				(*arr)[py][px] = '*';
				py++;
			}
		}
		if directions == "left" {
			for px >= 0 {
				if px == 0 {
					directions = "none";
					(*arr)[py][px] = '*';
					break;
				}
				if (*arr)[py][px-1] == '#' {
					(*arr)[py][px] = '*';
					directions = "top"
					break;
				}
				(*arr)[py][px] = '*';
				px--;
			}
		}
		if directions == "none" {
			break;
		}
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

