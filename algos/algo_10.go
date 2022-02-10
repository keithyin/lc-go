package algos

func isMatch(s string, p string) bool {
	cache := make(map[[2]int]bool)
	return isMatchCoreWithCache(s, p, 0, 0, cache)
}

func isMatchCore(s, p string, sPos, pPos int) bool {
	if sPos >= len(s) && pPos >= len(p) {
		return true
	}
	if sPos >= len(s) {
		if (pPos+1) < len(p) && p[pPos+1] == '*' {
			return isMatchCore(s, p, sPos, pPos+2)
		}
		return false
	}
	if pPos >= len(p) {
		return false
	}

	if (pPos+1) < len(p) && p[pPos+1] == '*' {
		if s[sPos] == p[pPos] || p[pPos] == '.' {
			return isMatchCore(s, p, sPos+1, pPos+2) || isMatchCore(s, p, sPos+1, pPos) || isMatchCore(s, p, sPos, pPos+2)
		} else {
			return isMatchCore(s, p, sPos, pPos+2)
		}
	} else {
		if s[sPos] == p[pPos] || p[pPos] == '.' {
			return isMatchCore(s, p, sPos+1, pPos+1)
		} else {
			return false
		}
	}
}

func isMatchCoreWithCache(s, p string, sPos, pPos int, cache map[[2]int]bool) bool {
	if match, ok := cache[[2]int{sPos, pPos}]; ok {
		return match
	}

	if sPos >= len(s) && pPos >= len(p) {
		cache[[2]int{sPos, pPos}] = true
		return true
	}
	if sPos >= len(s) {
		if (pPos+1) < len(p) && p[pPos+1] == '*' {
			match := isMatchCore(s, p, sPos, pPos+2)
			cache[[2]int{sPos, pPos}] = match
			return match
		}
		cache[[2]int{sPos, pPos}] = false
		return false
	}
	if pPos >= len(p) {
		cache[[2]int{sPos, pPos}] = false
		return false
	}

	if (pPos+1) < len(p) && p[pPos+1] == '*' {
		if s[sPos] == p[pPos] || p[pPos] == '.' {
			match := isMatchCore(s, p, sPos+1, pPos+2)
			cache[[2]int{sPos, pPos}] = match
			if match {
				return match
			}

			match = isMatchCore(s, p, sPos+1, pPos)
			cache[[2]int{sPos, pPos}] = match
			if match {
				return match
			}

			match = isMatchCore(s, p, sPos, pPos+2)
			cache[[2]int{sPos, pPos}] = match

			return match
		} else {
			match := isMatchCore(s, p, sPos, pPos+2)
			cache[[2]int{sPos, pPos}] = match
			return match
		}
	} else {
		if s[sPos] == p[pPos] || p[pPos] == '.' {
			match := isMatchCore(s, p, sPos+1, pPos+1)
			cache[[2]int{sPos, pPos}] = match
			return match
		} else {
			cache[[2]int{sPos, pPos}] = false
			return false
		}
	}
}
