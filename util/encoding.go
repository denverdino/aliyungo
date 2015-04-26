package util

import (
	"encoding/json"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const ISO8601_TIMESTAMP_FORMAT = "2014-05-26T12:00:00Z"

func ConvertToQueryValues(i interface{}) url.Values {
	values := url.Values{}
	SetQueryValues(i, &values)
	return values
}

func SetQueryValues(i interface{}, values *url.Values) {
	elem := reflect.ValueOf(i)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	elemType := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		// TODO Use Tag for validation
		// tag := typ.Field(i).Tag.Get("tagname")
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}
		var value string
		switch field.Interface().(type) {
		case int, int8, int16, int32, int64:
			i := field.Int()
			if i != 0 {
				value = strconv.FormatInt(i, 10)
			}
		case uint, uint8, uint16, uint32, uint64:
			i := field.Uint()
			if i != 0 {
				value = strconv.FormatUint(i, 10)
			}
		case float32:
			value = strconv.FormatFloat(field.Float(), 'f', 4, 32)
		case float64:
			value = strconv.FormatFloat(field.Float(), 'f', 4, 64)
		case []byte:
			if !field.IsNil() {
				value = string(field.Bytes())
			}
		case bool:
			value = strconv.FormatBool(field.Bool())
		case string:
			value = field.String()
		case []string:
			if !field.IsNil() && field.Len() > 0 {
				l := field.Len()
				strArray := make([]string, l)
				for i := 0; i < l; i++ {
					strArray[i] = field.Index(i).String()
				}
				bytes, err := json.Marshal(strArray)
				if err == nil {
					value = string(bytes)
				} else {
					log.Printf("Failed to convert JSON: %v", err)
				}
			}
		case time.Time:
			t := field.Interface().(time.Time)
			value = GetISO8601TimeStamp(t)
		default:
			ifc := field.Interface()
			if ifc != nil {
				SetQueryValues(ifc, values)
				continue
			}
		}
		if value != "" {
			values.Set(elemType.Field(i).Name, value)
		}
	}
}
