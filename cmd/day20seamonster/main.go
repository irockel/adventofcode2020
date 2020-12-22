// read input file input.txt and try to put the tiles together to an image
// the tiles are defined like puzzle elements
// try to find the hidden sea monster
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

	image := resolveImage(tiles)

	hashes := findSeamonster(image)

	fmt.Printf("The amount of hashes aside the sea monster is: %d\n", hashes)
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
	topTileID    int
	rightTileID  int
	bottomTileID int
}

//
// read values from input and image tiles with calculated border
// values as map keys
//
func readImageTiles() map[int]tileDef {
	fmt.Println("reading inputtest.txt")

	file, err := os.Open("./inputtest.txt")
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
// resolve the image from the given tile set
//
func resolveImage(tiles map[int]tileDef) []string {
	used := 0
	image := []string{}
	leftTile := findFirstCorner(tiles)
	fmt.Println(leftTile.id)
	currentTile := leftTile
	imagePos := 0
	for used < len(tiles) {
		for {
			fmt.Printf(" %d ", currentTile.id)
			for pos, line := range currentTile.data {
				if len(image) <= imagePos+pos {
					image = append(image, strings.TrimSpace(line))
				} else {
					image[imagePos+pos] = image[imagePos+pos] + " " + strings.TrimSpace(line)
				}
			}
			used++
			if currentTile.rightTileID > 0 {
				lastID := currentTile.id
				lastBorders := currentTile.right
				currentTile = tiles[currentTile.rightTileID]

				switch lastID {
				case currentTile.rightTileID : 
					if currentTile.right[1] == lastBorders[0] {
						// need to be turned two times
						currentTile = turnRightTile(currentTile)
						currentTile = turnRightTile(currentTile)
					} else {
						// it's just flipped around
						currentTile = flipTile(currentTile)
					}
				case currentTile.topTileID :
					currentTile = turnLeftTile(currentTile)
				case currentTile.bottomTileID :
					currentTile = turnRightTile(currentTile)
				}
			} else {
				break
			}
		}
		fmt.Println()
		imagePos += len(image)
		if leftTile.bottomTileID > 0 {
			lastID := currentTile.id
			lastBorders := currentTile.bottom
			leftTile = tiles[leftTile.bottomTileID]

			switch lastID {
				case leftTile.bottomTileID : 
					if leftTile.bottom[1] == lastBorders[0] {
						// need to be turned two times
						leftTile = turnRightTile(leftTile)
						leftTile = turnRightTile(leftTile)
					} else {
						// it's just flipped around
						leftTile = flipTile(leftTile)
					}
				case leftTile.topTileID :
					leftTile = turnLeftTile(leftTile)
				case currentTile.bottomTileID :
					leftTile = turnRightTile(leftTile)
				}

			currentTile = leftTile
		}
	}

	for _, line := range image {
		fmt.Println(line)
	}

	return image
}

func flipTile(tile tileDef) tileDef {
	flippedTile := tile
	flippedTile.data = []string{}
	for _, line := range tile.data {
		flippedLine := ""
		for i := 0; i < len(line); i++ {
			flippedLine = string(line[i]) + flippedLine 
		}
		flippedTile.data = append(flippedTile.data, flippedLine)
	}

	// flip references
	flippedTile.left = tile.right
	flippedTile.right = tile.left
	flippedTile.leftTileID = tile.rightTileID
	flippedTile.rightTileID = tile.leftTileID

	return flippedTile
}

func turnLeftTile(tile tileDef) tileDef {
	turnedTile := tile
	turnedTile.data = []string{}

	for i := 0; i < len(tile.data[0]); i++ {
		turnedLine := ""
		for j := 0; j < len(tile.data); j++ {
			turnedLine = string(tile.data[j][i]) + turnedLine
		}


	}
	turnedTile.left = tile.top
	turnedTile.top = tile.right
	turnedTile.right = tile.bottom
	turnedTile.bottom = tile.left
	turnedTile.leftTileID = tile.topTileID
	turnedTile.topTileID = tile.rightTileID
	turnedTile.rightTileID = tile.bottomTileID
	turnedTile.bottomTileID = tile.leftTileID
	
	return turnedTile
}

func turnRightTile(tile tileDef) tileDef {
	turnedTile := tile
	
	for i := 0; i < len(tile.data[0]); i++ {
		turnedLine := ""
		for j := 0; j < len(tile.data); j++ {
			turnedLine = turnedLine + string(tile.data[j][i])
		}
	}
	turnedTile.left = tile.bottom
	turnedTile.top = tile.left
	turnedTile.right = tile.top
	turnedTile.bottom = tile.right
	turnedTile.leftTileID = tile.bottomTileID
	turnedTile.topTileID = tile.leftTileID
	turnedTile.rightTileID = tile.topTileID
	turnedTile.bottomTileID = tile.rightTileID

	return turnedTile
}

//
// find the first corner to start with
//
func findFirstCorner(tiles map[int]tileDef) tileDef {
	for _, tile := range tiles {
		if tile.leftTileID == 0 && tile.topTileID == 0 && tile.rightTileID > 0 && tile.bottomTileID > 0 {
			return tile
		}
	}

	return tileDef{}
}

//
// try to find the sea monster which has the pattern
//                   #
// #    ##    ##    ###
//  #  #  #  #  #  #
//
// and marked the used hashes, count all remaining hashes
// and return them
func findSeamonster([]string) int {

	return 0
}
