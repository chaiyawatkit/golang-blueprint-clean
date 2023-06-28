package env

import (
	"os"
)

var (
	AppFeatureServiceUrl string
	RsaPublicKey         string
	EncryptKey           string
	XCorrelationID       string
)

func Init() {

	os.Setenv("AppFeatureServiceUrl", "http://localhost:8080")
	os.Setenv("RSA_PUBLIC_KEY", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3KWCjLIptcxlOUeGBE/8\n6m+oaufsDHVP4RLbj+5BQDxr16vFphaMz7WKUiwb2tu/CFUaVdPmXAG3O9klFCI/\nI8XpNxyn1FQ3/6hDijYXKNSXOS1YGZKLs81/tQb2YxmTvTk0wn55OdHgvUXizeBI\ngAqLhjvVDw5MB+lC0LV/onO9m5byyjHqY33hbEFdzXpajVfs5OTMTWIlDxKYNr4Z\nK2pcHtM63A0qKcnMN2+S4EodirOciQRNvEcnn1O4n3QeB04MaWiMmoeY4pLxFb2f\nl2PQU+NYLGCCtJAJF+znBltwN0V4TIfNAEdLhMYToo9O07ode86ilzVr9W6C8gi/\nPQIDAQAB\n-----END PUBLIC KEY-----")
	os.Setenv("ENCRYPT_KEY", "b054eb59dcf46db5da45ade306d005a2473def4a51f0ee93371bd64e77ae4b20")
	os.Setenv("X_CORRELATION_ID", "ad1ad1b903a4711506a2bfd6a8fd9086d2aaee36fc267b9be847963b9412b95e")

	AppFeatureServiceUrl = os.Getenv("AppFeatureServiceUrl")
	RsaPublicKey = os.Getenv("RSA_PUBLIC_KEY")
	EncryptKey = os.Getenv("ENCRYPT_KEY")
	XCorrelationID = os.Getenv("X_CORRELATION_ID")

}
