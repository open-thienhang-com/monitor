package momo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
)

// MomoIPN
type MomoIPN struct {
	PartnerCode  string `json:"partnerCode"  example:"MOMOIOLD20190129"`
	InvoiceId    string `json:"orderId"  example:"01234567890123451633504872421"`
	RequestId    string `json:"requestId" example:"01234567890123451633504872421"`
	Amount       int    `json:"amount"  example:"1000"`
	OrderInfo    string `json:"orderInfo"  example:"Test Thue 1234556"`
	OrderType    string `json:"orderType"  example:"momo_wallet"`
	TransId      int    `json:"transId"  example:"2588659987"`
	ResultCode   int    `json:"resultCode"  example:"0"`
	Message      string `json:"message"  example:"Giao dịch thành công."`
	PayType      string `json:"payType"  example:"qr"`
	ResponseTime int    `json:"responseTime"  example:"1633504902954"`
	ExtraData    string `json:"extraData"  example:"eyJyZXN1bHRfbmFtZXNwYWNlIjoidW1hcmtldCIsImVycm9yIjoiIiwic3RhdGUiOjZ9"`
	Signature    string `json:"signature"  example:"90482b3881bdf863d5f61ace078921bbc6dbb58b2fded35261c71c9af3b1ce4f"`
}

func CheckSignatureMoMo(r *http.Request) (amount int, orderID string, err error) {
	var mm MomoIPN
	if err := json.NewDecoder(r.Body).Decode(&mm); err != nil {
		return 0, "", err
	}
	var rawSignature bytes.Buffer
	rawSignature.WriteString("amount=")
	rawSignature.WriteString(fmt.Sprint(mm.Amount))
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(mm.ExtraData)
	rawSignature.WriteString("&message=")
	rawSignature.WriteString(mm.Message)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(mm.InvoiceId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(mm.OrderInfo)
	rawSignature.WriteString("&orderType=")
	rawSignature.WriteString(mm.OrderType)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(mm.PartnerCode)
	rawSignature.WriteString("&payType=")
	rawSignature.WriteString(mm.PayType)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(mm.RequestId)
	rawSignature.WriteString("&responseTime=")
	rawSignature.WriteString(fmt.Sprint(mm.ResponseTime))
	rawSignature.WriteString("&resultCode=")
	rawSignature.WriteString(fmt.Sprint(mm.ResultCode))
	rawSignature.WriteString("&transId=")
	rawSignature.WriteString(fmt.Sprint(mm.TransId))

	// Create a new HMAC by defining the hash type and the key (as byte array)
	var secretKey = "fXrgYaYR55WpeSHjZpIMp18gtIW4dDcA"
	hmac := hmac.New(sha256.New, []byte(secretKey))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())

	// Get result and encode as hexadecimal string
	// signature := hex.EncodeToString(hmac.Sum(nil))
	// fmt.Println(signature)
	return mm.Amount, mm.InvoiceId, nil
}
