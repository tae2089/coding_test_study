package coding

import "bytes"

func FindCenterChar(s string ) string{
		
	var b bytes.Buffer
	length := len(s)
	if length % 2 == 0 {
		b.WriteByte(s[(length/2)-1])
		b.WriteByte(s[length/2])
	}else{
		b.WriteByte(s[(length/2)] ) 
	}
	return b.String()
}