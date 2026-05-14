package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func buildLessonsExcluding(count int, exclude string) ([]int, error) {

	parts := strings.Split(exclude, ",")

	var excludeInts []int
	var finalLessons []int

	// Separate string numbers into array of strings
	for _, part := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(part))

		if err != nil {
			return nil, err
		}

		excludeInts = append(excludeInts, n)
	}

	// Build final lessons array by appending
	for i := 1; i <= count; i++ {
		shouldExclude := false

		// Look if it should be excluded
		for _, n := range excludeInts {
			if n == i {
				shouldExclude = true
				break
			}
		}

		// Add to array if not
		if !shouldExclude {
			finalLessons = append(finalLessons, i)
		}
	}

	return finalLessons, nil
}

func main() {
	var chapter int
	var lessonCount int
	var lessonArray []int
	var ans string
	var exclude string

	// Ask for chapter
	fmt.Print("Chapter: ")
	_, err := fmt.Scan(&chapter)

	if err != nil {
		fmt.Println("failed to read chapter:", err)
		return
	}

	// Ask for lesson count
	fmt.Print("Lesson count: ")
	_, err = fmt.Scan(&lessonCount)

	if err != nil {
		fmt.Println("failed to read lesson count:", err)
		return
	}

	// Ask for lessons to exclude/include
	for {
		fmt.Print("Would you like to exclude lessons? y/n: ")
		_, err = fmt.Scan(&ans)

		if err != nil {
			fmt.Println("failed to read y/n choice:", err)
			return
		}

		if ans == "y" {
			// Build array while excluding
			fmt.Print("Give the lessons you'd like to exclude, separated by ,:")
			_, err = fmt.Scan(&exclude)

			if err != nil {
				fmt.Println("failed to read lessons to exclude:", err)
				return
			}

			lessonArray, err = buildLessonsExcluding(lessonCount, exclude)

			if err != nil {
				fmt.Println("failed to build lesson list:", err)
				return
			}
			break
		} else if ans == "n" {
			// Build array
			for i := 1; i <= lessonCount; i++ {
				lessonArray = append(lessonArray, i)
			}
			break
		} else {
			fmt.Println("Invalid option.")
		}
	}

	// create directories and files

	chDir := fmt.Sprintf("ch%d", chapter)

	for _, lesson := range lessonArray {

		lesDir := fmt.Sprintf("l%02d", lesson)
		path := filepath.Join(chDir, lesDir)
		mainPath := filepath.Join(path, "main.go")
		testPath := filepath.Join(path, "main_test.go")

		err = os.MkdirAll(path, 0o755)
		if err != nil {
			fmt.Println("failed to create directory:", err)
			return
		}

		// Create main.go only if it does not already exist.
		_, err = os.Stat(mainPath)
		if os.IsNotExist(err) {
			err = os.WriteFile(mainPath, []byte{}, 0o644)
			if err != nil {
				fmt.Println("failed to create main.go:", err)
				return
			}
		} else if err != nil {
			fmt.Println("failed to check main.go:", err)
			return
		} else {
			fmt.Println("skipped existing:", mainPath)
		}

		// Create main_test.go only if it does not already exist.
		_, err = os.Stat(testPath)
		if os.IsNotExist(err) {
			err = os.WriteFile(testPath, []byte{}, 0o644)
			if err != nil {
				fmt.Println("failed to create main_test.go:", err)
				return
			}
		} else if err != nil {
			fmt.Println("failed to check main_test.go:", err)
			return
		} else {
			fmt.Println("skipped existing:", testPath)
		}
	}

	fmt.Println("Scaffold created successfully.")
}
