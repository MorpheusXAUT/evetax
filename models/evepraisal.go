package models

// EvePraisal represents a retrieved evepraisal and all values
type EvePraisal struct {
	Created    int64            `json:"created"`
	ID         int64            `json:"id"`
	Items      []EvePraisalItem `json:"items"`
	Kind       string           `json:"kind"`
	MarketID   int64            `json:"market_id"`
	MarketName string           `json:"market_name"`
	Totals     EvePraisalTotals `json:"totals"`
}

// EvePraisalItem represents one item as retrieved from evepraisal
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

// EvePraisalPrices represents all three prices of one EvePraisalItem
type EvePraisalPrices struct {
	All  EvePraisalPrice `json:"all"`
	Buy  EvePraisalPrice `json:"buy"`
	Sell EvePraisalPrice `json:"sell"`
}

// EvePraisalPrice represents a single price including average, minimum and maximum
type EvePraisalPrice struct {
	Average float64 `json:"avg"`
	Maximum float64 `json:"max"`
	Minimum float64 `json:"min"`
	Price   float64 `json:"price"`
}

// EvePraisalTotals represents the total of all prices of the EvePraisal and all items
type EvePraisalTotals struct {
	Buy    float64 `json:"buy"`
	Sell   float64 `json:"sell"`
	Volume float64 `json:"volume"`
}

// GetTotalBuyValue fetches the total buy value of the EvePraisal
func (e EvePraisal) GetTotalBuyValue() float64 {
	return e.Totals.Buy
}

// GetTotalSellValue fetches the total sell value of the EvePraisal
func (e EvePraisal) GetTotalSellValue() float64 {
	return e.Totals.Sell
}

// GetTotalVolume fetches the total volume of the EvePraisal
func (e EvePraisal) GetTotalVolume() float64 {
	return e.Totals.Volume
}
