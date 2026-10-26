package main

// Паттерн: Trie (префиксное дерево) — «мапа, у которой ключ
// разложен по буквам».
//
// Ключевая идея: путь от корня кодирует префикс. Каждый узел отвечает
// на вопрос «продолжается ли уже виденный префикс?», а флаг isWord
// отличает «просто префикс» от «законченного слова».
//
// Сложность: Insert/Search/StartsWith — O(L), L = длина слова
// (вне зависимости от числа слов в словаре!).
// Память: до O(сумма длин × 26 указателей) — дороже хэш-мапы,
// зато автодополнение и группировка по префиксу «бесплатно».

type TrieNode struct {
	children [26]*TrieNode // только a..z, как в условиях LeetCode
	isWord   bool          // здесь заканчивается полное слово
}

type Trie struct {
	root TrieNode // значение, а не указатель: не нужна проверка nil
}

func NewTrie() *Trie { return &Trie{} }

// letter — индекс буквы a..z; ok=false для всего остального.
func letter(b byte) (int, bool) {
	if b < 'a' || b > 'z' {
		return 0, false
	}
	return int(b - 'a'), true
}

// Insert добавляет word. Повторная вставка безопасна: слово уже
// помечено, ничего не ломается.
func (t *Trie) Insert(word string) {
	node := &t.root
	for i := 0; i < len(word); i++ {
		c, ok := letter(word[i])
		if !ok {
			return // не a..z — в условиях LC такого не бывает
		}
		if node.children[c] == nil {
			node.children[c] = &TrieNode{}
		}
		node = node.children[c]
	}
	node.isWord = true
}

// Search — есть ли слово целиком (не просто префикс).
func (t *Trie) Search(word string) bool {
	node := t.walk(word)
	return node != nil && node.isWord
}

// StartsWith — есть ли хоть одно слово с таким префиксом.
func (t *Trie) StartsWith(prefix string) bool {
	return t.walk(prefix) != nil
}

// walk — узел по данному пути или nil, если пути нет.
func (t *Trie) walk(path string) *TrieNode {
	node := &t.root
	for i := 0; i < len(path); i++ {
		c, ok := letter(path[i])
		if !ok {
			return nil
		}
		node = node.children[c]
		if node == nil {
			return nil
		}
	}
	return node
}

// WordsWithPrefix — автодополнение: все слова с префиксом,
// в алфавитном порядке (обход поддерева слева направо).
//
// Три шага: спуститься к узлу префикса (O(|prefix|)), затем DFS
// собрать все isWord под ним. Пустой префикс = все слова словаря.
func (t *Trie) WordsWithPrefix(prefix string) []string {
	node := t.walk(prefix)
	if node == nil {
		return nil
	}
	var out []string
	t.collect(node, prefix, &out)
	return out
}

func (t *Trie) collect(node *TrieNode, base string, out *[]string) {
	if node.isWord {
		*out = append(*out, base)
	}
	for c := 0; c < 26; c++ {
		if node.children[c] != nil {
			t.collect(node.children[c], base+string(rune('a'+c)), out)
		}
	}
}
