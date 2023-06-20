package jsonutil

import (
	"encoding/json"
	"io"
)

func WriteJSONMessage(conn io.Writer, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(string(data) + "\n"))
	if err != nil {
		return err
	}
	return nil
}

func ReadJSONMessage(conn io.Reader, message interface{}) error {
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(message)
	if err != nil {
		return err
	}
	return nil
}
