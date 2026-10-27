package main

import (
	"testing"
)

func TestNewLRUAlgorithm(t *testing.T) {
	c := NewLRUCache(2)

	// базовая вставка и чтение
	c.Put(1, 1)
	if got := c.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	c.Put(2, 2)
	if got := c.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}

	// тот же ключ — обновление
	c.Put(1, 10)
	if got := c.Get(1); got != 10 {
		t.Errorf("после обновления Get(1) = %d, want 10", got)
	}
}

func TestLRUEviction(t *testing.T) {
	c := NewLRUCache(2)

	c.Put(1, 1)
	c.Put(2, 2)
	// размер == 2, дальше вытеснение
	c.Put(3, 3) // вытеснит LRU (1, если 1 было не touched)

	if c.Get(1) != -1 {
		t.Fatal("ключ 1 должен быть вытеснен")
	}
	if got := c.Get(2); got != 2 {
		t.Errorf("ключ 2 должен остаться, got %d", got)
	}
	if got := c.Get(3); got != 3 {
		t.Errorf("ключ 3 должен быть в кэше, got %d", got)
	}

	// касание 2 делает его recent; затем 4 вытеснит 3 (если это LRU)
	c.Get(2)
	c.Put(4, 4)

	if c.Get(3) != -1 {
		t.Fatal("ключ 3 должен быть вытеснен после того, как 2 стал recent")
	}
	if got := c.Get(2); got != 2 {
		t.Errorf("ключ 2 должен остаться, got %d", got)
	}
	if got := c.Get(4); got != 4 {
		t.Errorf("ключ 4 должен быть в кэше, got %d", got)
	}
}

func TestLRUCapacityZero(t *testing.T) {
	c := NewLRUCache(0)
	c.Put(1, 1)
	if c.Get(1) != -1 {
		t.Errorf("при ёмкости 0 Get должен возвращать -1, got %d", c.Get(1))
	}
}

func TestLRUGetNonexistent(t *testing.T) {
	c := NewLRUCache(5)
	if got := c.Get(99); got != -1 {
		t.Errorf("несуществующий ключ должен вернуть -1, got %d", got)
	}
}

func TestLRUUpdateDoesNotIncreaseSize(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(1, 10)
	// размер всё ещё 2, ключей 2
	if c.Get(1) != 10 {
		t.Errorf("обновлённое значение должно быть 10, got %d", c.Get(1))
	}
	if c.Get(2) != 2 {
		t.Errorf("второй ключ должен остаться, got %d", c.Get(2))
	}
}

func TestLRUFrontRecent(t *testing.T) {
	// тот, что тронут позже, остаётся, а LRU — вытесняется.
	c := NewLRUCache(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)

	c.Get(1)    // touch 1 → 1 becomes most recent
	c.Put(4, 4) // вытеснит LRU, т.е. 2 (если 2 не touched)

	if c.Get(2) != -1 {
		t.Fatal("2 должен быть вытеснен (LRU)")
	}
	if c.Get(1) != 1 {
		t.Errorf("1 должен остаться (most recent), got %d", c.Get(1))
	}
	if c.Get(3) != 3 {
		t.Errorf("3 должен остаться, got %d", c.Get(3))
	}
	if c.Get(4) != 4 {
		t.Errorf("4 должен быть в кэше, got %d", c.Get(4))
	}
}

func TestLRUCapacityOne(t *testing.T) {
	c := NewLRUCache(1)
	c.Put(1, 100)
	if c.Get(1) != 100 {
		t.Errorf("единственный элемент должен храниться, got %d", c.Get(1))
	}
	c.Put(2, 200)
	if c.Get(1) != -1 {
		t.Fatal("ключ 1 должен быть вытеснен, когда пришел новый")
	}
	if c.Get(2) != 200 {
		t.Errorf("ключ 2 должен быть в кэше, got %d", c.Get(2))
	}
}

func TestLRUOrderReuse(t *testing.T) {
	c := NewLRUCache(2)

	c.Put(1, 11)
	c.Put(2, 22)
	c.Get(1)     // 1 recent
	c.Put(3, 33) // evict 2

	if c.Get(2) != -1 {
		t.Fatal("2 должен быть вытеснен")
	}
	if c.Get(1) != 11 {
		t.Errorf("1 должен остаться, got %d", c.Get(1))
	}
	if c.Get(3) != 33 {
		t.Errorf("3 должен быть в кэше, got %d", c.Get(3))
	}
}
