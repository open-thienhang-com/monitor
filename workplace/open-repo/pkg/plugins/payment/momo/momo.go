package momo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// define a payload, reference in https://developers.momo.vn/#cong-thanh-toan-momo-phuong-thuc-thanh-toan
type momoPayload struct {
	PartnerCode string `json:"partnerCode"`
	PartnerName string `json:"partnerName"`
	StoreId     string `json:"storeId"`
	RequestID   string `json:"requestId"`
	Amount      string `json:"amount"`
	OrderID     string `json:"orderId"`
	OrderInfo   string `json:"orderInfo"`
	RedirectUrl string `json:"redirectUrl"`
	IpnUrl      string `json:"ipnUrl"`
	ExtraData   string `json:"extraData"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
	Lang        string `json:"lang"`
}

//

func CreateMomo(id string, am string) IPayment {
	var orderId = id
	var requestId = id
	var partnerCode = "MOMOYSWU20211020"
	var accessKey = "FEbWzxcfpCYM6HBx"
	var secretKey = "fXrgYaYR55WpeSHjZpIMp18gtIW4dDcA"
	var orderInfo = "Thanh toán đơn hàng payment.xxxx.com"
	var redirectUrl = "xxxx://deposit"
	var ipnUrl = "https://payment.xxxx.com/api/v1/ipn/momo"
	var amount = am
	var requestType = "captureWallet"
	var extraData = "" //pass empty value or Encode base64 JsonString

	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(am)
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(ipnUrl)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(redirectUrl)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(requestType)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac := hmac.New(sha256.New, []byte(secretKey))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())
	fmt.Println("\nRaw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmac.Sum(nil))
	return &momoPayload{
		//AccessKey:   accessKey,
		Amount:      amount,
		ExtraData:   extraData,
		IpnUrl:      ipnUrl,
		OrderID:     orderId,
		OrderInfo:   orderInfo,
		PartnerCode: partnerCode,
		PartnerName: "xxxx.com",
		RedirectUrl: redirectUrl,
		RequestID:   requestId,
		RequestType: requestType,
		StoreId:     "N1",
		Signature:   signature,
		Lang:        "en",
	}
}

func (payload *momoPayload) CaptureWallet() (interface{}, error) {
	var jsonPayload []byte
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := http.Post(os.Getenv("MOMO_CREATE"), "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil || resp.StatusCode != 200 {
		log.Error(err)
		return nil, errors.New("momo no availabe")
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["payUrl"], nil
}

func (payload *momoPayload) GetInvoiceID() string {
	return payload.OrderID
}

func (payload *momoPayload) GetAmount() string {
	return payload.Amount
}
