package common

import "os"

func GetProToken() string {
	return os.Getenv("eb_pro_token")
}
