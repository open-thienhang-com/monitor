package zalopay

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zpmep/hmacutil"
)

type zaloPayload struct {
	Amount  string `json:"amount"`
	OrderID string `json:"orderId"`
}

func CreateZalo(orderId string, amount string) IPayment {
	return &zaloPayload{
		Amount:  amount,
		OrderID: orderId,
	}
}

func (payload *zaloPayload) CaptureWallet() (interface{}, error) {
	type object map[string]interface{}

	rand.Seed(time.Now().UnixNano())
	transID := rand.Intn(1000000) // Generate random trans id
	embedData, _ := json.Marshal(object{})
	items, _ := json.Marshal([]object{})
	// request data
	params := make(url.Values)
	params.Add("app_id", "2553")
	params.Add("amount", "100000") //payload.Amount)
	params.Add("app_user", "user123")
	params.Add("embed_data", string(embedData))
	params.Add("item", string(items))
	params.Add("description", "xxxx.com - Payment for the order #"+strconv.Itoa(transID))
	params.Add("bank_code", "zalopayapp")

	now := time.Now()
	params.Add("app_time", strconv.FormatInt(now.UnixNano()/int64(time.Millisecond), 10)) // miliseconds

	params.Add("app_trans_id", fmt.Sprintf("%02d%02d%02d_%v", now.Year()%100, int(now.Month()), now.Day(), transID)) // translation missing: vi.docs.shared.sample_code.comments.app_trans_id

	// appid|app_trans_id|appuser|amount|apptime|embeddata|item
	data := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v", params.Get("app_id"), params.Get("app_trans_id"), params.Get("app_user"),
		params.Get("amount"), params.Get("app_time"), params.Get("embed_data"), params.Get("item"))
	params.Add("mac", hmacutil.HexStringEncode(hmacutil.SHA256, "PcY4iZIKFCIdgZvA6ueMcMHHUbRLYjPL", data))
	log.Error(params)
	resp, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/create", params)
	log.Error(resp)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("zalopay is not available")
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["order_url"], nil
}

func (payload *zaloPayload) GetInvoiceID() string {
	return payload.OrderID
}

func (payload *zaloPayload) GetAmount() string {
	return payload.Amount
}
