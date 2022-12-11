package utils

import "strings"

func SimplifyOutput(outputArray []string) []string {
	var output []string
	for i := range outputArray {
		outputArray[i] = strings.TrimSpace(outputArray[i])
	}
	for j := range outputArray {
		if outputArray[j] != "" {
			output = append(output, outputArray[j])
		}
	}
	return output
}
