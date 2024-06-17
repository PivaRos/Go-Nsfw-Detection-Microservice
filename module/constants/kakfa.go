package constants

type KafkaTopic string

const (
	ImageUpload KafkaTopic = "image_upload"
)

var Topics = []string{string(ImageUpload)}
