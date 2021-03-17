package stringex

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGUID(t *testing.T) {
	for i := 0; i < 505; i++ {
		id := GUID()
		if id == "" {
			t.Fatal(id)
		}
	}
}

func TestJoin(t *testing.T) {
	arr := []string{"a", "b", "c"}
	str := Join("-", arr...)
	if str != strings.Join(arr, "-") {
		t.Fatal(str)
	}
}

func TestJoin_空切片(t *testing.T) {
	var arr []string
	str := Join("-", arr...)
	if str != "" {
		t.Fatal(str)
	}
}

func TestJoin_无元素(t *testing.T) {
	str := Join("-")
	if str != "" {
		t.Fatal(str)
	}
}

func TestSilce(t *testing.T) {
	rv := reflect.MakeSlice(
		reflect.SliceOf(
			reflect.TypeOf(0),
		),
		0,
		0,
	)
	for i := 0; i < 10; i++ {
		rv = reflect.Append(
			rv,
			reflect.ValueOf(i),
		)
	}
	arr, ok := rv.Interface().([]int)
	if !(ok && len(arr) == 10) {
		t.Fatal(ok, arr)
	}
}

var (
	one = make(map[int]map[int]string)
	two = make(map[int]map[int]string)
)

func opMap(dic map[int]map[int]string) {
	dic[1] = make(map[int]string)
	dic[1][11] = "a"
}

func Test_dic(t *testing.T) {
	opMap(one)
	fmt.Println("res", one)
	t.Error("err")
}
