package main

import (
	"bufio"
	"fmt"
	"io"
	"../numbers"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("triangle.txt")

	if err != nil {
		fmt.Println("Could not open 'triangle.txt'")
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var rows []string
	for {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		rows = append(rows, strings.Trim(row, "\n"))
	}

	nums := make([][]int, len(rows))
	start := time.Now()

	for y, row := range rows {
		cols := strings.Split(row, " ")
		nums[y] = make([]int, len(cols))

		for x, col := range cols {
			number, _ := strconv.Atoi(col)
			nums[y][x] = number
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < len(nums[i])-1; j++ {
			nums[i-1][j] += numbers.Maximum(nums[i][j], nums[i][j+1])
		}
	}

	fmt.Println(nums[0][0])
	end := time.Now()
	fmt.Printf("Calculated answer in %v.\n", end.Sub(start))
}
