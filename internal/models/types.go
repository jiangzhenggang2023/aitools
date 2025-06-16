package models

type UpdateData struct {
	KBID        string   `json:"kb_id"`
	AccessToken string   `json:"access_token"`
	Files       []string `json:"file_names"`
}

type ResponseData struct {
	RetCode int    `json:"retcode"`
	RetMsg  string `json:"retmsg"`
}
