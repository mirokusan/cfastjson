/*
 * @URL: github.com/mirokusan/cfastjson
 * @Date: 2021-01-04 23:46:47
 * @Last Modified time: 2021-01-04 23:49:43
 */

package cfastjson_test

import (
	"fmt"
	"mirokusan/cfastjson"
	"testing"

	"github.com/valyala/fastjson"
)

func TestA_1(t *testing.T) {
	str := `{}`
	fmt.Println(str)

	fjs, e := cfastjson.Parse(str)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(fjs.String())
	}

	str = `{"key": 1, "key2": 2, "key3": "val"}   `
	fjs, e = cfastjson.Parse(str)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(fjs.String())
		fmt.Println(fjs.Type())
	}
}

func TestA_2(t *testing.T) {
	str := `" abcdeftggg "`
	fjs, e := fastjson.Parse(str)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(fjs.String())
	}

	str = "[1,2,3,4]"
	fjs, e = fastjson.Parse(fmt.Sprintf(`"%s"`, str))
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(fjs.String())
		fmt.Println(fjs.Type())
	}
}

func TestA_3(t *testing.T) {
	str := `"  abcdeftggg "`
	fjs, e := cfastjson.Parse(str)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(fjs.String())
	fmt.Println(fjs.Type())

	str = "\"  abcdeftggg \""
	fjs, e = cfastjson.Parse(str)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(fjs.String())
	fmt.Println(fjs.Type())

	str = "abcde\""
	fjs, e = cfastjson.Parse(str)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(fjs.String())
	fmt.Println(fjs.Type())

	str = "D:\\test\\test2"
	fmt.Println(str)
	fjs, e = cfastjson.Parse(str)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(fjs.String())

	fmt.Printf("%s", fjs.GetStringBytes())
}

func Benchmark_cfastjson(b *testing.B) {

}
