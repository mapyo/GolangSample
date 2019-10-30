package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Human struct {
	HumanID string
	Name    string // maskしたい
}

func (h Human) Sample() string {
	return "human"
}

type SampleRequest struct {
	SampleId   string
	Human      *Human
	XXX_sample string
}

func (s SampleRequest) Sample() string {
	return "hoge"
}

// ==================

type Sampler interface {
	Sample() string
}

type MaskFilter func(structName string, structField reflect.StructField) bool

type RequestMasker struct {
	Sampler
	Filter MaskFilter
}

type TagSample struct {
	ID   string
	Name string `log:"-"`
}

func (t TagSample) GoString() string {
	return BuildMaskGoString(t, tagFilter)
}

func BuildMaskGoString(s interface{}, filter func(structName string, structField reflect.StructField) bool) string {
	var result string

	// struct名の取得
	st := reflect.Indirect(reflect.ValueOf(s))
	structName := st.Type().Name()

	result = fmt.Sprintf("%s{", structName)

	for i := 0; i < st.NumField(); i++ {
		var (
			field      = st.Type().Field(i)
			fieldValue = st.FieldByName(field.Name)
		)

		if strings.HasPrefix(field.Name, "XXX") {
			continue
		}

		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = reflect.Indirect(fieldValue)
		}

		var subResult string
		if fieldValue.Kind() == reflect.Struct {
			subResult = BuildMaskGoString(fieldValue.Interface(), filter)
		} else {
			if filter(structName, field) {
				subResult = field.Name + ": " + fieldValue.String() + ", "
			}
		}

		result = result + subResult
	}

	result = result + "}"
	return result
}

func (r RequestMasker) GoString() string {
	return BuildMaskGoString(r.Sampler, r.Filter)
}

var mask = map[string][]string{
	"Human": {"aaa", "Name"},
	"Noman": {"aaa", "nono"},
}

var listFilter = func(structName string, structField reflect.StructField, isWhite bool) bool {
	if a, ok := mask[structName]; ok {
		for _, v := range a {
			if v == structField.Name {
				return isWhite
			}
		}
	}
	return !isWhite
}

var blackListFilter MaskFilter = func(structName string, structField reflect.StructField) bool {
	return listFilter(structName, structField, false)
}

var whiteListFilter MaskFilter = func(structName string, structField reflect.StructField) bool {
	return listFilter(structName, structField, true)
}

var tagFilter MaskFilter = func(structName string, structField reflect.StructField) bool {
	log := structField.Tag.Get("log")
	if log == "-" {
		return false
	}
	return true
}

func main() {
	human := Human{"replaceValue", "nameValue"}
	sampleRequest := SampleRequest{SampleId: "10", Human: &human, XXX_sample: "hoge"}

	rs := RequestMasker{sampleRequest, whiteListFilter}

	fmt.Printf("%#v\n", sampleRequest)
	fmt.Printf("%#v\n", rs)

	tagsample := TagSample{ID: "tag_id", Name: "tag_name"}
	fmt.Printf("%#v\n", tagsample)
}
