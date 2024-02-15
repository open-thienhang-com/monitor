package zalopay

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/zpmep/hmacutil"
)

func CheckSignatureZalo(r *http.Request) (amount int, orderID string, err error) {
	defer r.Body.Close()
	var cbdata map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&cbdata)

	requestMac := cbdata["mac"].(string)
	dataStr := cbdata["data"].(string)
	mac := hmacutil.HexStringEncode(hmacutil.SHA256, "eG4r0GcoNtRGbO8", dataStr)
	log.Println("mac =", mac)

	// kiểm tra callback hợp lệ (đến từ ZaloPay server)
	if mac != requestMac {
		return 0, "", errors.New("mac not equal")
	}

	// merchant cập nhật trạng thái cho đơn hàng
	var dataJSON map[string]interface{}
	json.Unmarshal([]byte(dataStr), &dataJSON)
	log.Println("update order's status = success where app_trans_id =", dataJSON["app_trans_id"])

	return 0, "mm.InvoiceId", nil
}
