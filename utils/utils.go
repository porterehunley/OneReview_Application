package utils

//This file contains useful helper functions to build json messages and return the responses

import (
	"encoding/json"
	"net/http"
)


//The interface{} is just a representation for an arbitrary data type
//The map[string]interface{} maps strings onto values, like JSON
func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) error {

	//Adding the content type to the response header
	w.Header().Add("Content-Type", "application/json")

	//An encoder writes JSON values to an output stream.
	//New encoder creates an encoder that writes to the http.ResponseWriter output stream
	//The Encode method parses a JSON value to a bytestream for output to io
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		return err
	}
	return nil
}