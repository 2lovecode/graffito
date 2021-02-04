package pkg

type Stock struct {
	sku *Sku
}

func NewStock(sku *Sku) *Stock {
	return &Stock{
		sku: sku,
	}
}

// 预锁定
func (s *Stock) PreLock(num uint64) {

}

// 添加库存
func (s *Stock) Add(num uint64) {

}

// 扣减库存
func (s *Stock) Deduct(num uint64) {

}

// 库存回退
func (s *Stock) Rollback(num uint64) {

}
