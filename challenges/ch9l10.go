package main

func getNameCounts(names []string) map[rune]map[string]int {
	nestedMap := make(map[rune]map[string]int)

	for _, name := range names {
		firstChar := rune(name[0])

		// if firstChar == 240 {
		// 	nestedMap['😊'][name] = 1
		// 	continue
		// }

		if _, ok := nestedMap[firstChar]; !ok {
			nestedMap[firstChar] = make(map[string]int)
			// fmt.Println("made innermap", firstChar)
		}

		// namesList := nestedMap[rune(firstChar)]

		// if _, ok := namesList[name]; !ok {
		// 	namesList[name] = 1
		// } else {
		// 	namesList[name]++
		// }

		nestedMap[firstChar][name]++
	}

	return nestedMap
}

// need to figure out how to fix edge case of smiley face not being recognized
