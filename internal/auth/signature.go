package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Signature(appId, appKey string, dataModel any) string {

	baseStr := DataModelGenSignature(dataModel, fmt.Sprintf("app_id=%s", appId))
	return HS256(baseStr, appKey)
}

func GenRequestId(dataModel any, appId string) string {
	return DataModelGenSignature(dataModel, fmt.Sprintf("app_id=%s", appId))
}
func DataModelGenSignature(dataModel any, extraData ...string) string {
	var d = encodeToStrings(dataModel, "signature")
	d = append(d, extraData...)
	//sort strings slice
	sort.Strings(d)
	return m5(strings.Join(d, "&&"))
}
func encodeToStrings(dataModel any, excludedField ...string) []string {
	var t = reflect.TypeOf(dataModel)
	var v = reflect.ValueOf(dataModel)
	var elements []string

	if t.Kind() != reflect.Struct && t.Kind() != reflect.Ptr {
		panic(errors.New("the input structure type must be a structure or a structure pointer"))
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

For1:
	for i := 0; i < t.NumField(); i++ {

		var tagName = getTagName(t, i)
		if !t.Field(i).IsExported() {
			continue
		}
		if len(excludedField) > 0 {
			for _, field := range excludedField {
				if tagName == field {
					continue For1
				}
			}
		}

		ele := fmt.Sprintf("%s=%v",
			tagName,
			analyzeField(v.Field(i), t.Field(i).Type),
		)
		elements = append(elements, ele)
	}
	return elements
}

func analyzeField(v reflect.Value, t reflect.Type) interface{} {

	switch t.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int64, reflect.Int16, reflect.Int32:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Bool:
		return v.Bool()
	case reflect.Interface:
		if v.CanInterface() {
			return v.Interface()
		}
		return nil
	case reflect.Slice:
		return handleSlice(v, v.Type())
	case reflect.Struct:
		if !v.IsNil() {
			return handleStruct(v, v.Type())
		}
	case reflect.Ptr:
		if !v.IsNil() {
			return handleStruct(v.Elem(), v.Type())
		}
	default:
		return nil
	}
	return nil
}

func handleSlice(v reflect.Value, t reflect.Type) interface{} {

	var length = v.Len()
	if length <= 0 {
		return nil
	}

	var data = []string{}
	for i := 0; i < length; i++ {

		var f = analyzeField(v.Index(i), v.Index(i).Type())
		//
		if f != nil {
			data = append(data, f.(string))
		}
	}

	if len(data) > 0 {
		//sort
		sort.Strings(data)
		//join
		return base64Encode(strings.Join(data, ","))
	}
	return nil
}

func handleStruct(v reflect.Value, t reflect.Type) interface{} {

	var data []string

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		var f = analyzeField(v.Field(i), t.Field(i).Type)
		if f == nil {
			continue
		} else {
			data = append(data, fmt.Sprintf("%s=%v", getTagName(t, i), f))
		}
	}
	//sort the string
	sort.Strings(data)
	//join the string
	return base64Encode(strings.Join(data, "&&"))
}

func getTagName(f reflect.Type, i int) string {
	var tagName = f.Field(i).Tag.Get("json")

	if len(tagName) <= 0 {
		tagName = strings.ToLower(f.Field(i).Name)
	}
	return tagName
}
func base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}
func m5(src string) string {
	var m = md5.New()
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

func HS256(md5String string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(md5String))
	return fmt.Sprintf("%x", h.Sum(nil))
}
