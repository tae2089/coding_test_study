package coding

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func SumDivision(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		if n % i == 0{
			result += i
		}
	}
	return result
}

