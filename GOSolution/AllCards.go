package main

type OneColorCards struct {
	cards []int
}

var allCards map[string]OneColorCards

func makeCards() {
	allCards = nil
	allCards = make(map[string]OneColorCards)
	allCards["Diamonds"] = OneColorCards{[]int{2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10, 11}}
	allCards["Hearts"] = OneColorCards{[]int{2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10, 11}}
	allCards["Spadles"] = OneColorCards{[]int{2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10, 11}}
	allCards["Clubs"] = OneColorCards{[]int{2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10, 11}}
}

func binarySearch(array []int, lookingFor int) bool {
	if len(array) == 0 {
		return false
	}
	if len(array) == 0 && array[0] == lookingFor {
		return false
	}
	var mid, assumption int
	var min = 0
	var high = len(array) - 1

	for min <= high {
		mid = (min + high) / 2
		assumption = mid
		if assumption == lookingFor {
			array = append(array[:assumption], array[assumption+1:]...)
			return true
		}
		if assumption > lookingFor {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return false
}

func existsCard(elem int) (bool, string) {
	for name, x := range allCards {
		if binarySearch(x.cards, elem) {
			return true, name
		}
	}
	return false, ""
}

//{2, 2, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10, 11}
