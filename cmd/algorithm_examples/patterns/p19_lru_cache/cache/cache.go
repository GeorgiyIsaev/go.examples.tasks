package main

// Паттерн: LRU Cache — кэш с вытеснением «наиболее старый» за O(1).
//
// Ключевая идея: хэш-мапа key → узел + двусвязный список (most-recent →
// least-recent). При Get/Put touched-ключ перемещается в front; при превышении
// capacity вытесняется узел перед tail.
//
// Инвариант: в кэше всегда ≤ capacity элементов; мапа даёт прямой указатель
// на любой узел, поэтому move/remove — O(1) без линейного поиска.
// Фиктивные head/tail убирают ветки «пустой список».

// node — элемент кэша в двусвязном списке.
type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

// LRUCache — LRU-кэш за O(1): Get/Put, eviction.
type LRUCache struct {
	cap   int
	head  *node
	tail  *node
	byKey map[int]*node
}

// NewLRUCache создаёт кэш заданной ёмкости; при capacity ≤ 0 — пустой кэш.
func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		capacity = 0
	}
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		cap:   capacity,
		head:  head,
		tail:  tail,
		byKey: make(map[int]*node, capacity),
	}
}

// Get возвращает значение по ключу; -1 если ключа нет. Touched-ключ → front.
func (c *LRUCache) Get(key int) int {
	n, ok := c.byKey[key]
	if !ok {
		return -1
	}
	c.moveToFront(n)
	return n.value
}

// Put вставляет или обновляет значение; при переполнении evict LRU.
func (c *LRUCache) Put(key int, value int) {
	if c.cap <= 0 {
		return
	}
	if n, ok := c.byKey[key]; ok {
		n.value = value
		c.moveToFront(n)
		return
	}
	n := &node{key: key, value: value}
	c.byKey[key] = n
	c.insertAtFront(n)
	if len(c.byKey) > c.cap {
		c.evictLRU()
	}
}

// moveToFront — отцепить и вставить после head (most recent).
func (c *LRUCache) moveToFront(n *node) {
	c.remove(n)
	c.insertAtFront(n)
}

// insertAtFront — вставить n сразу после head.
func (c *LRUCache) insertAtFront(n *node) {
	n.prev = c.head
	n.next = c.head.next
	c.head.next.prev = n
	c.head.next = n
}

// remove — отцепить n от текущего места; не трогает мапу.
func (c *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

// evictLRU — удалить least-recent (узел перед tail) из списка и из мапы.
func (c *LRUCache) evictLRU() {
	lru := c.tail.prev
	if lru == c.head {
		return
	}
	c.remove(lru)
	delete(c.byKey, lru.key)
}
