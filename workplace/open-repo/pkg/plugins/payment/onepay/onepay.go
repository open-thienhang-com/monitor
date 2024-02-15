package onepay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
)

type onePayPayload struct {
	AgainLink   string `json:"againLink"`
	Title       string `json:"title"`
	AccessCode  string `json:"accessCode"`
	Amount      string `json:"amount"`
	Command     string `json:"command"`
	Currency    string `json:"currency"`
	Locale      string `json:"locale"`
	MerchTxnRef string `json:"merchTxnRef"`
	Merchant    string `json:"merchant"`
	OrderInfo   string `json:"orderInfo"`
	TicketNo    string `json:"ticketNo"`
	Version     string `json:"version"`
	ReturnURL   string `json:"returnURL"`
	Signature   string `json:"signature"  example: "90482b3881bdf863d5f61ace078921bbc6dbb58b2fded35261c71c9af3b1ce4f"`
}

func CreateOnePay(invoiceID string, am string, ip string) IPayment {
	var title = "xxxx.com"
	var againLink = "payment.xxxx.com"
	var accessCode = "C0905B01"
	// var accessCode = "6BEB2566" //test OP
	// var accessCode = "6BEB2546" //test OP
	var command = "pay"
	var currency = "VND"
	var locale = "vn"
	var merchTxnRef = invoiceID
	// var merchant = "TESTONEPAY25" //test OP
	// var merchant = "TESTONEPAY" //test OP
	var merchant = "OP_SANDEX"
	var orderInfo = "xxxx.com"
	var ticketNo = ip
	var version = "2"
	var returnURL = "https://payment.xxxx.com/api/v1/ipn/onepay"
	// var returnURL = "http://localhost:8001/api/v1/ipn/onepay" //test local
	// var returnURL = "xxxx://onepay/"
	var amount = am

	// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120D6") //test OP
	// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120DB") //test OP
	secretKey, _ := hex.DecodeString("B157D0AB54E32DF09156BF5E4D7E9988")
	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("vpc_AccessCode=")
	rawSignature.WriteString(accessCode)
	rawSignature.WriteString("&vpc_Amount=")
	rawSignature.WriteString(am)
	rawSignature.WriteString("&vpc_Command=")
	rawSignature.WriteString(command)
	rawSignature.WriteString("&vpc_Currency=")
	rawSignature.WriteString(currency)
	rawSignature.WriteString("&vpc_Locale=")
	rawSignature.WriteString(locale)
	rawSignature.WriteString("&vpc_MerchTxnRef=")
	rawSignature.WriteString(merchTxnRef)
	rawSignature.WriteString("&vpc_Merchant=")
	rawSignature.WriteString(merchant)
	rawSignature.WriteString("&vpc_OrderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&vpc_ReturnURL=")
	rawSignature.WriteString(returnURL)
	rawSignature.WriteString("&vpc_TicketNo=")
	rawSignature.WriteString(ticketNo)
	rawSignature.WriteString("&vpc_Version=")
	rawSignature.WriteString(version)

	hmac := hmac.New(sha256.New, []byte(secretKey))
	hmac.Write(rawSignature.Bytes())
	signature := hex.EncodeToString(hmac.Sum(nil))

	return &onePayPayload{
		AgainLink:   againLink,
		Title:       title,
		AccessCode:  accessCode,
		Amount:      amount,
		Command:     command,
		Currency:    currency,
		Locale:      locale,
		MerchTxnRef: merchTxnRef,
		Merchant:    merchant,
		OrderInfo:   orderInfo,
		TicketNo:    ticketNo,
		Version:     version,
		ReturnURL:   returnURL,
		Signature:   signature,
	}
}
func (payload *onePayPayload) GetAmount() string {
	return payload.Amount
}

func (payload *onePayPayload) CaptureWallet() (interface{}, error) {
	log.Info("CaptureWallet")
	// var endpoint = os.Getenv("ONEPAY_URL")
	var endpoint = "https://onepay.vn/paygate/vpcpay.op" //Hòa theme
	// var endpoint = "https://mtf.onepay.vn/paygate/vpcpay.op" //test OP
	log.Error(os.Getenv("ONEPAY_URL"))
	req, err := http.NewRequest("GET", endpoint, nil)
	log.Error(req)
	if err != nil {
		log.Error(err)
		return "", err
	}
	// defer req.Body.Close()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	parm := req.URL.Query()
	parm.Add("vpc_AccessCode", payload.AccessCode)
	parm.Add("vpc_Amount", payload.Amount)
	parm.Add("vpc_Command", payload.Command)
	parm.Add("vpc_Currency", payload.Currency)
	parm.Add("vpc_Locale", payload.Locale)
	parm.Add("vpc_MerchTxnRef", payload.MerchTxnRef)
	parm.Add("vpc_Merchant", payload.Merchant)
	parm.Add("vpc_OrderInfo", payload.OrderInfo)
	parm.Add("vpc_ReturnURL", payload.ReturnURL)
	parm.Add("vpc_TicketNo", payload.TicketNo)
	parm.Add("vpc_Version", payload.Version)
	parm.Add("vpc_SecureHash", payload.Signature)
	parm.Add("AgainLink", payload.AgainLink)
	parm.Add("Title", payload.Title)
	req.URL.RawQuery = parm.Encode()
	log.Info(req.URL.String())
	return req.URL.String(), nil
}

func (payload *onePayPayload) GetInvoiceID() string {
	return payload.MerchTxnRef
}
