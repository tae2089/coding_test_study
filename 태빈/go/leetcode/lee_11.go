package leetcode

func maxArea(height []int) int {
	startPoint := 0
	endPoint := len(height) - 1

	maxArea := 0
	for startPoint < endPoint {
		area := minPoints(height[startPoint], height[endPoint]) * (endPoint - startPoint)
		if area > maxArea {
			maxArea = area
		}
		if height[startPoint] < height[endPoint] {
			startPoint++
		}else{
			endPoint--
		}
	}
	return maxArea
}


func minPoints(startPoint, endPoint int) int {
	if startPoint < endPoint {
		return startPoint
	}
	return endPoint
}