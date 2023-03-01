package web

type PiggyBankReponse struct {
	Id            string  `json:"id"`
	UserId        string  `json:"user_id"`
	PiggyBankName string  `json:"piggy_bank_name"`
	Type          bool    `json:"type"`
	Total         float32 `json:"total"`
}
