package onepay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"mono.thienhang.com/pkg/utils"
)

func CheckSignatureOnePay(r *http.Request) (amount int, orderID string, err error) {
	//
	keys := []string{}
	for k := range r.URL.Query() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//
	result := r.URL.Query().Get("vpc_TxnResponseCode")
	// if result != "0" {
	// 	return 0, "", utils.ErrSignature
	// }
	if result == "1" {
		return 0, "", utils.ErrSignature1
	}
	if result == "2" {
		return 0, "", utils.ErrSignature2
	}
	if result == "3" {
		return 0, "", utils.ErrSignature3
	}
	if result == "4" {
		return 0, "", utils.ErrSignature4
	}
	if result == "5" {
		return 0, "", utils.ErrSignature5
	}
	if result == "6" {
		return 0, "", utils.ErrSignature6
	}
	if result == "7" {
		return 0, "", utils.ErrSignature7
	}
	if result == "8" {
		return 0, "", utils.ErrSignature8
	}
	if result == "9" {
		return 0, "", utils.ErrSignature9
	}
	if result == "10" {
		return 0, "", utils.ErrSignature10
	}
	if result == "11" {
		return 0, "", utils.ErrSignature11
	}
	if result == "12" {
		return 0, "", utils.ErrSignature12
	}
	if result == "13" {
		return 0, "", utils.ErrSignature13
	}
	if result == "21" {
		return 0, "", utils.ErrSignature21
	}
	if result == "22" {
		return 0, "", utils.ErrSignature22
	}
	if result == "23" {
		return 0, "", utils.ErrSignature23
	}
	if result == "24" {
		return 0, "", utils.ErrSignature24
	}
	if result == "25" {
		return 0, "", utils.ErrSignature25
	}
	if result == "253" {
		return 0, "", utils.ErrSignature253
	}
	if result == "99" {
		return 0, "", utils.ErrSignature99
	}
	if result == "B" {
		return 0, "", utils.ErrSignatureB
	}
	if result == "E" {
		return 0, "", utils.ErrSignatureE
	}
	if result == "F" {
		return 0, "", utils.ErrSignatureF
	}
	if result == "Z" {
		return 0, "", utils.ErrSignatureZ
	}
	if result == "0" {
		qData := ""
		for _, k := range keys {
			if k != "vpc_SecureHash" {
				if qData == "" {
					qData += k + "=" + strings.Join(r.URL.Query()[k], "")
					continue
				}
				qData += "&" + k + "=" + strings.Join(r.URL.Query()[k], "")
			}
		}
		//
		secretKey, _ := hex.DecodeString("B157D0AB54E32DF09156BF5E4D7E9988")
		// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120DB") //test OP
		// secretKey, _ := hex.DecodeString("6D0870CDE5F24F34F3915FB0045120D6") //test OP
		hmac := hmac.New(sha256.New, []byte(secretKey))

		hmac.Write([]byte(qData))
		//
		signature := hex.EncodeToString(hmac.Sum(nil))
		if strings.ToUpper(signature) == r.URL.Query()["vpc_SecureHash"][0] {
			am, err := strconv.Atoi(r.URL.Query()["vpc_Amount"][0])
			if err != nil {
				return -1, "", err
			}
			return am, r.URL.Query()["vpc_MerchTxnRef"][0], nil
		} else {
			// return 0, "", utils.ErrSignature
			return -1, "", utils.ErrSignatureOther //return -1 khi sai hash
		}
	} else {
		return 0, "", utils.ErrSignatureOther
	}
}
