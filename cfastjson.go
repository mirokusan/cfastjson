/*
 * @URL: github.com/mirokusan/cfastjson
 * @Date: 2021-01-04 01:06:14
 * @Last Modified time: 2021-01-04 23:48:53
 */

package cfastjson

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

// FormatString -
func (c *Cfastjson) FormatString() string {
	s := c.Value.String()
	var result interface{}
	json.Unmarshal(String2Bytes(s), &result)
	buf, _ := json.MarshalIndent(result, "", " ")
	return Bytes2String(buf)
}

// Set -
func (c *Cfastjson) Set(key string, val interface{}) *Cfastjson {
	fjs := c.Value
	if fjs.Type() != fastjson.TypeObject {
		panic("json type is not object")
	}

	jsVal, e := parse(val)
	if e != nil {
		panic(e)
	}

	fjs.Set(key, jsVal)
	return c
}

// Append -
func (c *Cfastjson) Append(val interface{}) *Cfastjson {
	fjs := c.Value
	if fjs.Type() != fastjson.TypeArray {
		panic("json type is not array")
	}

	jsVal, e := parse(val)
	if e != nil {
		panic(e)
	}

	fjs.SetArrayItem(len(fjs.GetArray()), jsVal)
	return c
}
