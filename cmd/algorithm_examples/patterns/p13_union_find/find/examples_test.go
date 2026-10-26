package main

import (
	"testing"
)

func TestNewDSU(t *testing.T) {
	d := NewDSU(5)
	if got := d.Components(); got != 5 {
		t.Errorf("изначально компонент должно быть 5, got %d", got)
	}
	for i := 0; i < 5; i++ {
		if d.Find(i) != i {
			t.Errorf("изначально каждый элемент — свой корень, d.Find(%d) = %d", i, d.Find(i))
		}
		if d.Size(i) != 1 {
			t.Errorf("изначально размер каждой компоненты = 1, got %d", d.Size(i))
		}
	}
}

func TestUnionAndSame(t *testing.T) {
	// Связность: 0-1-2, 3-4.
	d := NewDSU(5)
	d.Union(0, 1)
	if !d.Same(0, 1) {
		t.Fatal("0 и 1 должны быть в одной компоненте")
	}
	if d.Same(0, 2) {
		t.Fatal("0 и 2 ещё не соединены")
	}
	d.Union(1, 2)
	if !d.Same(0, 2) {
		t.Fatal("после 1-2 все {0,1,2} должны быть вместе")
	}
	if d.Same(0, 3) {
		t.Fatal("3 — отдельная компонента")
	}
	d.Union(3, 4)
	if !d.Same(3, 4) {
		t.Fatal("3 и 4 вместе")
	}
	if d.Same(0, 3) {
		t.Fatal("0 и 3 всё ещё разделены")
	}
	// Компоненты: {0,1,2}, {3,4}.
	comps := d.Components()
	if comps != 2 {
		t.Errorf("ожидалось 2 компоненты, got %d", comps)
	}
	// Попытка объединить уже внутри.
	if d.Union(0, 2) {
		t.Fatal("0 и 2 уже вместе, Union должен вернуть false")
	}
}

func TestRankAndSize(t *testing.T) {
	d := NewDSU(6)
	d.Union(0, 1)
	d.Union(1, 2)
	if got := d.Size(0); got != 3 {
		t.Errorf("размер компоненты после 0-1-2 должен быть 3, got %d", got)
	}
	// соединяем маленькую с большой: 3-0, маленькая подчинится.
	d.Union(3, 0)
	if got := d.Size(3); got != 4 {
		t.Errorf("размер компоненты после вклеивания 3 должен быть 4, got %d", got)
	}
	if d.Same(3, 0) {
		t.Skip("ранг поведён так, что 3 и 0 вместе — ок")
	} else {
		t.Fatal("3 стал в одной компоненте с {0,1,2}")
	}
	// изолированный 4.
	if d.Same(4, 0) {
		t.Fatal("4 изначально изолирован")
	}
	if d.Size(4) != 1 {
		t.Errorf("изолированный элемент остаётся размером 1, got %d", d.Size(4))
	}
}

func TestFindCompressionDoesNotOscillate(t *testing.T) {
	// После сжатия повторный find возвращает тот же корень.
	d := NewDSU(10)
	for i := 1; i < 10; i++ {
		d.Union(0, i)
	}
	r1 := d.Find(9)
	r2 := d.Find(9)
	if r1 != r2 {
		t.Errorf("после сжатия Find должен быть детерминирован: %d != %d", r1, r2)
	}
	if r1 != d.Find(0) {
		t.Error("корень 9 должен совпадать с корнем 0")
	}
}

func TestComponentsDecreaseOnlyOnSuccessfulUnion(t *testing.T) {
	d := NewDSU(4)
	if d.Components() != 4 {
		t.Fatal("изначально 4")
	}
	d.Union(0, 1)
	if d.Components() != 3 {
		t.Errorf("после успешного merge должно быть 3, got %d", d.Components())
	}
	before := d.Components()
	d.Union(0, 1)
	if after := d.Components(); after != before {
		t.Errorf("дубликат merge должен не менять количество компонент: %d -> %d", before, after)
	}
}

func TestEmptyDSU(t *testing.T) {
	d := NewDSU(0)
	if d.Components() != 0 {
		t.Errorf("пустой DSU: ожидалось 0 компонент, got %d", d.Components())
	}
	// к искать нечего — Find на пустом слайсе не должен падать, но
	// вызов с x>=0 на пустом DSU технически невалиден. Проверим, что
	// код не паникует на создание.
	_ = d
}

func TestSingleElement(t *testing.T) {
	d := NewDSU(1)
	if d.Components() != 1 {
		t.Fatalf("должна быть 1 компонента, got %d", d.Components())
	}
	if !d.Same(0, 0) {
		t.Fatal("один и тот же элемент — в одной компоненте")
	}
	if d.Size(0) != 1 {
		t.Errorf("размер единичного элемента = 1, got %d", d.Size(0))
	}
}

func TestComponentsNotConnectedToSelf(t *testing.T) {
	// Убедимся, что компоненты не «сливаются сами в себя» без явного merge.
	d := NewDSU(3)
	d.Union(0, 1)
	if d.Same(0, 2) {
		t.Fatal("0 и 2 не должны быть соединены")
	}
	d.Union(1, 2)
	if !d.Same(0, 2) {
		t.Fatal("после 1-2 всё должно быть вместе")
	}
}
