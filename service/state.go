package service

func StateInNum(state string) uint {
	switch state {
	case "Public Sea":
		return 0
	case "First":
		return 1
	case "Second":
		return 2
	case "Final":
		return 3
	}
	return 999
}

func StateInChinese(state string) string {
	switch state {
	case "Public Sea":
		return "公海"
	case "First":
		return "一面"
	case "Second":
		return "二面"
	case "Final":
		return "入潮"
	}
	return "null"
}

func NextState(curState string) string {
	switch curState {
	case "Public Sea":
		return StateInChinese("First")
	case "First":
		return StateInChinese("Second")
	case "Second":
		return StateInChinese("Final")
	}
	return "null"
}