package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var numbers map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func findDigit(s string) (int, error) {
	var result []string

	for _, chr := range s {
		if unicode.IsDigit(chr) {
			result = append(result, string(chr))
		}
	}

	if len(result) == 0 {
		return 0, fmt.Errorf("cannot find at least 1 digit in %q", s)
	}

	number, err := strconv.Atoi(fmt.Sprintf("%s%s", result[0], result[len(result)-1]))
	if err != nil {
		return 0, err
	}

	return number, nil
}

func findQuestionableDigit(s string) (int, error) {
	var result []string

	for i, chr := range s {
		if unicode.IsDigit(chr) {
			result = append(result, string(chr))
		} else {
			for key, value := range numbers {
				if len(key)+i > len(s) {
					continue
				}

				if key == s[i:i+len(key)] {
					result = append(result, value)
					break
				}
			}
		}
	}

	if len(result) == 0 {
		return 0, fmt.Errorf("cannot find at least 1 digit in %q", s)
	}

	number, err := strconv.Atoi(fmt.Sprintf("%s%s", result[0], result[len(result)-1]))
	if err != nil {
		return 0, err
	}

	return number, nil
}

func run() error {
	modeFlag := flag.String("mode", "1", "Whether to solve part 1 or part 2")
	flag.Parse()

	if flag.NArg() < 1 {
		return errors.New("need input file")
	}

	filename := flag.Arg(0)
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	sum := 0

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		var num int
		var err error
		if *modeFlag == "1" {
			num, err = findDigit(line)
		} else {
			num, err = findQuestionableDigit(line)
		}
		if err != nil {
			return fmt.Errorf("error finding digits in %q: %w", line, err)
		}

		fmt.Printf("%q %d\n", line, num)
		sum += num
	}

	fmt.Printf("total = %d\n", sum)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
