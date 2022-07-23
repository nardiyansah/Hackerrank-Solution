package main

func isSubsequence(s string, t string) bool {
	idxs, idxt := 0, 0

	if len(s) == 0 {
		return true
	}

	for {
		// condition to out from loop and string not subsequence
		if idxs >= len(s) || idxt >= len(t) {
			break
		}

		if s[idxs] == t[idxt] {
			if idxs == len(s)-1 {
				return true
			}
			idxs += 1
		}
		idxt += 1
	}

	return false
}
