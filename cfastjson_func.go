/*
 * @URL: github.com/mirokusan/cfastjson
 * @Date: 2021-11-08 22:20:09
 * @Last Modified time: 2021-11-08 22:20:09
 */

package cfastjson

import (
	"reflect"
	"unsafe"

	"github.com/valyala/fastjson"
)

// String2Bytes -
func String2Bytes(s string) (b []byte) {
	hb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hs := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hb.Cap = hs.Len
	hb.Len = hs.Len
	hb.Data = hs.Data
	return
}

// Bytes2String -
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Member -
func Member(fjs *Cfastjson) (members []string) {
	if fjs.Type() != fastjson.TypeObject {
		return
	}

	fjs.GetObject().Visit(func(key []byte, v *fastjson.Value) {
		members = append(members, Bytes2String(key))
	})
	return
}
