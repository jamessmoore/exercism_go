package parsinglogfiles

import (
	"regexp" 
	"fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]+\d*`)

	if re.MatchString(text) {
		fmt.Println(text)
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	//sl = re.Split("123 456", -1) 
	re := regexp.MustCompile(`^<(-*|~*|\**)\>`)
	return re.Split(text, -1) 
}

func CountQuotedPasswords(lines []string) int {
	// panic("Please implement the CountQuotedPasswords function")
	var count int
	return count
}

func RemoveEndOfLineText(text string) string {
	var return string
	return return
}

func TagWithUserName(lines []string) []string {
	var return string
	return return
}
