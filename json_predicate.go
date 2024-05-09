package jsonpredicate

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Op string

const (
	Exists     Op = "exists"
	Is         Op = "is"
	IsOneOf    Op = "isOneOf"
	IsNot      Op = "isNot"
	IsNotOneOf Op = "isNotOneOf"
	And        Op = "and"
	Or         Op = "or"
	Between    Op = "between"
	Above      Op = "above"
	Below      Op = "below"
)

type Object struct {
	Op    Op       `json:"op"`
	Path  string   `json:"path"`
	Value string   `json:"value"`
	Apply []Object `json:"apply"`
}

func (r *Object) FromRaw(raw string) bool {
	err := json.Unmarshal([]byte(raw), r)
	return err == nil
}

func (r *Object) ValidateIntValues() bool {
	vals := strings.Split(r.Value, ",")

	for _, iVal := range vals {
		if _, err := strconv.Atoi(iVal); err != nil {
			return false
		}
	}

	return true
}

func (r *Object) BoolValue() bool {
	return r.Value == "1"
}

func (r *Object) IntValue() int {
	valInt, err := strconv.Atoi(r.Value)
	if err != nil {
		return 0
	}

	return valInt
}

func (r *Object) InterfaceIntValues() []interface{} {
	vals := strings.Split(r.Value, ",")
	output := make([]interface{}, 0, len(vals))

	for _, iVal := range vals {
		if iValInt, err := strconv.Atoi(iVal); err == nil {
			output = append(output, iValInt)
		}
	}

	return output
}

func (r *Object) InterfaceStringValues() []interface{} {
	vals := strings.Split(r.Value, ",")
	output := make([]interface{}, 0, len(vals))

	for _, iVal := range vals {
		output = append(output, iVal)
	}

	return output
}

func (r *Object) And() bool {
	return r.Op == "and"
}

func (r *Object) Or() bool {
	return r.Op == "or"
}
