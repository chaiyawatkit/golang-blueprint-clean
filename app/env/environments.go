package env

import (
	"os"
)

var (
	OracleUser     string
	OraclePassword string
	OracleHost     string
	OraclePort     string
	OracleService  string
	RsaPublicKey   string
	RsaPrivateKey  string
	EncryptKey     string
)

func Init() {

	os.Setenv("ORACLE_USER", "system")
	os.Setenv("ORACLE_PASSWORD", "tiger")
	os.Setenv("ORACLE_HOST", "localhost")
	os.Setenv("ORACLE_PORT", "1521")
	os.Setenv("ORACLE_SERVICE", "orcl")
	os.Setenv("RSA_PUBLIC_KEY", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3KWCjLIptcxlOUeGBE/8\n6m+oaufsDHVP4RLbj+5BQDxr16vFphaMz7WKUiwb2tu/CFUaVdPmXAG3O9klFCI/\nI8XpNxyn1FQ3/6hDijYXKNSXOS1YGZKLs81/tQb2YxmTvTk0wn55OdHgvUXizeBI\ngAqLhjvVDw5MB+lC0LV/onO9m5byyjHqY33hbEFdzXpajVfs5OTMTWIlDxKYNr4Z\nK2pcHtM63A0qKcnMN2+S4EodirOciQRNvEcnn1O4n3QeB04MaWiMmoeY4pLxFb2f\nl2PQU+NYLGCCtJAJF+znBltwN0V4TIfNAEdLhMYToo9O07ode86ilzVr9W6C8gi/\nPQIDAQAB\n-----END PUBLIC KEY-----")
	os.Setenv("RSA_PRIVATE_KEY", "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDcpYKMsim1zGU5\nR4YET/zqb6hq5+wMdU/hEtuP7kFAPGvXq8WmFozPtYpSLBva278IVRpV0+ZcAbc7\n2SUUIj8jxek3HKfUVDf/qEOKNhco1Jc5LVgZkouzzX+1BvZjGZO9OTTCfnk50eC9\nReLN4EiACouGO9UPDkwH6ULQtX+ic72blvLKMepjfeFsQV3NelqNV+zk5MxNYiUP\nEpg2vhkralwe0zrcDSopycw3b5LgSh2Ks5yJBE28RyefU7ifdB4HTgxpaIyah5ji\nkvEVvZ+XY9BT41gsYIK0kAkX7OcGW3A3RXhMh80AR0uExhOij07Tuh17zqKXNWv1\nboLyCL89AgMBAAECggEBAMDaoUrchnxyVJ9GFqs6RrOaOLiIUddpyhHN677X3oyI\nwzHU34KV/ymlYWsFRqt6mnRr/gK6nYQZs9BJsZDxhgJ8f6U+FKsPmYqJxc7ODFwR\nxMOjEHf5kMXtfgv6M+zWdD1V/PlqlgalnSuSI/USbgwMK3t+XHE2a9oWVeIaYZ7Y\nr3b1vT8LLsxN3ns41c2fsmdN57rL3dIMPyDBDhztZGgFBle3SyRHbWP9pfrTFVrb\ns+6AH33B9kpfNUFJUWuTpn4JKbFKdQpAXOzzP+WX1GMwdhFWnKavvp+2ug6g5mn9\nAap3t5pekNszUMRZQFOxp3YrTHTiYFRmGTHxHLRg+nkCgYEA8t9buVp4tcV6Wn0b\n9Ti2Zw4gtHXZ9aTAwlJ7Rs5gmBbHyjsDmSu7LKTN6a/riHD7PpN093QrKs8992Wk\nmpsUDtmMPWSxXw+mR46l3PVJvmZAloSN3LInkkMRxNlIuOYDi+P+mQHIqnZTbSnI\nxCFSgYTe9/Lp1Q6RHnH6elxHCq8CgYEA6JKcCtsUJNNfP2xkSNAE7E5d6xsfgD6G\nOeAjO7bz0k4THleNzDNuxd0Z8b14nezE9Zr+MPchiTIal8Ez+NbGyT9UyAJBwTwX\nnwNbJjI0iNLxU7Rx1Z9nWYF6SkNUYkbnzI1MNUJxNv/iBys9ybjXpI81O3fmAQzf\nRSfzf7K8X9MCgYBE79UvNVjR722dRamr3x8W/VGXJ2Robw0vmw3WuTl9semfo0SQ\nM3N7ZoPz8rUeE0OdLWmj21ttWUmzcSxZne8BkApYWr5lxyFhakH2B8GYw5lNn+5M\nWF0XDOZ8Q1h68v7KtDpN382/ZLqlX5jW4sZycO6A3ng0u1/BjoSUOqEHjQKBgDcz\nTQwakX5oGOxRyr+pd/GqbAo4ZtKmhG9KLrKmqtpYo5sJBuPXtUwLPXQuF2nMX7c3\nsWtnzLLmjH9GYRKHz6jGelX6iybVH3ojbsfzFJsBDjxi0L32Vq5c1+y2bFnjIHNh\nehwOCBcYxsoSzliJoc0yHF1gCNxo9LlSUS4W+zT3AoGBAMEwWkg8OjeiGuoHYySI\nG/rIZ0NSiIrkoo6HxsAENjO2t9iLbUsbcA1zFhwekUCgx9ROqtpEbGekgKeTALiM\nr0rLyCQ7i8XocVjen5qPhWQPe8WsfNoThLc1H/qCzESEQ6BEdd5B+OFtkcnibDGu\nyf2JkpeeoVZ+eA7CtGhQGyF2\n-----END PRIVATE KEY-----")
	os.Setenv("ENCRYPT_KEY", "b054eb59dcf46db5da45ade306d005a2473def4a51f0ee93371bd64e77ae4b20")

	OracleUser = os.Getenv("ORACLE_USER")
	OraclePassword = os.Getenv("ORACLE_PASSWORD")
	OracleHost = os.Getenv("ORACLE_HOST")
	OraclePort = os.Getenv("ORACLE_PORT")
	OracleService = os.Getenv("ORACLE_SERVICE")
	RsaPublicKey = os.Getenv("RSA_PUBLIC_KEY")
	RsaPrivateKey = os.Getenv("RSA_PRIVATE_KEY")
	EncryptKey = os.Getenv("ENCRYPT_KEY")

}
