package param

type AuthParam struct {
	Key       string `json:"key"`
	Signature string `json:"signature"`
	Type      int64  `json:"type"`
}
