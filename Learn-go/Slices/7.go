package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for index, element := range msg {

		for _, badelement := range badWords {
			if element == badelement {
				return index
			}
		}

	}
	return -1
}
