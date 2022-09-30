package leetcode

func destCity(paths [][]string) string {
	cityA := map[string]bool{}
	for _, path := range paths {
		cityA[path[0]] = true
	}
	for _, path := range paths {
		if !cityA[path[1]] {
			return path[1]
		}
	}
	return ""
}
