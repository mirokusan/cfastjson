/*
 * @URL: github.com/mirokusan/cfastjson
 * @Date: 2021-01-04 01:06:07
 * @Last Modified time: 2021-01-04 23:48:42
 */

package cfastjson

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/valyala/fastjson"
)

// Cfastjson -
type Cfastjson struct {
	*fastjson.Value
}

func trim(s string) string {
	return strings.TrimLeft(strings.TrimRight(s, " "), " ")
}

func convertChar(s string) string {
	pBuf := &bytes.Buffer{}

	for _, v := range s {
		switch v {
		case '"', '\\':
			pBuf.WriteByte('\\')
		}
		pBuf.WriteRune(v)
	}

	return pBuf.String()
}

/// getFormatStr - 获取fastjson格式字符串
func strFormat(s string) string {
	s = trim(s)

	if len(s) > 0 && (s[0] == '[' && s[len(s)-1] == ']' || s[0] == '{' && s[len(s)-1] == '}') {
		return s
	}

	s = convertChar(s)
	return fmt.Sprintf(`"%s"`, s)
}

/// parse - 解析json
func parse(v interface{}) (*fastjson.Value, error) {
	if v == nil {
		return &fastjson.Value{}, nil
	}

	s := ""
	switch v2 := v.(type) {
	case *string:
		s = strFormat(*v2)

	case string:
		s = strFormat(v2)

	case nil:
		return &fastjson.Value{}, nil

	case []string, []interface{}:
		return parseSlice(v2)

	default:
		s = fmt.Sprintf("%v", v2)
	}
	return fastjson.Parse(s)
}

///
func parseSlice(item interface{}) (*fastjson.Value, error) {
	return nil, nil
}
