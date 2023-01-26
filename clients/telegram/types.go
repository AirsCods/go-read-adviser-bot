package telegram

// UpdatesResponse Структура полученная из request
type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update Структура получения сообщений от telegram api
type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}
