package coding

import (
	"bytes"
	"strconv"
)

func TranslateString(s string) int {
	var result int
	var resultByte bytes.Buffer
	numberDict := map[string]string{
		"zero":"0",
		"one":"1",
		"two":"2",
		"three":"3",
		"four":"4",
		"five":"5",
		"six":"6",
		"seven":"7",
		"eight":"8",
		"nine":"9",
	}
	runes := []rune(s)
	var b bytes.Buffer
	for i:=0; i < len(runes); i++ {
		r := runes[i]
		if r >47 && r <58{
			resultByte.WriteRune(r)
			continue
		}
		b.WriteRune(r)
		v, f := numberDict[b.String()]
		if(f){
			resultByte.WriteString(v)
			b = *bytes.NewBuffer([]byte(""))
		}
	}
	result, _ = strconv.Atoi(resultByte.String())
	return result
}