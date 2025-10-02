package functions

import (
	"github.com/blutspende/go-astm/v3/enums/lineseparator"
	"github.com/blutspende/go-astm/v3/errmsg"
	"github.com/blutspende/go-astm/v3/models/astmmodels"
	"strings"
)

func SliceLines(input string, config *astmmodels.Configuration) (output []string, err error) {
	// Check for empty input
	if input == "" {
		return nil, errmsg.ErrLineProcessingEmptyInput
	}

	// A line separator has to be provided if auto-detect is disabled
	if !config.AutoDetectLineSeparator && config.LineSeparator == "" {
		return nil, errmsg.ErrLineProcessingNoLineSeparator
	}

	var lines []string
	if !config.AutoDetectLineSeparator {
		// Line separator provided in config, no auto-detect
		lines = strings.Split(input, config.LineSeparator)
	} else {
		// Auto-detect line separator
		lfCnt := 0
		crCnt := 0
		for _, c := range input {
			if c == rune(lineseparator.LF[0]) {
				lfCnt++
			} else if c == rune(lineseparator.CR[0]) {
				crCnt++
			}
		}

		if lfCnt == 0 && crCnt == 0 {
			// Note: single line inputs are allowed, but it could indicate a problem
			lines = append(lines, input)
		} else {
			if lfCnt > 0 && crCnt > 0 && lfCnt != crCnt {
				// Mixed number of different line endings are not allowed
				return nil, errmsg.ErrLineProcessingInvalidLinebreak
			}
			if lfCnt == 0 {
				// Only CR line endings: replace with LF
				input = strings.ReplaceAll(input, lineseparator.CR, lineseparator.LF)
			} else {
				// Equally mixed line endings: only keep LF
				input = strings.ReplaceAll(input, lineseparator.CR, "")
			}
			lines = strings.Split(input, lineseparator.LF)
		}
	}

	for i := range lines {
		lines[i] = strings.Trim(lines[i], " ")
		if lines[i] != "" {
			output = append(output, lines[i])
		}
	}

	return output, nil
}

func BuildLines(input []string, config *astmmodels.Configuration) (output string) {
	linebreak := lineseparator.LF
	if config.LineSeparator != "" && !config.AutoDetectLineSeparator {
		linebreak = config.LineSeparator
	}

	for i, line := range input {
		output += line
		if i < len(input)-1 {
			output += linebreak
		}
	}

	return output
}
