package service

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