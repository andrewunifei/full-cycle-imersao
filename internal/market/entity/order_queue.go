package entity

type OrderQueue struct {
	Orders []*Order
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}

// Less - Compara i e j
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Swap - Inverte i por j
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Len - Tamanho do dado
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Push - Adiciona o objeto
func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Pop - Remove o último objeto e o retorna
func (oq *OrderQueue) Pop() any {
	old := oq.Orders
	n := len(old)
	item := old[n - 1]
	oq.Orders = old[0: n - 1]

	return item
}