package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

var defaultUserAgent = fmt.Sprintf("AlibabaCloud (%s; %s) Golang/%s Core/%s TeaDSL/1", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), "0.01")

type RuntimeOptions struct {
	Autoretry      *bool   `json:"autoretry" xml:"autoretry"`
	IgnoreSSL      *bool   `json:"ignoreSSL" xml:"ignoreSSL"`
	MaxAttempts    *int    `json:"maxAttempts" xml:"maxAttempts"`
	BackoffPolicy  *string `json:"backoffPolicy" xml:"backoffPolicy"`
	BackoffPeriod  *int    `json:"backoffPeriod" xml:"backoffPeriod"`
	ReadTimeout    *int    `json:"readTimeout" xml:"readTimeout"`
	ConnectTimeout *int    `json:"connectTimeout" xml:"connectTimeout"`
	LocalAddr      *string `json:"localAddr" xml:"localAddr"`
	HttpProxy      *string `json:"httpProxy" xml:"httpProxy"`
	HttpsProxy     *string `json:"httpsProxy" xml:"httpsProxy"`
	NoProxy        *string `json:"noProxy" xml:"noProxy"`
	MaxIdleConns   *int    `json:"maxIdleConns" xml:"maxIdleConns"`
	Socks5Proxy    *string `json:"socks5Proxy" xml:"socks5Proxy"`
	Socks5NetWork  *string `json:"socks5NetWork" xml:"socks5NetWork"`
}

func (s RuntimeOptions) String() string {
	return tea.Prettify(s)
}

func (s RuntimeOptions) GoString() string {
	return s.String()
}

func (s *RuntimeOptions) SetAutoretry(v bool) *RuntimeOptions {
	s.Autoretry = &v
	return s
}

func (s *RuntimeOptions) SetIgnoreSSL(v bool) *RuntimeOptions {
	s.IgnoreSSL = &v
	return s
}

func (s *RuntimeOptions) SetMaxAttempts(v int) *RuntimeOptions {
	s.MaxAttempts = &v
	return s
}

func (s *RuntimeOptions) SetBackoffPolicy(v string) *RuntimeOptions {
	s.BackoffPolicy = &v
	return s
}

func (s *RuntimeOptions) SetBackoffPeriod(v int) *RuntimeOptions {
	s.BackoffPeriod = &v
	return s
}

func (s *RuntimeOptions) SetReadTimeout(v int) *RuntimeOptions {
	s.ReadTimeout = &v
	return s
}

func (s *RuntimeOptions) SetConnectTimeout(v int) *RuntimeOptions {
	s.ConnectTimeout = &v
	return s
}

func (s *RuntimeOptions) SetHttpProxy(v string) *RuntimeOptions {
	s.HttpProxy = &v
	return s
}

func (s *RuntimeOptions) SetHttpsProxy(v string) *RuntimeOptions {
	s.HttpsProxy = &v
	return s
}

func (s *RuntimeOptions) SetNoProxy(v string) *RuntimeOptions {
	s.NoProxy = &v
	return s
}

func (s *RuntimeOptions) SetMaxIdleConns(v int) *RuntimeOptions {
	s.MaxIdleConns = &v
	return s
}

func (s *RuntimeOptions) SetLocalAddr(v string) *RuntimeOptions {
	s.LocalAddr = &v
	return s
}

func (s *RuntimeOptions) SetSocks5Proxy(v string) *RuntimeOptions {
	s.Socks5Proxy = &v
	return s
}

func (s *RuntimeOptions) SetSocks5NetWork(v string) *RuntimeOptions {
	s.Socks5NetWork = &v
	return s
}

func ReadAsString(body io.Reader) (*string, error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return tea.String(""), err
	}
	return tea.String(string(byt)), nil
}

func StringifyMapValue(a map[string]interface{}) map[string]*string {
	res := make(map[string]*string)
	for key, value := range a {
		if value != nil {
			switch value.(type) {
			case string:
				res[key] = tea.String(value.(string))
			default:
				byt, _ := json.Marshal(value)
				res[key] = tea.String(string(byt))
			}
		}
	}
	return res
}

func AnyifyMapValue(a map[string]*string) map[string]interface{} {
	res := make(map[string]interface{})
	for key, value := range a {
		res[key] = tea.StringValue(value)
	}
	return res
}

func ReadAsBytes(body io.Reader) ([]byte, error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

func DefaultString(reaStr, defaultStr *string) *string {
	if reaStr == nil {
		return defaultStr
	}
	return reaStr
}

func ToJSONString(a interface{}) *string {
	byt, _ := json.Marshal(a)
	return tea.String(string(byt))
}

func DefaultNumber(reaNum, defaultNum *int) *int {
	if reaNum == nil {
		return defaultNum
	}
	return reaNum
}

func ReadAsJSON(body io.Reader) (result interface{}, err error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	if string(byt) == "" {
		return
	}
	err = json.Unmarshal(byt, &result)
	return
}

func GetNonce() *string {
	return tea.String(getUUID())
}

func Empty(val *string) *bool {
	return tea.Bool(val == nil || tea.StringValue(val) == "")
}

func ValidateModel(a interface{}) error {
	if a == nil {
		return nil
	}
	err := tea.Validate(a)
	return err
}

func EqualString(val1, val2 *string) *bool {
	return tea.Bool(tea.StringValue(val1) == tea.StringValue(val2))
}

func EqualNumber(val1, val2 *int) *bool {
	return tea.Bool(tea.IntValue(val1) == tea.IntValue(val2))
}

func IsUnset(val interface{}) *bool {
	if val == nil {
		return tea.Bool(true)
	}

	valType := reflect.TypeOf(val)
	valZero := reflect.Zero(valType)
	v := reflect.ValueOf(val)
	return tea.Bool(valZero == v)
}

func ToBytes(a *string) []byte {
	return []byte(tea.StringValue(a))
}

func AssertAsMap(a interface{}) map[string]interface{} {
	res, ok := a.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("%v is not a map[string]interface{}", a))
	}
	return res
}

func AssertAsNumber(a interface{}) *int {
	res, ok := a.(int)
	if !ok {
		panic(fmt.Sprintf("%v is not a int", a))
	}
	return tea.Int(res)
}

func AssertAsBoolean(a interface{}) *bool {
	res, ok := a.(bool)
	if !ok {
		panic(fmt.Sprintf("%v is not a bool", a))
	}
	return tea.Bool(res)
}

func AssertAsString(a interface{}) *string {
	res, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("%v is not a string", a))
	}
	return tea.String(res)
}

func ParseJSON(a *string) interface{} {
	tmp := make(map[string]interface{})
	err := json.Unmarshal([]byte(tea.StringValue(a)), &tmp)
	if err == nil {
		return tmp
	}

	if num, err := strconv.Atoi(tea.StringValue(a)); err == nil {
		return num
	}

	if ok, err := strconv.ParseBool(tea.StringValue(a)); err == nil {
		return ok
	}

	if floa64tVal, err := strconv.ParseFloat(tea.StringValue(a), 64); err == nil {
		return floa64tVal
	}
	return nil
}

func ToString(a []byte) *string {
	return tea.String(string(a))
}

func ToMap(in interface{}) map[string]interface{} {
	if in == nil {
		return nil
	}
	res := tea.ToMap(in)
	return res
}

func ToFormString(a map[string]interface{}) *string {
	if a == nil {
		return tea.String("")
	}
	res := ""
	first := true
	for k, v := range a {
		if first {
			first = false
		} else {
			res += "&"
		}
		res += url.QueryEscape(k)
		res += "="
		res += url.QueryEscape(fmt.Sprintf("%v", v))
	}
	return tea.String(res)
}

func GetDateUTCString() *string {
	return tea.String(time.Now().UTC().Format(http.TimeFormat))
}

func GetUserAgent(userAgent *string) *string {
	if userAgent != nil && tea.StringValue(userAgent) != "" {
		return tea.String(defaultUserAgent + " " + tea.StringValue(userAgent))
	}
	return tea.String(defaultUserAgent)
}

func Is2xx(code *int) *bool {
	tmp := tea.IntValue(code)
	return tea.Bool(tmp >= 200 && tmp < 300)
}

func Is3xx(code *int) *bool {
	tmp := tea.IntValue(code)
	return tea.Bool(tmp >= 300 && tmp < 400)
}

func Is4xx(code *int) *bool {
	tmp := tea.IntValue(code)
	return tea.Bool(tmp >= 400 && tmp < 500)
}

func Is5xx(code *int) *bool {
	tmp := tea.IntValue(code)
	return tea.Bool(tmp >= 500 && tmp < 600)
}
