package main

import "regexp"

func part2() {
	input := file_to_array("./input.txt");
	total := 0;
	var comman_mul []string;
	for _, istring := range input {
		mul := find_mul_part2(istring);
		for _, m := range mul {
			comman_mul = append(comman_mul, m);
		}
	}
	comman_mul = real_mul_part2(comman_mul);
	for _, mu := range comman_mul {
		multipication := give_mul(mu);
		total += multipication;
	}
	println(total);
}

func real_mul_part2(arr []string) []string {
	sig := 1;
	var rmul []string;
	for _, m := range arr {
		if m == `don't()` {
			sig = 0;
			continue;
		}
		if m == `do()` {
			sig = 1;
			continue;
		}
		if sig == 1 {
			rmul = append(rmul, m);
		}
	}
	return rmul;
}


func find_mul_part2(arr string) []string {
	pattern := `mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`;
	re := regexp.MustCompile(pattern);
	mul := re.FindAllString(arr, -1);
	return mul;
}
