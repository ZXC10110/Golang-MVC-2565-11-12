package Models

func (b *Product) TableName() string {
	return "product"
}

type Product struct {
	Sequence int     `json:"sequence"`
	Pid      string  `json:"pid"`
	Pname    string  `json:"pname"`
	Price    float64 `json:"price"`
	Qty      int     `json:"qty"`
}
