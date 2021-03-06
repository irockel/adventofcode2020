// read input file input.txt and try to put the tiles together to an image
// the tiles are defined like puzzle elements
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//
// read in data, calculate the results and print the sum
//
func main() {
	tiles := readImageTiles()

	tiles = findMatchingTiles(tiles)

	edgeValues := calculateEdgeValues(tiles)

	fmt.Printf("The product of the corner ids is: %d\n", edgeValues)
}

//
// describes an image tile
//
type tileDef struct {
	id   int
	data []string

	// calculated border value of all borders
	// every border with the "normal" value
	// and the reversed value in second position
	// left, left, right, right, bottom, bottom, top, top
	left   [2]int
	top    [2]int
	right  [2]int
	bottom [2]int

	// linked tiles
	leftTileID   int
	rightTileID  int
	topTileID    int
	bottomTileID int
}

//
// read values from input and image tiles with calculated border
// values as map keys
//
func readImageTiles() map[int]tileDef {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return nil
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	tiles := map[int]tileDef{}
	currentTile := tileDef{}

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			if strings.Contains(trimmedLine, "#") || strings.Contains(trimmedLine, ".") {
				currentTile.data = append(currentTile.data, trimmedLine)
			} else if strings.Contains(trimmedLine, "Tile") {
				if currentTile.id > 0 {
					currentTile = calculateBorders(currentTile)
					tiles[currentTile.id] = currentTile
				}
				tileIDStr := strings.Split(trimmedLine, " ")[1]
				// strip ":"
				tileIDStr = tileIDStr[:len(tileIDStr)-1]
				currentTile = tileDef{}
				currentTile.id, err = strconv.Atoi(tileIDStr)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		if err != nil {
			break
		}
	}

	// add remaining tile
	if currentTile.id > 0 {
		currentTile = calculateBorders(currentTile)
		tiles[currentTile.id] = currentTile
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
	}

	return tiles
}

//
// calculate the border values of the given tile,
// from left to right and from right to left
//
func calculateBorders(tile tileDef) tileDef {
	left := ""
	right := ""
	for i := 0; i < len(tile.data); i++ {
		left = left + string(tile.data[i][0])
		right = right + string(tile.data[i][len(tile.data[i])-1])
	}

	tile.left[0], _ = getBorderValue(left)
	tile.left[1], _ = getBorderValue(reverse(left))
	tile.right[0], _ = getBorderValue(right)
	tile.right[1], _ = getBorderValue(reverse(right))
	tile.top[0], _ = getBorderValue(tile.data[0])
	tile.top[1], _ = getBorderValue(reverse(tile.data[0]))
	tile.bottom[0], _ = getBorderValue(tile.data[len(tile.data)-1])
	tile.bottom[1], _ = getBorderValue(reverse(tile.data[len(tile.data)-1]))

	return tile
}

//
// get the associated border value for the given string
//
func getBorderValue(str string) (int, error) {
	valStr := strings.ReplaceAll(str, "#", "1")
	valStr = strings.ReplaceAll(valStr, ".", "0")

	val, err := strconv.ParseInt(valStr, 2, 64)

	return int(val), err
}

//
// reverse string order for a turned or flipped tile
//
func reverse(s string) string {
	r := []rune(s)
	var output strings.Builder
	for i := len(r) - 1; i >= 0; i-- {
		output.WriteString(string(r[i]))
	}

	return output.String()
}

//
// find all matching tiles and properly link them together
//
func findMatchingTiles(tiles map[int]tileDef) map[int]tileDef {
	for _, tile := range tiles {
		for _, otherTile := range tiles {
			if tile.id != otherTile.id {
				tile, otherTile = matchTiles(tile, otherTile)
				tiles[tile.id] = tile
				tiles[otherTile.id] = otherTile
			}
		}
	}

	for _, tile := range tiles {
		fmt.Println(tile)
	}

	return tiles
}

//
// check if the given tiles are a match and in this case
// return the linked tiles
//
func matchTiles(tile tileDef, otherTile tileDef) (tileDef, tileDef) {
	if tile.leftTileID == 0 {
		tile.leftTileID = checkOtherTile(tile.left, &otherTile, tile.id)
	}
	if tile.topTileID == 0 {
		tile.topTileID = checkOtherTile(tile.top, &otherTile, tile.id)
	}
	if tile.rightTileID == 0 {
		tile.rightTileID = checkOtherTile(tile.right, &otherTile, tile.id)
	}
	if tile.bottomTileID == 0 {
		tile.bottomTileID = checkOtherTile(tile.bottom, &otherTile, tile.id)
	}

	return tile, otherTile
}

//
// check if the other tile is linked to the given tile
//
func checkOtherTile(border [2]int, otherTile *tileDef, tileID int) int {
	if border[0] == otherTile.left[0] || border[0] == otherTile.left[1] ||
		border[1] == otherTile.left[0] || border[1] == otherTile.left[1] {
		(*otherTile).leftTileID = tileID
		return otherTile.id
	}
	if border[0] == otherTile.top[0] || border[0] == otherTile.top[1] ||
		border[1] == otherTile.top[0] || border[1] == otherTile.top[1] {
		(*otherTile).topTileID = tileID
		return otherTile.id
	}
	if border[0] == otherTile.right[0] || border[0] == otherTile.right[1] ||
		border[1] == otherTile.right[0] || border[1] == otherTile.right[1] {
		(*otherTile).rightTileID = tileID
		return otherTile.id
	}
	if border[0] == otherTile.bottom[0] || border[0] == otherTile.bottom[1] ||
		border[1] == otherTile.bottom[0] || border[1] == otherTile.bottom[1] {
		(*otherTile).bottomTileID = tileID
		return otherTile.id
	}

	return 0
}

//
// get product of the four corner tiles' id
//
func calculateEdgeValues(tiles map[int]tileDef) int {
	result := 1
	for _, tile := range tiles {
		if isCorner(tile) {
			result *= tile.id
		}
	}
	return result
}

//
// check if the given tile is a corner tile with
// just two linked tiles.
//
func isCorner(tile tileDef) bool {
	var activeCorners int
	if tile.leftTileID > 0 {
		activeCorners++
	}

	if tile.rightTileID > 0 {
		activeCorners++
	}

	if tile.topTileID > 0 {
		activeCorners++
	}

	if tile.bottomTileID > 0 {
		activeCorners++
	}

	return activeCorners == 2
}
