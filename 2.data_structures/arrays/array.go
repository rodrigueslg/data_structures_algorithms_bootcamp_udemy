package arrays

type Array struct {
	Lenght int
	Data   map[int]any
}

func NewArray() *Array {
	return &Array{
		Lenght: 0,
		Data:   make(map[int]any),
	}
}

func (a *Array) Get(index int) any {
	if a.Lenght > 0 {
		return a.Data[index]
	}
	return nil
}

func (a *Array) Push(item any) int {
	a.Data[a.Lenght] = item
	a.Lenght++
	return a.Lenght
}

func (a *Array) Pop() any {
	if a.Lenght > 0 {
		lastItem := a.Data[a.Lenght-1]
		delete(a.Data, a.Lenght-1)
		a.Lenght--
		return lastItem
	}
	return nil
}

func (a *Array) Delete(index int) {
	if a.Lenght > 0 {
		a.shiftItems(index)
		a.Lenght--
	}
}

func (a *Array) shiftItems(index int) {
	for i := index; i < a.Lenght-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	delete(a.Data, a.Lenght-1)
}

func (a *Array) Reverse() {
	reversed := make(map[int]any)
	j := 0
	for i := a.Lenght - 1; i >= 0; i-- {
		reversed[j] = a.Data[i]
		j++
	}
	a.Data = reversed
}
