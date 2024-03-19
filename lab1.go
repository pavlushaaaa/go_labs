package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var totalSize = 34
var memoryBlocks = [][2]int{}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func isDigit(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func allocateMemory(size int) int {
	if len(memoryBlocks) == 0 {
		if size <= totalSize {
			memoryBlocks = append(memoryBlocks, [2]int{0, size - 1})
			return 0
		}
		return -1
	}

	if memoryBlocks[0][0] >= size {
		memoryBlocks = append([][2]int{{0, size - 1}}, memoryBlocks...)
		return 0
	}
	for i := 0; i < len(memoryBlocks)-1; i++ {
		if memoryBlocks[i+1][0]-memoryBlocks[i][1]-1 >= size {
			firstFreeIndex := memoryBlocks[i][1] + 1
			newSubslice := [2]int{firstFreeIndex, firstFreeIndex + size - 1}
			memoryBlocks = append(memoryBlocks[:i+1], memoryBlocks[i:]...)
			memoryBlocks[i+1] = newSubslice
			return firstFreeIndex
		}
	}
	lastIndex := memoryBlocks[len(memoryBlocks)-1][1]
	if lastIndex+size < totalSize {
		memoryBlocks = append(memoryBlocks, [2]int{
			lastIndex + 1,
			lastIndex + size,
		})
		return lastIndex + 1
	}
	return -1
}

func freeMemory(index int) int {
	for i := 0; i < len(memoryBlocks); i++ {
		if memoryBlocks[i][0] == index {
			memoryBlocks = append(memoryBlocks[:i], memoryBlocks[i+1:]...)
			return index
		}
	}
	return -1
}

func visualizeMemory() {
	cellSize := len(strconv.Itoa(totalSize)) + 1
	resultString := ""
	for i := 0; i < totalSize; i++ {
		resultString += repeat(" ", cellSize)
	}

	byteSlice := []byte(resultString)
	for block := 0; block < len(memoryBlocks); block++ {
		firstIndexSize := len(strconv.Itoa(memoryBlocks[block][0]))
		firtIndexSlice := []byte(strconv.Itoa(memoryBlocks[block][0]))
		for i := memoryBlocks[block][0] * cellSize; i < (memoryBlocks[block][1]+1)*cellSize; i++ {
			byteSlice[i] = 'x'
		}
		byteSlice[memoryBlocks[block][1]*cellSize+(cellSize-1)] = '|'

		insertNumberIndex := memoryBlocks[block][0] * cellSize
		for i := 0; i < firstIndexSize; i++ {
			byteSlice[insertNumberIndex] = firtIndexSlice[i]
			insertNumberIndex++
		}
	}

	for i := 1; i < len(byteSlice); i++ {
		if isDigit(byteSlice[i]) {
			if byteSlice[i-1] == ' ' {
				byteSlice[i-1] = '|'
			}
		}
	}

	bytesCounter := 0
	resultSlice := []byte{}
	resultSlice = append(resultSlice, '|')

	for ; bytesCounter < len(byteSlice); bytesCounter++ {
		resultSlice = append(resultSlice, byteSlice[bytesCounter])
		if (bytesCounter+1)%cellSize == 0 && ((bytesCounter+1)/cellSize)%10 == 0 {
			resultSlice[len(resultSlice)-1] = '|'
			resultSlice = append(resultSlice, '\n')
			if bytesCounter != len(byteSlice)-1 {
				resultSlice = append(resultSlice, '|')
			}
		}
	}
	if totalSize%10 != 0 {
		if resultSlice[len(resultSlice)-1] != '|' {
			resultSlice[len(resultSlice)-1] = '|'
		}
	}

	fmt.Println(string(resultSlice))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		parsedCommand := strings.Split(input, " ")
		command := parsedCommand[0]
		args := parsedCommand[1:]

		switch command {
		case "allocate":
			if len(args) != 1 {
				fmt.Println("Error: allocate command requires one argument")
				continue
			}
			size, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error: invalid size argument")
				continue
			}
			exitCode := allocateMemory(size)
			if exitCode == -1 {
				fmt.Println("Error: no free blocks are available")
			} else {
				fmt.Printf("Info: allocated %d block(s) of memory\n", size)
			}
		case "free":
			if len(args) != 1 {
				fmt.Println("Error: free command requires one argument")
				continue
			}
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error: invalid index argument")
				continue
			}
			exitCode := freeMemory(index)
			if exitCode == -1 {
				fmt.Println("Error: block that starts with given index doesn't exist")
			} else {
				fmt.Printf("Info: freed block at index %d\n", index)
			}
		case "print":
			visualizeMemory()
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		case "help":
			fmt.Println(" - help: show this output")
			fmt.Println(" - allocate BLOCK_SIZE: allocate block of size BLOCK_SIZE")
			fmt.Println(" - free BLOCK_INDEX: free block that starts at index BLOCK_INDEX")
			fmt.Println(" - print: visualize current memory status")
			fmt.Println(" - exit: exit emulator")
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}
