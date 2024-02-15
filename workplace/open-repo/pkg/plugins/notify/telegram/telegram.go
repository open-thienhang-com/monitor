package telegram

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SendNotification(content string) {
	data := url.Values{}
	data.Set("text", content)
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot2082980755:AAHkEB4RsO2x-6YBVjBbOZtQXMe4_AJdAMg/sendMessage?chat_id=-1001749804629", strings.NewReader(data.Encode()))
	// Header - API get user information
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Print(err.Error())
	}
	//
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
}

func GetIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			v := ip.String()
			return v
		}
	}
	return ""
}

func SendNotificationStarted() {
	SendNotification("🆘🆘🆘 Server started at: " + time.Now().String() + "\n IP: " + GetIP())
}

func SendHealthCheck(cpu, ram string) {
	SendNotification("Server : " + GetIP() +
		"\n 🎄 HEAP: " + ram +
		"\n 🎄 STACK: " + cpu)
}

func SendCreateInvoice(invoiceID, userID string) {
	SendNotification(" 📤 📤 📤  ĐƠN HÀNG ID: #" + invoiceID + " \n ⚠️ TRẠNG THÁI: ĐANG CHỜ XỬ LÍ \n ⚠️ TÀI KHOẢN: " + userID + "\n ⚠️ LÚC: " + time.Now().Local().Format(time.RFC1123Z) + "\n ⚠️ PHƯƠNG THỨC THANH TOÁN: Momo")
}

func SendDepositInvoice(invoiceID, userID, amount string) {
	SendNotification(" 📩 📩 📩  HOÁ ĐƠN ĐÃ  ID: #" + invoiceID + " \n ⚠️ TRẠNG THÁI: ĐÃ THANH TOÁN THÀNH CÔNG \n ✅ TÀI KHOẢN: " + userID + "\n ✅ LÚC: " + time.Now().Local().Format(time.RFC1123Z) + "\n ✅ PHƯƠNG THỨC THANH TOÁN: Momo. Số tiền cần đối soát:" + amount)
}

func SendErrorInvoice(invoiceID, userID string) {
	SendNotification(" 📩 📩 📩  HOÁ ĐƠN ĐÃ  ID: #" + invoiceID + " \n ⚠️ TRẠNG THÁI: ĐÃ THANH TOÁN KHÔNG THÀNH CÔNG \n ⛔ TÀI KHOẢN: " + userID + "\n ⛔ LÚC: " + time.Now().Local().Format(time.RFC1123Z) + "\n ⛔ PHƯƠNG THỨC THANH TOÁN: Momo")
}
