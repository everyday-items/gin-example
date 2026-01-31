package util

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/everyday-items/gin-example/library/logging"
)

// JSONDecode 解码json数据
func JSONDecode(r io.Reader, obj interface{}) error {
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return nil
}

// MapToJson map转json
func MapToJson(m map[string]interface{}) string {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		logging.Error(fmt.Sprintf("{\"err\":%s, \"func\":\"MapToJson\"}", err))
		return ""
	}
	return string(jsonByte)
}

func MapToJsonStr(m map[string]string) string {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		logging.Error(fmt.Sprintf("{\"err\":%s, \"func\":\"MapToJsonStr\"}", err))
		return ""
	}
	return string(jsonByte)
}

// DateToTimestamp Y-M-D H:i:s 日期转时间戳
func DateToTimestamp(timeStr string) int64 {
	var loc, _ = time.LoadLocation("Asia/Shanghai")
	timeTmeplate := "2006-01-02 15:04:05"
	tim, _ := time.ParseInLocation(timeTmeplate, timeStr, loc)
	return tim.UnixMilli()
}

func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GetInterfaceToString interface 转 string
func GetInterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case time.Time:
		t, _ := value.(time.Time)
		key = t.String()
		key = strings.Replace(key, " +0800 CST", "", 1)
		key = strings.Replace(key, " +0000 UTC", "", 1)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func MapToBytes(m map[string]interface{}) (b []byte, err error) {
	b, err = json.Marshal(m)
	if err != nil {
		return b, err
	}
	return
}

// IntToBytes 整型转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// BytesToInt 字节转换成整型
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

// MergeMap 合并多个map
func MergeMap(mObj ...map[string]string) map[string]string {
	newObj := make(map[string]string)
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

func InArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func StructToJson(v interface{}) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
