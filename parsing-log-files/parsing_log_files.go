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
			fmt.Println(line)
			ssub := re.FindStringSubmatch(line)
			//fmt.Println(ssub)
			count += len(ssub)
			// for range ssub {
			// 	count++
			// }
		}
	fmt.Println(count)
	return count
}

func RemoveEndOfLineText(text string) string {
	var result string
	return result
}

func TagWithUserName(lines []string) []string {
	var result []string
	return result
}
