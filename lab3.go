package main

import (
	"bufio"
	"fmt"
	"os"
)

func hashString(s string, tableSize int) int {
	// Simple hashing function - sum of ASCII values
	hashValue := 0
	for _, char := range s {
		hashValue += int(char)
	}
	return hashValue % tableSize
}

func checkUniqueLines(file1, file2 *os.File) bool {
	tableSize := 100 // Initial size of the hash table
	hashTable := make([]string, tableSize)

	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		line := scanner.Text()
		index := hashString(line, tableSize)
		if hashTable[index] == "" {
			hashTable[index] = line
		}
	}

	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		line := scanner.Text()
		index := hashString(line, tableSize)
		if hashTable[index] == "" || hashTable[index] != line {
			return false
		}
	}

	return true
}

func main() {
	file1, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println("Error opening file1:", err)
		return
	}
	defer file1.Close()

	file2, err := os.Open("file2.txt")
	if err != nil {
		fmt.Println("Error opening file2:", err)
		return
	}
	defer file2.Close()

	result := checkUniqueLines(file1, file2)
	if result {
		fmt.Println("Множини унікальних рядків у двох файлах співпадають.")
	} else {
		fmt.Println("Множини унікальних рядків у двох файлах не співпадають.")
	}
}
