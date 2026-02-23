package parsinglogfiles

import (
	"regexp" 
	"fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]+\d*`)

	if re.MatchString(text) {
		// fmt.Println(text)
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<([-*~=]*)\>`)
	// fmt.Println(re.Split(text, -1))
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
	// fmt.Println(count)
	return count
}

func RemoveEndOfLineText(text string) string {
	// re := regexp.MustCompile(`^.*end-of-line(?:\\d+.*)*$`)
	re := regexp.MustCompile(`end-of-line\d*`)
	// fmt.Println(text)
	//.*end-of-line\d*.*
	s := re.ReplaceAllString(text, "")
	// fmt.Println(s)
	return s
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s*[A-Z]\w*`)
	var result []string
	var newLine string
	for _, line := range lines {
		fmt.Println(line)
		if sm := re.FindStringSubmatch(line)
			// for _, s := range sm {
			// 	if s != "User" {
			// 		fmt.Println(s)
			// 	}
			// }
			// fmt.Printf("%q", sm)
			// fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
			newLine = "[USR] " + line 
		} 
		newLine += line
		// fmt.Println(newLine)
		result = append(result, newLine)
	}
	return result
}
