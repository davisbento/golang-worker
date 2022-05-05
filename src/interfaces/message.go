package interfaces

type Message struct {
	UserId     string `json:"userId"`
	ItemId     string `json:"itemId"`
	Type       string `json:"type"`
	ValueInBtc float64  `json:"valueInBtc"`
	Timestamp  string `json:"@timestamp"`
}
