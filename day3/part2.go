package main

import "regexp"

func part2() {
	input := file_to_array("./input.txt");
	var total []int;
	var tot int;
	sig := 1;
	for _, istring := range input {
		mul := find_mul_part2(istring, sig);
		mult := 0;
		for _, mu := range mul {
			// println(mul);
			multipication := give_mul(mu);
			mult += multipication;
		}
		total = append(total, mult)
	}
	for _, tol := range total {
		println(tol);
	}
	for i := 0; i < len(total); i++ {
		tot = tot + total[i]
	}
	println(tot);
}


func find_mul_part2(arr string, sig int) []string {
	pattern := `mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`;
	re := regexp.MustCompile(pattern);
	mul := re.FindAllString(arr, -1);
	// for _, ml := range mul {
	// 	println(ml)
	// }
	var rmul []string;
	for _, m := range mul {
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
