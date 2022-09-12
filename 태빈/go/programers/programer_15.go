package coding

import "bytes"

func HidePhoneNumer(phone_number string) string{
	if len(phone_number) == 4 {
		return phone_number
	}
	phoneNumberLength := len(phone_number)
	backNumber := phone_number[phoneNumberLength-4:]
	
	var b bytes.Buffer
	for i:= 0; i<phoneNumberLength-4;i++{
		b.WriteString("*")
	}
	return b.String()+backNumber
}