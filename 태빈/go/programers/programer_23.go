package coding

import "bytes"

func CreateSquare(a,b int){
	var bc bytes.Buffer
	for i:=0;i<b;i++{
		for j:=0;j<a;j++{
			bc.WriteString("*")
		}
		bc.WriteString("\n")
	}
}