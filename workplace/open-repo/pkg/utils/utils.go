package utils

import (
	"bytes"

	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NebulousLabs/fastrand"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	E_WELCOME = 0
	E_OTP     = 1
	E_RESET   = 2
	E_VERSION = 3
)

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
// 	w.WriteHeader(statusCode)
// 	err := json.NewEncoder(w).Encode(data)

// 	if err != nil {
// 		fmt.Fprintf(w, "%s", err.Error())
// 	}
// }

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	// if err != nil {
	// 	JSON(w, statusCode, struct {
	// 		Error string `json:"error"`
	// 	}{
	// 		Error: err.Error(),
	// 	})
	// 	return
	// }
	// JSON(w, http.StatusBadRequest, nil)
}

func GetExpirationTime(hours int) int64 {
	return time.Now().Add(time.Hour * time.Duration(hours)).Unix()
}

func GetDurationByHour(hours int) time.Duration {
	return time.Hour * time.Duration(hours)
}

func ParseTemplate(fileName string, data interface{}) (content string, err error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func ResponseWithJson(w http.ResponseWriter, status int, object interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var res Response
	if status == http.StatusOK {
		res = Response{
			false,
			"success",
			object,
		}
	} else {
		res = Response{
			true,
			"error",
			object,
		}
	}
	json.NewEncoder(w).Encode(res)
}

func SendEmail(to []string, msg string, eType int) (err error) {
	var t *template.Template
	var title string = "Thông báo"
	switch eType {
	case E_WELCOME:
		title = "Thư chào mừng thienhang.com"
		t, err = template.ParseFiles("template/welcome.html")
		if err != nil {
			logrus.Error(err)
		}
	case E_OTP:
		t, err = template.ParseFiles("template/sendotp.html")
		if err != nil {
			logrus.Error(err)
		}
	case E_RESET:
		t, _ = template.ParseFiles("./././template/resetpassword.html")
	case E_VERSION:
		title = "🐷 Thông báo bản cập nhật mới"
		t, _ = template.ParseFiles("./././template/version.html")
	default:
		title = "test"
		t, err = template.ParseFiles("template/welcome.html")
		if err != nil {
			logrus.Error(err)
		}
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", title, mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{

		Name:    to[0],
		Message: msg,
	})

	auth := smtp.PlainAuth("", "postmaster@mail.thienhang.com", "f815d1178c6b3f30aaa1bed91ea13d2b-78651cec-f6284dac", "smtp.eu.mailgun.org")
	err = smtp.SendMail("smtp.eu.mailgun.org:587", auth, "noreply@thienhang.com", to, body.Bytes())
	if err != nil {
		logrus.Error(err)
		return err
	}
	fmt.Println(auth)
	return nil
}

func GenerateOtp(number string) string {
	return "XXXXX"
}

func GetEnvVar(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func ComparePasswordAndHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return errors.New("Can't set value\n")
	}

	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	}

	return nil
}

func SetDefault2(ptr interface{}) error {
	tag := "default"
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return fmt.Errorf("Not a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get(tag); defaultVal != "-" {
			if err := setField(v.Field(i), defaultVal); err != nil {
				return err
			}

		}
	}
	return nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

func LoadJsonFile(filename string) []byte {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
	}
	return file
}

func getIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
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
			return "", err
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
			return v, nil
		}
	}
	return "", errors.New("are you connected to the network")
}

var (
	Err_ABNORMAL             = errors.New("lỗi không xác định từ hệ thống. Anh chị vui lòng liên lạc lại với CSKH xxxx.com")
	ErrAmountOut             = errors.New("amount out")
	ErrSignature             = errors.New("invalid signature")
	ErrSignature1            = errors.New("Giao dịch không thành công, Ngân hàng phát hành thẻ không cấp phép cho giao dịch hoặc thẻ chưa được kích hoạt dịch vụ thanh toán trên Internet. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignature2            = errors.New("Giao dịch không thành công, Ngân hàng phát hành thẻ từ chối cấp phép cho giao dịch. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ để biết chính xác nguyên nhân Ngân hàng từ chối.")
	ErrSignature3            = errors.New("Giao dịch không thành công, Cổng thanh toán không nhận được kết quả trả về từ ngân hàng phát hành thẻ. Vui lòng liên hệ với ngân hàng theo số điện thoại sau mặt thẻ để biết chính xác trạng thái giao dịch và thực hiện thanh toán lại")
	ErrSignature4            = errors.New("Giao dịch không thành công do thẻ hết hạn sử dụng hoặc nhập sai thông tin tháng/ năm hết hạn của thẻ. Vui lòng kiểm tra lại thông tin và thanh toán lại")
	ErrSignature5            = errors.New("Giao dịch không thành công, Thẻ không đủ hạn mức hoặc tài khoản không đủ số dư để thanh toán. Vui lòng kiểm tra lại thông tin và thanh toán lại")
	ErrSignature6            = errors.New("Giao dịch không thành công, Quá trình xử lý giao dịch phát sinh lỗi từ ngân hàng phát hành thẻ. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignature7            = errors.New("Giao dịch không thành công, Đã có lỗi phát sinh trong quá trình xử lý giao dịch. Vui lòng thực hiện thanh toán lại.")
	ErrSignature8            = errors.New("Giao dịch không thành công. Số thẻ không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature9            = errors.New("Giao dịch không thành công. Tên chủ thẻ không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature10           = errors.New("Giao dịch không thành công. Thẻ hết hạn/Thẻ bị khóa. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature11           = errors.New("Giao dịch không thành công. Thẻ chưa đăng ký sử dụng dịch vụ thanh toán trên Internet. Vui lòng liên hê ngân hàng theo số điện thoại sau mặt thẻ để được hỗ trợ.")
	ErrSignature12           = errors.New("Giao dịch không thành công. Ngày phát hành/Hết hạn không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature13           = errors.New("Giao dịch không thành công. thẻ/ tài khoản đã vượt quá hạn mức thanh toán. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature21           = errors.New("Giao dịch không thành công. Số tiền không đủ để thanh toán. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature22           = errors.New("Giao dịch không thành công. Thông tin tài khoản không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature23           = errors.New("Giao dịch không thành công. Tài khoản bị khóa. Vui lòng liên hê ngân hàng theo số điện thoại sau mặt thẻ để được hỗ trợ")
	ErrSignature24           = errors.New("Giao dịch không thành công. Thông tin thẻ không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature25           = errors.New("Giao dịch không thành công. OTP không đúng. Vui lòng kiểm tra và thực hiện thanh toán lại")
	ErrSignature253          = errors.New("Giao dịch không thành công. Quá thời gian thanh toán. Vui lòng thực hiện thanh toán lại")
	ErrSignature99           = errors.New("Giao dịch không thành công. Người sử dụng hủy giao dịch")
	ErrSignatureB            = errors.New("Giao dịch không thành công do không xác thực được 3D-Secure. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignatureE            = errors.New("Giao dịch không thành công do nhập sai CSC (Card Security Card) hoặc ngân hàng từ chối cấp phép cho giao dịch. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignatureF            = errors.New("Giao dịch không thành công do không xác thực được 3D-Secure. Vui lòng liên hệ ngân hàng theo số điện thoại sau mặt thẻ được hỗ trợ chi tiết.")
	ErrSignatureZ            = errors.New("Giao dịch không thành công do vi phạm quy định của hệ thống. Vui lòng liên hệ với OnePAY để được hỗ trợ (Hotline: 1900 633 927)")
	ErrSignatureOther        = errors.New("Giao dịch không thành công. Vui lòng liên hệ với OnePAY để được hỗ trợ (Hotline: 1900 633 927)")
	Err_STRING_METHOD        = "Anh chị vui lòng lựa chọn phương thức thanh toán hợp lệ."
	Err_STRING_TOKEN         = "Phiên làm việc của anh chị đã hết hạn hoặc tài khoản của anh chị đã bị khoá. Anh chị vui lòng đăng nhập lại hoặc liên hệ bộ phận chăm sóc khách hàng của xxxx để được hỗ trợ."
	Err_STRING_RANGE_DEPOSIT = "Định dạng số tiền thanh toán của anh chị chưa hợp lệ. Số tiền thanh toán phải lớn hơn hoặc bằng 50.000 đồng, bội số của 10.000"
	// Err_STRING_METHOD =
	// Err_STRING_METHOD =
	Err_Unauthorization = ""
	Err_Signature       = "Chữ kí chưa hợp lệ"
)

func Uuid(length int64) string {
	ele := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "v", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "Driver", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	ele, _ = Random(ele)
	uuid := ""
	var i int64
	for i = 0; i < length; i++ {
		uuid += ele[fastrand.Intn(59)]
	}
	return uuid
}

func Random(strings []string) ([]string, error) {
	for i := len(strings) - 1; i > 0; i-- {
		num := fastrand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := make([]string, 0)
	for i := 0; i < len(strings); i++ {
		str = append(str, strings[i])
	}
	return str, nil
}

func CompressedContent(h *template.HTML) {
	st := strings.Split(string(*h), "\n")
	var ss []string
	for i := 0; i < len(st); i++ {
		st[i] = strings.TrimSpace(st[i])
		if st[i] != "" {
			ss = append(ss, st[i])
		}
	}
	*h = template.HTML(strings.Join(ss, "\n"))
}

func ReplaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func WrapURL(u string) string {
	uarr := strings.Split(u, "?")
	if len(uarr) < 2 {
		return url.QueryEscape(strings.Replace(u, "/", "_", -1))
	}
	v, err := url.ParseQuery(uarr[1])
	if err != nil {
		return url.QueryEscape(strings.Replace(u, "/", "_", -1))
	}
	return url.QueryEscape(strings.Replace(uarr[0], "/", "_", -1)) + "?" +
		strings.Replace(v.Encode(), "%7B%7B.Id%7D%7D", "{{.Id}}", -1)
}

func JSON(a interface{}) string {
	if a == nil {
		return ""
	}
	b, _ := json.Marshal(a)
	return string(b)
}

func ParseBool(s string) bool {
	b1, _ := strconv.ParseBool(s)
	return b1
}

func ParseFloat32(f string) float32 {
	s, _ := strconv.ParseFloat(f, 32)
	return float32(s)
}

func SetDefault(value, condition, def string) string {
	if value == condition {
		return def
	}
	return value
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func CopyMap(m map[string]string) map[string]string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	var cm map[string]string
	err = dec.Decode(&cm)
	if err != nil {
		panic(err)
	}
	return cm
}

func CompareVersion(src, toCompare string) bool {
	if toCompare == "" {
		return false
	}

	exp, _ := regexp.Compile(`-(.*)`)
	src = exp.ReplaceAllString(src, "")
	toCompare = exp.ReplaceAllString(toCompare, "")

	srcs := strings.Split(src, "v")
	srcArr := strings.Split(srcs[1], ".")
	op := ">"
	srcs[0] = strings.TrimSpace(srcs[0])
	if InArray([]string{">=", "<=", "=", ">", "<"}, srcs[0]) {
		op = srcs[0]
	}

	toCompare = strings.Replace(toCompare, "v", "", -1)

	if op == "=" {
		return srcs[1] == toCompare
	}

	if srcs[1] == toCompare && (op == "<=" || op == ">=") {
		return true
	}

	toCompareArr := strings.Split(strings.Replace(toCompare, "v", "", -1), ".")
	for i := 0; i < len(srcArr); i++ {
		v, err := strconv.Atoi(srcArr[i])
		if err != nil {
			return false
		}
		vv, err := strconv.Atoi(toCompareArr[i])
		if err != nil {
			return false
		}
		switch op {
		case ">", ">=":
			if v < vv {
				return true
			} else if v > vv {
				return false
			} else {
				continue
			}
		case "<", "<=":
			if v > vv {
				return true
			} else if v < vv {
				return false
			} else {
				continue
			}
		}
	}

	return false
}

const (
	Byte  = 1
	KByte = Byte * 1024
	MByte = KByte * 1024
	GByte = MByte * 1024
	TByte = GByte * 1024
	PByte = TByte * 1024
	EByte = PByte * 1024
)

var bytesSizeTable = map[string]uint64{
	"b":  Byte,
	"kb": KByte,
	"mb": MByte,
	"gb": GByte,
	"tb": TByte,
	"pb": PByte,
	"eb": EByte,
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1024, sizes)
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time, m map[string]string) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff, m)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func computeTimeDiff(diff int64, m map[string]string) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 " + m["second"]
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d "+m["seconds"], diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 " + m["minute"]
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d "+m["minutes"], diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 " + m["hour"]
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d "+m["hours"], diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 " + m["day"]
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d "+m["days"], diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 " + m["week"]
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d "+m["weeks"], diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 " + m["month"]
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d "+m["months"], diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 " + m["year"]
	default:
		diffStr = fmt.Sprintf("%d "+m["years"], diff/Year)
		diff = 0
	}
	return diff, diffStr
}

// func GetIdRequets(ctx *context.Context) int {
// 	request_url := ctx.Request

// 	pathParts := strings.Split(request_url.URL.Path, "/")

// 	id_product, error_id := strconv.Atoi(pathParts[len(pathParts)-1])
// 	if error_id != nil {
// 		return -1
// 	} else {
// 		return id_product
// 	}
// }

// func GetBodyRequest(ctx *context.Context, type_object interface{}) (interface{}, error) {
// 	body_request, readErr := io.ReadAll(ctx.Request.Body)
// 	fmt.Println(string(body_request))
// 	data := json.Unmarshal(body_request, &type_object)
// 	fmt.Println(type_object)
// 	fmt.Println(data)
// 	if readErr != nil {
// 		fmt.Println("Error when read request", readErr)
// 		return nil, readErr
// 	}
// 	if data != nil {
// 		fmt.Println("Error when exxtract JSON", data)
// 		return nil, data
// 	}
// 	return type_object, nil
// }

// func GetParamRequest(ctx *context.Context, param_name string) interface{} {
// 	request_url := ctx.Request
// 	params := request_url.URL.Query()
// 	value_param := params.Get(param_name)

// 	return value_param
// }
