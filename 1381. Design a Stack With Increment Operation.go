/**
 * Придумать свой стэк со следующими операциями.
 * Конструктор с параметром maxSize int - махимальное кол-во элементов в стэке. Если размер стека достигает этого
 * значения, дальнейшего увеличение не происходит.
 * func push(int x) Добавляет х на вершину стэка, если размер стека не должен превышать максимум.
 * func pop() int Извлекает и возвращает элемент с вершины или -1 если стэк пуст.
 * func Increment(k int, val int) Увеличивает значения последних k элементов стэка на val.
 * Если в стэке меньше чем k элементов, просто увеличиваем все элементы в стэке.
 * Инициализируем:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

package leetcode

type CustomStack struct {
	s       []int
	maxSize int
}

func New(maxSize int) CustomStack {
	return CustomStack{
		s:       make([]int, 0, maxSize),
		maxSize: maxSize,
	}
}

func (cs *CustomStack) Push(x int) {
	if len(cs.s) < cap(cs.s) {
		cs.s = append(cs.s, x)
	}
}

func (cs *CustomStack) Pop() int {
	if len(cs.s) == 0 {
		return -1
	}
	var x int

	x, cs.s = cs.s[len(cs.s)-1], cs.s[:len(cs.s)-1]

	return x
}

func (cs *CustomStack) Increment(k, val int) {
	if k > len(cs.s) {
		k = len(cs.s)
	}

	for i := 0; i < k; i++ {
		cs.s[i] += val
	}
}
