package tzkt

import "time"

type DelegationItems struct {
	Items   []*DelegationItem
	HasMore bool
}

type DelegationItem struct {
	Type         string      `json:"type"`
	ID           int64       `json:"id"`
	Level        int         `json:"level"`
	Timestamp    time.Time   `json:"timestamp"`
	Block        string      `json:"block"`
	Hash         string      `json:"hash"`
	Counter      int         `json:"counter"`
	Sender       Sender      `json:"sender"`
	GasLimit     int         `json:"gasLimit"`
	GasUsed      int         `json:"gasUsed"`
	StorageLimit int         `json:"storageLimit"`
	BakerFee     int         `json:"bakerFee"`
	Amount       int64       `json:"amount"`
	NewDelegate  NewDelegate `json:"newDelegate"`
	Status       string      `json:"status"`
}

type Sender struct {
	Address string `json:"address"`
}

type NewDelegate struct {
	Alias   string `json:"alias"`
	Address string `json:"address"`
}
