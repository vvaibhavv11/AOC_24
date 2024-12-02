package main

import (
	str "strings"
)

func part2() {
    input := file_to_array("./input.txt");
    count := 0;
    for _, istring := range input {
        sreport := str.Split(istring, " ");
        report := to_int(sreport);
        if is_sorted(report) {
            if is_diff_right(report) {
                count += 1;
            } else {
                if is_safe(report) {
                    count += 1;
                }
            }
        } else {
            if is_safe(report) {
                count += 1;
            }
        }
    }
    println(count);
}

func is_safe(arr []int) bool {
    safe := false;
    for idx, _ := range arr {
        new_arr := remove_index(arr, idx);
        if is_sorted(new_arr) {
            if is_diff_right(new_arr) {
                safe = true;
                break;
            }
        }
    }

    return safe;
}

func remove_index(s []int, index int) []int {
    var new_arr []int;
    for idx, ele := range s {
        if idx == index {
            continue;
        }
        new_arr = append(new_arr, ele)
    }
    return new_arr;
}

