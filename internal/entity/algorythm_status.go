package entity

type AlgorithmStatus struct {
	ID       int64
	ClientID int64
	VWAP     bool
	TWAP     bool
	HFT      bool
}
