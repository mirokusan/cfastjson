/*
 * @URL: github.com/mirokusan/cfastjson
 * @Date: 2021-01-04 01:06:14
 * @Last Modified time: 2021-01-04 23:48:53
 */

package cfastjson

import "encoding/json"

// FormatString -
func (c *Cfastjson) FormatString() string {
	s := c.Value.String()
	var result interface{}
	json.Unmarshal(String2Bytes(s), &result)
	buf, _ := json.MarshalIndent(result, "", " ")
	return Bytes2String(buf)
}
