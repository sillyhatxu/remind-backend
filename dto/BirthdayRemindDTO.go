package dto

type BirthdayRemindDTO struct {
	Id         string `json:"id"`
	Fallback   string `json:"fallback"`
	Text       string `json:"text"`
	AuthorName string `json:"author_name"`
}
