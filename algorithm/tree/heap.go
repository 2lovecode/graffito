package tree

// 数组实现最大堆
type MaxIntHeap struct {
	Data      []int
	Len       int
	Cap       int
	OriginCap int
}

func NewIntHeap(cap int) *MaxIntHeap {
	heap := &MaxIntHeap{
		Data:      make([]int, cap),
		Len:       0,
		Cap:       cap,
		OriginCap: cap,
	}
	return heap
}

func (heap *MaxIntHeap) Push(value int) {
	if (heap.Len + 1) >= heap.Cap {
		data := make([]int, heap.Cap*2)
		copy(data, heap.Data)
		heap.Data = data
		heap.Cap = heap.Cap * 2
	}
	//新值放到最后
	heap.Data[heap.Len] = value
	heap.Len++

	//上浮操作
	heap.up()
}

func (heap *MaxIntHeap) Pop() int {
	value := heap.Data[0]

	heap.Len--

	heap.Data[0] = heap.Data[heap.Len]
	heap.Data[heap.Len] = 0

	if (heap.Len == (heap.Cap - heap.OriginCap)) && (heap.Len != 0) {
		heap.Data = heap.Data[0:heap.Len]
		heap.Cap = heap.Cap - heap.OriginCap
	}

	heap.down()

	return value
}

func (heap *MaxIntHeap) up() {
	currentKey := heap.Len - 1
	parentKey := (currentKey+1)/2 - 1
	for parentKey >= 0 {
		if heap.Data[parentKey] < heap.Data[currentKey] {
			heap.Data[parentKey], heap.Data[currentKey] = heap.Data[currentKey], heap.Data[parentKey]
			currentKey = parentKey
			parentKey = (currentKey+1)/2 - 1
		} else {
			break
		}
	}
}

func (heap *MaxIntHeap) down() {
	currentKey := 0
	leftKey := currentKey*2 + 1
	rightKey := currentKey*2 + 2

	for (currentKey < heap.Len) && (rightKey <= heap.Len) {
		tmpKey := leftKey
		if heap.Data[leftKey] < heap.Data[rightKey] {
			tmpKey = rightKey
		}
		if heap.Data[currentKey] < heap.Data[tmpKey] {
			heap.Data[currentKey], heap.Data[tmpKey] = heap.Data[tmpKey], heap.Data[currentKey]
		} else {
			break
		}
		currentKey = tmpKey
		leftKey = currentKey*2 + 1
		rightKey = currentKey*2 + 2
	}
}
