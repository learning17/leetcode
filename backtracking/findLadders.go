package main
import "fmt"
// https://leetcode-cn.com/problems/word-transformer-lcci/

func main() {
	beginWord, endWord := "qa", "sq"
	wordList := []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}
	res := findLadders(beginWord, endWord, wordList)
	fmt.Println(res)
}

func findLadders(beginWord string, endWord string, wordList []string) []string {
	dict := findWords(beginWord, wordList)
	if _, ok := dict[endWord]; !ok {
		return []string{}
	}
	path, res := []string{beginWord}, []string{}
	visited := make(map[string]bool)
	var backtrack func(string) bool 
	backtrack = func(word string) bool {
		if word == endWord {
			res = append(res, path...)
			return true
		}
		choices, ok := dict[word]
		if !ok {
			return false
		}
		for _, choice := range choices {
			if visited[choice] {
				continue
			}
			path = append(path, choice)
			visited[choice] = true
			if backtrack(choice) {
				return true
			}
			path = path[:len(path)-1]
		}
		return false
	}
	backtrack(beginWord)
	return res
}

func findWords(word string, wordList []string) map[string][]string {
	compare := func(word1, word2 string) bool {
		diffNum := 0
		for i := 0; i < len(word1); i++ {
			if diffNum > 1 {
				return false
			}
			if word1[i] != word2[i] {
				diffNum++
			}
		}
		return diffNum == 1
	}
	dict := make(map[string][]string)
	wordList = append(wordList, word)
	size := len(wordList)
	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			if i == j || !compare(wordList[i], wordList[j]){
				continue
			}
			dict[wordList[i]] = append(dict[wordList[i]], wordList[j])
		}
	}
	return dict
}
