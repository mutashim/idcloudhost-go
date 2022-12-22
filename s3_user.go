package idcloudhost

type S3User struct {
	Caps             interface{}    `json:"caps"`
	DisplayName      string         `json:"displayName"`
	Email            string         `json:"email"`
	MaxBuckets       int64          `json:"maxBuckets"`
	S3Credentials    []S3Credential `json:"s3Credentials"`
	Subusers         interface{}    `json:"subusers"`
	Suspended        int64          `json:"suspended"`
	SwiftCredentials interface{}    `json:"swiftCredentials"`
	UserID           string         `json:"userId"`
}
