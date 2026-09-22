package result

func CalculateTotal(marks []float64) float64 {
	var total float64

	for _, mark := range marks {
		total += mark
	}

	return total
}

func CalculateAverage(marks []float64) float64 {
	if len(marks) == 0 {
		return 0
	}

	return CalculateTotal(marks) / float64(len(marks))
}

func CalculateGrade(average float64) string {
	if average >= 90 {
		return "A+"
	} else if average >= 80 {
		return "A"
	} else if average >= 70 {
		return "B"
	} else if average >= 60 {
		return "C"
	} else if average >= 50 {
		return "D"
	}

	return "F"
}

func CalculateResult(marks []float64) string {
	for _, mark := range marks {
		if mark < 40 {
			return "FAIL"
		}
	}

	return "PASS"
}
