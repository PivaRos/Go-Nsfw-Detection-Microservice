package structs

type ImageUploadMessage struct {
	Base64    string `json:"base64"`
	ImageGuid string `json:"imageGuid"`
}
