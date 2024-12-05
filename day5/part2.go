package main

import (
	"strconv"
	"strings"
	"math"
)

func part2() {
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
	valide_update := invalide_numarr(&inputprod, &inputupdate);
	for _, k := range valide_update {
		middle_idx := math.Floor(float64(len(inputupdate[k])/2));
		final := uint(middle_idx);
		sum += inputupdate[k][final];
	}
	println(sum);
}

func invalide_numarr(inputprod *map[int][]int, inputupdate *[][]int) []int {
	var valide_num []int
	for i, _ := range *inputupdate {
		if len((*inputupdate)[i]) < 2 {
			continue;
		}
		valide := is_valide_part2(&(*inputupdate)[i], inputprod);
		if valide != true {
			valide_num = append(valide_num, i);
		}
	}
	return valide_num;
}

func is_valide_part2(arr *[]int, rule *map[int][]int) bool {
	valide := true;
	for i := 0; i < len(*arr); i++ {
		for j := i + 1; j < len(*arr); j++ {
			value := (*rule)[(*arr)[j]];
			for _, prod := range value {
				if prod == (*arr)[i] {
					temp := (*arr)[i];
					(*arr)[i] = (*arr)[j];
					(*arr)[j] = temp;
					valide = false;
					continue;
				}
			}
		}
	}
	return valide;
}

