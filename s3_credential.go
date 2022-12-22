package idcloudhost

type S3Credential struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	UserID    string `json:"userId"`
}
