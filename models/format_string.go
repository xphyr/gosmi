// Code generated by "enumer -type=Format -autotrimprefix -json format.go"; DO NOT EDIT

package models

import (
	"encoding/json"
	"fmt"
)

const (
	_Format_name_0 = "None"
	_Format_name_1 = "EnumName"
	_Format_name_2 = "EnumValue"
	_Format_name_3 = "Bits"
	_Format_name_4 = "String"
	_Format_name_5 = "Units"
	_Format_name_6 = "All"
)

var (
	_Format_index_0 = [...]uint8{0, 4}
	_Format_index_1 = [...]uint8{0, 8}
	_Format_index_2 = [...]uint8{0, 9}
	_Format_index_3 = [...]uint8{0, 4}
	_Format_index_4 = [...]uint8{0, 6}
	_Format_index_5 = [...]uint8{0, 5}
	_Format_index_6 = [...]uint8{0, 3}
)

func (i Format) String() string {
	switch {
	case i == 0:
		return _Format_name_0
	case i == 2:
		return _Format_name_1
	case i == 4:
		return _Format_name_2
	case i == 8:
		return _Format_name_3
	case i == 16:
		return _Format_name_4
	case i == 32:
		return _Format_name_5
	case i == 255:
		return _Format_name_6
	default:
		return fmt.Sprintf("Format(%d)", i)
	}
}

var _FormatNameToValue_map = map[string]Format{
	_Format_name_0[0:4]: 0,
	_Format_name_1[0:8]: 2,
	_Format_name_2[0:9]: 4,
	_Format_name_3[0:4]: 8,
	_Format_name_4[0:6]: 16,
	_Format_name_5[0:5]: 32,
	_Format_name_6[0:3]: 255,
}

func FormatFromString(s string) (Format, error) {
	if val, ok := _FormatNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Format values", s)
}

func FormatAsList() []Format {
	list := make([]Format, len(_FormatNameToValue_map))
	idx := 0
	for _, v := range _FormatNameToValue_map {
		list[idx] = v
		idx++
	}
	return list
}

func FormatAsListString() []string {
	list := make([]string, len(_FormatNameToValue_map))
	idx := 0
	for k := range _FormatNameToValue_map {
		list[idx] = k
		idx++
	}
	return list
}

func FormatIsValid(t Format) bool {
	for _, v := range FormatAsList() {
		if t == v {
			return true
		}
	}
	return false
}

func (i Format) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Format) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Format should be a string, got %s", data)
	}

	var err error
	*i, err = FormatFromString(s)
	return err
}