package models

type EvePraisal struct {
	Created    int64            `json:"created"`
	ID         int64            `json:"id"`
	Items      []EvePraisalItem `json:"items"`
	Kind       string           `json:"kind"`
	MarketID   int64            `json:"market_id"`
	MarketName string           `json:"market_name"`
	Totals     EvePraisalTotals `json:"totals"`
}

type EvePraisalItem struct {
	GroupID  int64            `json:"groupID"`
	Market   bool             `json:"market"`
	Name     string           `json:"name"`
	Prices   EvePraisalPrices `json:"prices"`
	Quantity int64            `json:"quantity"`
	TypeID   int64            `json:"typeID"`
	TypeName string           `json:"typeName"`
	Volume   float64          `json:"volume"`
}

type EvePraisalPrices struct {
	All  EvePraisalPrice `json:"all"`
	Buy  EvePraisalPrice `json:"buy"`
	Sell EvePraisalPrice `json:"sell"`
}

type EvePraisalPrice struct {
	Average float64 `json:"avg"`
	Maximum float64 `json:"max"`
	Minimum float64 `json:"min"`
	Price   float64 `json:"price"`
}

type EvePraisalTotals struct {
	Buy    float64 `json:"buy"`
	Sell   float64 `json:"sell"`
	Volume float64 `json:"volume"`
}

func (e EvePraisal) GetTotalBuyValue() float64 {
	return e.Totals.Buy
}

func (e EvePraisal) GetTotalSellValue() float64 {
	return e.Totals.Sell
}

func (e EvePraisal) GetTotalVolume() float64 {
	return e.Totals.Volume
}
