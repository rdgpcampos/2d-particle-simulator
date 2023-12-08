package util

import (
    "os"
    "log"
    "bufio"
    "strings"
)

func RemoveParticleByIndex[T any](s []*T, index int) []*T {
    ret := make([]*T, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func ParseFileToLines(filePath string) ([]string, error) {
    	// open file describing particle types
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// read particle types into string array
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	defer func() {
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

    return lines, err
}

func AppendLineToFile(line string, filePath string) error {
    file, err := os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

    if _, err := file.WriteString(line+"\n"); err != nil {
        log.Fatal(err)
    }

    return err
}

func SplitLine(s string) []string {
    splitString := strings.Split(s, " ")
    return splitString
}