/*
 * @URL: github.com/mirokusan/cfastjson
 * @Date: 2021-11-08 22:14:10
 * @Last Modified time: 2021-11-08 22:14:10
 */

package cfastjson

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

// Parse - 解析生成cfjsV
func Parse(v interface{}) (*fastjson.Value, error) {
	js, e := parse(v)
	if e == nil {
		return js, nil
	}
	return fastjson.Parse(fmt.Sprintf("%v", v))
}

// Parse2 - 解析生成cfjsV，失败后，尝试使用标准json库进行解析
func Parse2(v interface{}) (*fastjson.Value, error) {
	r, e1 := parse(v)
	if e1 == nil {
		return r, nil
	}

	b, e := json.Marshal(v)
	if e != nil {
		return nil, e1
	}

	return fastjson.ParseBytes(b)
}

// MustParse - 解析生成cfjsV，失败后panic
func MustParse(v interface{}) *fastjson.Value {
	r, e := parse(v)
	if e != nil {
		panic(e)
	}
	return r
}

// MustParse2 - 解析生成cfjsV，失败后，尝试使用标准json库进行解析，失败后panic
func MustParse2(v interface{}) *fastjson.Value {
	r, e := Parse2(v)
	if e != nil {
		panic(e)
	}
	return r
}

// FormatString - 格式化字符串输出
func FormatString(v *fastjson.Value) string {
	if v == nil {
		return "null"
	}

	s := v.String()

	var itf interface{}
	e := json.Unmarshal([]byte(s), &itf)
	if e != nil {
		return "null"
	}
	b, e := json.MarshalIndent(itf, "", " ")
	if e != nil {
		return "null"
	}
	return string(b)
}

// NewCfastJSON -
func NewCfastJSON(fjs *fastjson.Value) *Cfastjson {
	return &Cfastjson{
		Value: *fjs,
	}
}

// NewCfastJSONObj -
func NewCfastJSONObj() *Cfastjson {
	return &Cfastjson{
		Value: *fastjson.MustParse("{}"),
	}
}

// NewCfastJSONArray -
func NewCfastJSONArray() *Cfastjson {
	return &Cfastjson{
		Value: *fastjson.MustParse("[]"),
	}
}
