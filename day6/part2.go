package main

type positions struct {
	x int
	y int
}

func part2() {
	finput := file_to_array("./test.txt");
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
	for y, str := range input {
		for x, st := range str {
			if st == '#' || st == '^' {
				continue;
			}
			input[y][x] = '#';
			if replae_part2(px, py, &input) {
				count += 1;
			}
			input[y][x] = '.';
		}
	}
	println(count);
}

func replae_part2(px int, py int, arr *[][]rune) bool {
	var visited_positions []positions;
	looped := false;
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
					if position_visited(px, py-1, visited_positions) {
						directions = "none";
						looped = true;
						break;
					}
					visited_positions = append(visited_positions, positions{x: px, y: py-1});
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
					if position_visited(px+1, py, visited_positions) {
						directions = "none";
						looped = true;
						break;
					}
					visited_positions = append(visited_positions, positions{x: px+1, y: py});
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
					if position_visited(px, py+1, visited_positions) {
						directions = "none";
						looped = true;
						break;
					}
					visited_positions = append(visited_positions, positions{x: px, y: py+1});
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
					if position_visited(px-1, py, visited_positions) {
						directions = "none";
						looped = true;
						break;
					}
					visited_positions = append(visited_positions, positions{x: px-1, y: py});
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
	return looped;
}

func position_visited(px int, py int, arr []positions) bool {
	for _, position := range arr {
		if position.x == px && position.y == py {
			return true;
		}
	}
	return false;
}
