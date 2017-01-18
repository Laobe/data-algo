package enum

import (
	"fmt"
	"strings"
)

// 炸弹人问题
//
// 13 13
// #############
// #GG.GGG#GGG.#
// ###.#G#G#G#G#
// #.......#..G#
// #G#.###.#G#G#
// #GG.GGG.#.GG#
// #G#.#G#.#.#G#
// ##G...G.....#
// #G#.#G###.#G#
// #...G#GGG.GG#
// #G#.#G#G#.#G#
// #GG.GGG#G.GG#
// #############

func Enum2() {
	var i, j int
	var n, m, sum int
	var x, y int
	var p, q int
	var result int = 0
	fmt.Scanf("%d %d", &n, &m)
	a := read_2d_string()
	for i = 0; i <= n-1; i++ {
		for j = 0; j <= m-1; j++ {
			if a[i][j] == "." {
				sum = 0
				x = i
				y = j
				for a[x][y] != "#" {
					if a[x][y] == "G" {
						sum++
					}
					x--
				}
				x = i
				y = j
				for a[x][y] != "#" {
					if a[x][y] == "G" {
						sum++
					}
					x++
				}
				x = i
				y = j
				for a[x][y] != "#" {
					if a[x][y] == "G" {
						sum++
					}
					y--
				}
				x = i
				y = j
				for a[x][y] != "#" {
					if a[x][y] == "G" {
						sum++
					}
					y++
				}

				if sum > result {
					result = sum
					p = i
					q = j
				}
			}
		}
	}
	fmt.Printf("将炸弹放置在(%d, %d), 最懂可以消灭%d个敌人\n", p, q, result)
}

func read_2d_string() [][]string {
	var tmp string
	var line []string
	var input [][]string
	for {
		_, err := fmt.Scanf("%s", &tmp)
		if err != nil {
			break
		}
		line = strings.Split(tmp, "")
		input = append(input, line)
	}
	return input
}
