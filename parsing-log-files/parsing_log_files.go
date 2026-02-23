package parsinglogfiles

import (
	"regexp" 
	"fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]+\d*`)

	if re.MatchString(text) {
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<([-*~=]*)\>`)
	return re.Split(text, -1) 
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*?password.*?"`)
	var count int
		for _, line := range lines{
			// fmt.Println(line)
			ssub := re.FindStringSubmatch(line)
			count += len(ssub)
		}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	s := re.ReplaceAllString(text, "")
	return s
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s*[A-Z]\w*`)
	result := []string{}
	for _, line := range lines {
		fmt.Println(line)
		sm := re.FindStringSubmatch(line)
		
		if sm != nil {
			fmt.Println(sm)
		// 	line = fmt.Sprintf("[USR] %s %s", sm[1], line)
		}
		result = append(result, line)
	}
	return result
}
