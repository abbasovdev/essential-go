package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	CHAPTERS_PATH   = "chapters"
	TOC_PATH        = "README.md"
	MAIN_FILE       = "main.go"
	VALIDATION_FILE = "output.txt"
	TERMINAL_GREEN  = "\033[32m"
	TERMINAL_RED    = "\033[31m"
	TERMINAL_RESET  = "\033[0m"
)

func main() {
	topicNo, validate, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	topicMap, err := getTopicMap()
	if err != nil {
		fmt.Println(err)
		return
	}

	topicPath, ok := topicMap[topicNo]
	if !ok {
		fmt.Printf("Couldn't able to find a #%v topic\n", topicNo)
		return
	}

	if validate {
		err = validateTopic(topicPath)
		if err != nil {
			os.Exit(1)
		}
	} else {
		err = executeTopic(topicPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func parseArgs() (string, bool, error) {
	topicNo := ""
	validate := false
	if len(os.Args) < 2 {
		return topicNo, validate, fmt.Errorf("Usage: go run topic.go <topic no> (optional: <validate>)")
	}
	topicNo = os.Args[1]
	if !isArgValid(topicNo) {
		return topicNo, validate, fmt.Errorf("Invalid topic no. Use a positive integer (e.g. 1, 10)")
	}
	if len(os.Args) == 3 {
		validate = os.Args[2] == "validate"
	}
	return topicNo, validate, nil
}

func isArgValid(topicNo string) bool {
	n, err := strconv.Atoi(topicNo)
	return err == nil && n > 0
}

func getTopicMap() (map[string]string, error) {
	tocBytes, err := os.ReadFile(TOC_PATH)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", TOC_PATH, err)
	}

	topicMap := make(map[string]string)
	nextSeq := 1
	lines := strings.Split(string(tocBytes), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "|") || strings.HasPrefix(line, "| Chapters |") {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}
		topicNo := strings.TrimSpace(parts[2])
		topicsCell := strings.TrimSpace(parts[3])
		tocLinkRegex := regexp.MustCompile(`\]\((.*?)\)`)
		link := tocLinkRegex.FindStringSubmatch(topicsCell)
		if len(link) < 2 {
			continue
		}
		path := strings.TrimSpace(link[1])
		if path == "" || strings.HasPrefix(path, "#") {
			continue
		}
		dir := path
		if strings.HasSuffix(dir, "/README.md") {
			dir = strings.TrimSuffix(dir, "/README.md")
		} else if strings.HasSuffix(dir, "README.md") {
			dir = filepath.Dir(dir)
		}
		dir = strings.TrimPrefix(dir, "./")
		dir = filepath.ToSlash(dir)
		if dir == "" {
			continue
		}
		if _, err := os.Stat(dir); err != nil {
			continue
		}
		var key string
		if n, err := strconv.Atoi(topicNo); err == nil && n > 0 {
			key = strconv.Itoa(n)
			if n >= nextSeq {
				nextSeq = n + 1
			}
		} else {
			key = strconv.Itoa(nextSeq)
			nextSeq++
		}
		topicMap[key] = dir
	}
	return topicMap, nil
}

func executeTopic(topicPath string) error {
	fmt.Printf("--- Running: %s ---\n", topicPath)
	_, skipped, err := runGo(topicPath, MAIN_FILE, false)
	if skipped {
		printOK()
		return nil
	}
	if err != nil {
		return fmt.Errorf("Error executing %s: %v", MAIN_FILE, err)
	}
	return nil
}

func validateTopic(topicPath string) error {
	fmt.Printf("--- Validating: %s ---\n", topicPath)
	lessonDir, err := filepath.Abs(topicPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sERROR!%s validate: %v\n", TERMINAL_RED, TERMINAL_RESET, err)
		return err
	}
	expectedPath := filepath.Join(lessonDir, VALIDATION_FILE)
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		printOK()
		return nil
	}
	expectedBytes, err := os.ReadFile(expectedPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sERROR!%s validate: reading %s: %v\n", TERMINAL_RED, TERMINAL_RESET, VALIDATION_FILE, err)
		return err
	}
	expected := strings.TrimSpace(string(expectedBytes))

	out, skipped, err := runGo(lessonDir, MAIN_FILE, true)
	if skipped {
		fmt.Fprintf(os.Stderr, "%sERROR!%s validate: %s not found\n", TERMINAL_RED, TERMINAL_RESET, MAIN_FILE)
		return fmt.Errorf("%s not found", MAIN_FILE)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sERROR!%s validate: running lesson failed: %v\n%s", TERMINAL_RED, TERMINAL_RESET, err, out)
		return err
	}
	actual := strings.TrimSpace(string(out))
	if actual != expected {
		fmt.Fprintf(os.Stderr, "%sERROR!%s validation failed\n  expected: %q\n  actual:   %q\n", TERMINAL_RED, TERMINAL_RESET, expected, actual)
		return nil
	}
	printOK()
	return nil
}

func runGo(topicPath, fileName string, capture bool) (out []byte, skipped bool, err error) {
	filePath := filepath.Join(topicPath, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, true, nil
	}
	cmd := exec.Command("go", "run", filePath)
	if capture {
		out, err = cmd.CombinedOutput()
		return out, false, err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return nil, false, cmd.Run()
}

func printOK() {
	fmt.Printf("%sOK!%s\n", TERMINAL_GREEN, TERMINAL_RESET)
}
