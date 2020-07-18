package lc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type bitree struct {
	bit []int
}

func (b *bitree) get(index int) int {
	rst := 0
	index = index + 1
	for index > 0 {
		rst += b.bit[index]
		index -= index & (-index)
	}
	return rst
}

func (b *bitree) update(index int, val int) {
	index = index + 1
	for index < len(b.bit) {
		b.bit[index] += val
		index += index & (-index)
	}
}

type NumMatrix struct {
	trees []*bitree
	data  [][]int
}

func ConstructorNumMatrix(matrix [][]int) NumMatrix {
	trees := make([]*bitree, len(matrix))
	for i := 0; i < len(matrix); i++ {
		arr := matrix[i]
		trees[i] = &bitree{
			bit: make([]int, len(arr)+1),
		}
		for j, v := range matrix[i] {
			trees[i].update(j, v)
		}
	}
	return NumMatrix{
		trees: trees,
		data:  matrix,
	}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	v := this.data[row][col]
	this.trees[row].update(col, val-v)
	this.data[row][col] = val
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for row := row1; row <= row2; row++ {
		tr := this.trees[row]
		left := 0
		if col1 > 0 {
			left = tr.get(col1 - 1)
		}
		sum += tr.get(col2) - left
	}
	return sum
}

func removeDuplicateLetters(s string) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}
	pos := 0
	for i := 0; i < len(s); i++ {
		if s[i] < s[pos] {
			pos = i
		}
		cnt[s[i]-'a']--
		if cnt[s[i]-'a'] == 0 {
			break
		}
	}
	if pos == len(s) {
		return ""
	}
	return string(s[pos]) + removeDuplicateLetters(strings.ReplaceAll(s[pos+1:], string(s[pos]), ""))
}

func shortestDistance(grid [][]int) int {
	iL := len(grid)
	jL := len(grid[0])
	var buf []int
	var tmp []int
	var seen = make([]bool, iL*jL)
	fmt.Println("id", iL, jL)
	var bfs = func(id int) []int {
		buf = buf[:0]
		tmp = tmp[:0]
		for i := 0; i < len(seen); i++ {
			seen[i] = false
		}
		rst := make([]int, iL*jL)
		buf = append(buf, id)
		level := 0
		for len(buf) > 0 {
			for _, v := range buf {
				if v < 0 || v >= iL*jL || seen[v] || (level > 0 && grid[v/jL][v%jL] != 0) {
					continue
				}
				rst[v] = level
				seen[v] = true
				if v%jL != 0 {
					tmp = append(tmp, v-1)
				}
				if v%jL != jL-1 {
					tmp = append(tmp, v+1)
				}
				tmp = append(tmp, v-jL, v+jL)
			}
			level++
			tmp, buf = buf, tmp
			tmp = tmp[:0]
		}
		return rst
	}

	var buff [][]int
	for i := 0; i < iL; i++ {
		for j := 0; j < jL; j++ {
			if grid[i][j] == 1 {
				buff = append(buff, bfs(i*jL+j))
			}
		}
	}
	rst := math.MaxInt32
	for i := 0; i < iL*jL; i++ {
		if grid[i/jL][i%jL] == 0 {
			fail := false
			sum := 0
			for t := 0; t < len(buff); t++ {
				if buff[t][i] == 0 {
					fail = true
					break
				}
				sum += buff[t][i]
			}
			if !fail && sum < rst {
				rst = sum
			}
		}
	}
	if rst != math.MaxInt32 {
		return rst
	}
	return -1
}

func generateAbbreviations(word string) []string {
	if len(word) == 0 {
		return []string{""}
	}
	var rst, buf []string
	for i := 0; i < len(word); i++ {
		c := word[i]
		if len(rst) == 0 {
			rst = append(rst, string(c), "1")
			continue
		}
		buf = buf[:0]
		for _, v := range rst {
			buf = append(buf, v+string(c))
			last := v[len(v)-1]
			if last < '0' || last > '9' {
				buf = append(buf, v+"1")
			}
			for j := len(v) - 1; j >= 0; j-- {
				if (v[j] < '0' || v[j] > '9') && (j == 0 || v[j-1] < '0' || v[j-1] > '9') {
					buf = append(buf, v[:j]+strconv.Itoa(len(v)-j+1))
				} else {
					break
				}
			}
		}
		rst, buf = buf, rst
	}
	return rst
}
