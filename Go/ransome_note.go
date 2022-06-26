package main

func canConstruct(ransomNote string, magazine string) bool {
	charRansom := []rune(ransomNote)
	charMagazine := []rune(magazine)

	mapRansom := make(map[rune]int16)
	mapMagazine := make(map[rune]int16)
	
	for i := 0; i < len(charRansom); i++ {
		mapRansom[charRansom[i]] += 1
	}

	for i := 0; i < len(charMagazine); i++ {
		mapMagazine[charMagazine[i]] += 1
	}

	for k, _ := range mapRansom {
		if mapRansom[k] > mapMagazine[k] {
			return false
		}
	}

	return true
}
