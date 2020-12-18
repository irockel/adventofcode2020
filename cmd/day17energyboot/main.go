// read input file input.txt and calculate active cubes
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//
// read in data, calculate the active cubes for the given
// amount of cycles.
//
func main() {
	grid := readStartCubes()

	calculateCycles(grid, 6)

	activeCubes := countActiveCubes(grid)

	fmt.Printf("Counted active cubes: %d\n", activeCubes)
}

type gridInfo struct {
	grid map[int]map[int]map[int]string

	lowerX int
	upperX int
	lowerY int
	upperY int
	lowerZ int
	upperZ int
}

//
// read values from input and return as int slices
//
func readStartCubes() gridInfo {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return gridInfo{}
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	grid := make(map[int]map[int]map[int]string)
	gridSlice := map[int]map[int]string{}
	index := 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) > 0 {
			spike := map[int]string{}
			for pos, token := range line {
				spike[pos] = string(token)
			}
			gridSlice[index] = spike
			index++
		}

		if err != nil {
			break
		}
	}
	grid[0] = gridSlice

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	var powerGridInfo gridInfo
	powerGridInfo.upperZ = 0
	powerGridInfo.upperZ = 1
	powerGridInfo.lowerX = 0
	powerGridInfo.lowerY = 0
	powerGridInfo.upperY = len(grid) - 1
	powerGridInfo.upperX = len(grid[0]) - 1

	powerGridInfo.grid = grid

	return powerGridInfo
}

func calculateCycles(powerGridInfo gridInfo, cycles int) {
	for cycle := 0; cycle < cycles; cycles++ {
		newGrid := makeNewGrid(powerGridInfo)
		grid := powerGridInfo.grid
		for i := 0; i < len(newGrid.grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				for k := 0; k < len(grid[i][j]); k++ {
					/*active := countActiveCubesAround(grid, []int{i, j, k})
					if grid[i][j][k] == "#" && (active == 2 || active == 3) {
						//newGrid
					}*/
				}
			}
		}
	}

}

func countActiveCubesAround(grid map[int]map[int]string, coord []int) int {
	fmt.Printf("checking ")
	active := 0
	for ii := coord[0] - 1; ii <= coord[0]+1; ii++ {
		for jj := coord[1] - 1; jj <= coord[1]+1; jj++ {
			for kk := coord[2] - 1; jj <= coord[2]+1; kk++ {
				fmt.Printf("{ %d, %d, %d} ", ii, jj, kk)
				if kk > 0 && kk < len(grid[jj][kk]) &&
					jj > 0 && jj < len(grid[jj]) &&
					ii > 0 && ii < len(grid) {
					if grid[ii][jj][kk] == '#' {
						active++
					}
				}
			}
		}
	}
	fmt.Printf(", active: %d\n", active)

	return active
}

func makeNewGrid(powerGridInfo gridInfo) gridInfo {
	//newSliceBefore := make(map[int]map[int]string)

	return gridInfo{}
}

func countActiveCubes(powerGridInfo gridInfo) int {
	return 0
}
