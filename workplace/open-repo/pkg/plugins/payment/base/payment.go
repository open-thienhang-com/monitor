package payment

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"
	"time"
)

type IPayment interface {
	CaptureWallet() (interface{}, error)
	GetInvoiceID() string
	GetAmount() string
}

var pubKeyData = []byte(`
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA49mGWEzZt4KD9O0nsVoG6ibZ6C6rZtkTOCuXuBYoG2eAhFZL6u/A5JRvtSONcficyopXq2DVlY3RdTLq3+MFWF46KBcnWs7ThNZB7Vy1UmhwpG7WOg7YOBJ2BydLFhs5Y+yBMAsEBY9E7BwkLl3S+DFpC7vI0NJWrd8yHCCOfAY5uWGqTp7MNpSzwRu3YCa9jzs9H61EsyEbCeodtsX+VrlT0qhmh9P5P99zLMFuRkYceLpo7B0P7KQ9P3H3IvtkIJ0YnE2ltFRlWAi5bCjEK0qoiF6p4cnlyXPG2xq/KHSdXPKwHLcHulsqtJWK9WFu45iW+sr8ZU5oZU4eg5+BY/IzacVZqqCkLzCjeqy1JHMduBpDmzoR5w7BkOO1clQ1O82NhVYSzfkxFBjPW9AMZACiHBfqgg5ESjKTG3VhgkQpmnfGi0+tnGg56m6yzPl30Vkv1MYnu8eNR7atRpEbcktiBezvFzokvQCFPYKBIME2KC2XRb4B8pDhh/4XxZThl9xquI3itdrPD/rF/G4oUOmkpprF3hryKm3cKajC5iUO9qeAL6Z3qjKBJSa4MfMz87JvKbLOUXYyP+um0ioDz8hpeTsB+mfdx70rGGONtB+PCxAMnhYC+9uYNqAhPDHpSJtWsjUV1x8zf6pQKGpBN5NdH1tzF1lujE77+3VyHzkCAwEAAQ==
-----END PUBLIC KEY-----
`)

type Payment struct {
	OrderInfo string `json:"order_info"`
	Amount    string `json:"amount"`
	BankCode  string `json:"bank_code"`
	OrderType string `json:"order_type"`
	VnpUrl    string
	ReturnUrl string
	IpAddr    string
	SecretKey string
	TmnCode   string
}

func (pay Payment) GeneratePaymentUrl() string {
	currentDate := time.Now()
	m := map[string]string{}

	m["vnp_Version"] = "2"
	m["vnp_Command"] = "pay"
	m["vnp_TmnCode"] = pay.TmnCode
	m["vnp_Amount"] = pay.Amount
	m["vnp_CreateDate"] = currentDate.Format("20060102150405")
	m["vnp_CurrCode"] = "VND"
	m["vnp_IpAddr"] = pay.IpAddr
	m["vnp_Locale"] = "vn"
	m["vnp_OrderInfo"] = pay.OrderInfo
	m["vnp_OrderType"] = pay.OrderType
	m["vnp_ReturnUrl"] = pay.ReturnUrl
	m["vnp_TxnRef"] = currentDate.Format("150405")
	m["vnp_BankCode"] = pay.BankCode
	m["vnp_ExpireDate"] = currentDate.Format("20060102150405")

	req, _ := http.NewRequest("GET", pay.VnpUrl, nil)
	q := req.URL.Query()

	for key := range m {
		q.Add(key, m[key])
	}
	keys := []string{}
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	qData := ""
	for _, key := range keys {
		qData += key + "=" + strings.Join(q[key], "") + "&"
	}

	signData := pay.SecretKey + qData

	lengSignData := len(signData)
	signData = signData[:lengSignData-1]

	hash := sha256.New()
	hash.Write([]byte(signData))
	hashed := hex.EncodeToString(hash.Sum(nil))

	q.Add("vnp_SecureHashType", "SHA256")
	q.Add("vnp_SecureHash", hashed)

	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}
