package main

import (
	"strconv"
	str "strings"
)

func part2() {
    input := File_to_array("./input.txt");
    var array1 []int;
    var array2 []int;
    for _, istring := range input {
        number := str.Split(istring, "   ");
        n1, _ := strconv.Atoi(number[0]);
        array1 = append(array1, n1);
        n2, _ := strconv.Atoi(number[1]);
        array2 = append(array2, n2);
    }
    total := 0;
    for _, number := range array1 {
        count := iscontains(number, array2);
        total = total + (number * count);
    }
    println(total);
}

func iscontains(num int, arr []int) int {
    count := 0;
    for _, number := range arr {
        if num == number {
            count += 1;
        }
    }
    return count
}
