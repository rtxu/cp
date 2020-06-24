type Trie struct {
    isWordBoundary bool
    Children [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    root := &Trie{}
    return *root
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    n := len(word)
    current := this
    for i := 0; i < n; i++ {
        childIdx := byte(word[i] - 'a')
        if current.Children[childIdx] == nil {
            current.Children[childIdx] = &Trie{}
        }
        current = current.Children[childIdx]
    }
    current.isWordBoundary = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    n := len(word)
    current := this
    for i := 0; i < n; i++ {
        childIdx := byte(word[i] - 'a')
        if current.Children[childIdx] == nil {
            return false
        }
        current = current.Children[childIdx]
    }
    return current.isWordBoundary
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    n := len(prefix)
    current := this
    for i := 0; i < n; i++ {
        childIdx := byte(prefix[i] - 'a')
        if current.Children[childIdx] == nil {
            return false
        }
        current = current.Children[childIdx]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
