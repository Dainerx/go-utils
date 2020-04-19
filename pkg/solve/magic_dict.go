package solve

type MagicDictionary struct {
	words map[string]bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	words := make(map[string]bool)
	return MagicDictionary{words}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, w := range dict {
		this.words[w] = true
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	for w, _ := range this.words {
		if len(w) == len(word) {
			kifkif := 0
			for i, _ := range w {
				if w[i] == word[i] {
					kifkif++
				}
			}
			if kifkif == len(word)-1 {
				return true
			}
		}
	}

	return false
}
