package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	friends := friendships[username]

	suggestedFriends := make([]string, 0)

	for _, friend := range friends {
		for _, newFriend := range friendships[friend] {

			notNew := false

			if newFriend == username {
				continue
			}

			for _, check := range suggestedFriends {
				if check == newFriend {
					notNew = true
					break
				}
			}

			for _, check := range friends {
				if check == newFriend || notNew {
					notNew = true
					break
				}
			}

			if notNew {
				continue
			}

			suggestedFriends = append(suggestedFriends, newFriend)

		}
	}

	if len(suggestedFriends) == 0 {
		return nil
	}

	return suggestedFriends
}
