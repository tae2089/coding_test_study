package coding

import (
	"bytes"
	"strings"
)

func StrangeString(	s string) string {
	sArray := strings.Split(s," ")
	
	for i := 0; i < len(sArray); i++ {
		word := sArray[i]
		var b bytes.Buffer
		for idx, v := range word {
			if idx == 0 || idx % 2 == 0{
				b.WriteString(strings.ToUpper(string(v)))
			}else{
				b.WriteString(strings.ToLower(string(v)))
			}
		}
		sArray[i] = b.String()
	}
	
	return strings.Join(sArray, " ")
}