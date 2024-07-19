package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to get memory info from /proc/meminfo
func getMemInfo() (map[string]int, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	memInfo := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			key := strings.Trim(fields[0], ":")
			value, err := strconv.Atoi(fields[1])
			if err == nil {
				memInfo[key] = value
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return memInfo, nil
}

// Function to determine if available memory is near zero
func isMemoryNearZero(threshold int) (bool, error) {
	memInfo, err := getMemInfo()
	if err != nil {
		return false, err
	}

	availableMem, exists := memInfo["MemAvailable"]
	if !exists {
		return false, fmt.Errorf("MemAvailable not found in /proc/meminfo")
	}

	return availableMem < threshold, nil
}

func main() {
	threshold := 1024 // Threshold in kB (1 MB)
	nearZero, err := isMemoryNearZero(threshold)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if nearZero {
		fmt.Println("Memory is near zero!")
	} else {
		fmt.Println("Memory is sufficient.")
	}
}
