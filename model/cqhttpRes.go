package model

type CqhttpRes struct {
	Data struct {
		MessageID int `json:"message_id"`
	} `json:"data"`
	Echo    string `json:"echo"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}
