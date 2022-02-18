// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    error, err := UnmarshalError(bytes)
//    bytes, err = error.Marshal()

package types

import "encoding/json"

func UnmarshalError(data []byte) (Error, error) {
	var r Error
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Error) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Error struct {
	Code    string      `json:"code"`   
	Details interface{} `json:"details"`
	Msg     string      `json:"msg"`    
}
