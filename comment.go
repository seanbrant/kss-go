package kss

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
)

var singleLineRegexp, _ = regexp.Compile(`^\s*\/\/`)
var singleLineStripRegexp, _ = regexp.Compile(`\s*\/\/`)

var multiLineStartRegexp, _ = regexp.Compile(`^\s*\/\*`)
var multiLineEndRegexp, _ = regexp.Compile(`.*\*\/`)
var multiLineStartStripRegexp, _ = regexp.Compile(`\s*\/\*`)
var multiLineEndStripRegexp, _ = regexp.Compile(`\*\/`)
var multiLineMiddleStripRegexp, _ = regexp.Compile(`^(\s*\*+)`)

var precedingWhiteSpaceRegexp, _ = regexp.Compile(`^\s*`)

func isSingleLineComment(line string) bool {
	return singleLineRegexp.MatchString(line)
}

func isMultiLineCommentStart(line string) bool {
	return multiLineStartRegexp.MatchString(line)
}

func isMultiLineCommentEnd(line string) bool {
	if isSingleLineComment(line) {
		return false
	} else {
		return multiLineEndRegexp.MatchString(line)
	}
}

func parseSingleLine(line string) string {
	cleaned := singleLineStripRegexp.ReplaceAllLiteralString(line, "")
	return strings.TrimSpace(cleaned)
}

func parseMultiLine(line string) string {
	cleaned := multiLineStartStripRegexp.ReplaceAllLiteralString(line, "")
	cleaned = multiLineEndStripRegexp.ReplaceAllLiteralString(cleaned, "")
	return strings.TrimSpace(cleaned)
}

func normalize(lines []string) string {
	cleaned := []string{}
	indents := []int{}

	for i := range lines {
		line := lines[i]
		line = multiLineMiddleStripRegexp.ReplaceAllLiteralString(line, "")

		cleaned = append(cleaned, line)
		whitespace := len(precedingWhiteSpaceRegexp.FindString(line))

		if whitespace > 0 {
			indents = append(indents, whitespace)
		}
	}

	var indent int

	if len(indents) > 0 {
		sort.Ints(indents)
		indent = indents[0]
	}

	output := []string{}

	for i := range cleaned {
		line := cleaned[i]

		// Strip indent from all but the first line. The
		// first lines indent will already be stripped.
		if len(line) > 0 && i > 0 {
			output = append(output, line[indent:])
		} else {
			output = append(output, line)
		}
	}

	return strings.TrimSpace(strings.Join(output, "\n"))
}

func CommentParser(name string) []string {
	file, _ := os.Open(name)
	reader := bufio.NewReader(file)

	blocks := []string{}
	currentBlock := []string{}

	insideSingleLineBlock := false
	insideMultiLineBlock := false

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		// Parse single-line style
		if isSingleLineComment(line) {
			parsed := parseSingleLine(line)

			if insideSingleLineBlock {
				currentBlock = append(currentBlock, parsed)
			} else {
				currentBlock = []string{parsed}
				insideSingleLineBlock = true
			}
		}

		// Prase multi-line style
		if isMultiLineCommentStart(line) || insideMultiLineBlock {
			parsed := parseMultiLine(line)

			if insideMultiLineBlock {
				currentBlock = append(currentBlock, parsed)
			} else {
				currentBlock = []string{parsed}
				insideMultiLineBlock = true
			}
		}

		// End a multi-line block if detected
		if isMultiLineCommentEnd(line) {
			insideMultiLineBlock = false
		}

		// Store the current block if we're done
		if !isSingleLineComment(line) && !insideMultiLineBlock {
			if len(currentBlock) > 0 {
				blocks = append(blocks, normalize(currentBlock))
			}

			insideSingleLineBlock = false
			currentBlock = []string{}
		}
	}

	return blocks
}
