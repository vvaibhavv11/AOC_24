package main

import "strings"

func part2() {
	input := file_to_array("./input.txt");
	var input_matrix [][]rune = make([][]rune, len(input));
	count := 0;
	for idx, istring := range input {
		for _, w := range istring {
			input_matrix[idx] = append(input_matrix[idx], w);
		}
	}
	for i := 1; i < len(input_matrix) - 1; i++ {
		for j := 1; j < len(input_matrix[i]) - 1; j++ {
			if input_matrix[i][j] == 'A' {
				count += find_all_xmas_part2(i, j, input_matrix);
			}
		}
	}
	print(count);
}

func find_all_xmas_part2(i int, j int, arr [][]rune) int {
	count := 0;
	var all_str []string;
	var sb1 strings.Builder;
	sb1.WriteRune(arr[i-1][j-1]);
	sb1.WriteRune(arr[i][j]);
	sb1.WriteRune(arr[i+1][j+1]);
	all_str = append(all_str, sb1.String());
	var sb2 strings.Builder;
	sb2.WriteRune(arr[i-1][j+1]);
	sb2.WriteRune(arr[i][j]);
	sb2.WriteRune(arr[i+1][j-1]);
	all_str = append(all_str, sb2.String());
	if all_str[0] == "MAS" || all_str[0] == "SAM" {
		if all_str[1] == "MAS" || all_str[1] == "SAM" {
			count = 1;
		}
	}
	return count;
}
