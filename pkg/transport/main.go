package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Fetch basically exposes you a good api to make interacting with external apis a bit easier.

type FetchParams struct {
	// Endpoint URL
	Url string

	// HTTP Method
	Method string

	// Body
	Body interface{}

	// Error Json Object
	ErrorObj interface{}

	// Success Json Object
	SuccessObj interface{}

	// Headers
	Headers *map[string]string

	// This makes the function return logs
	ShouldLog *bool
}

func logData(shouldLog *bool) func(input ...interface{}) {
	return func(input ...interface{}) {
		if shouldLog != nil {
			stringLogs := ""
			for _, log := range input {
				stringLogs += fmt.Sprintf("%s", log)
			}
			fmt.Print("Fetch Logs :: ", stringLogs, "\n")
		}
	}
}

func Fetch(input FetchParams) error {

	log := logData(input.ShouldLog)

	var request *http.Request

	if input.Body != nil {
		jsonBytes, marshalJsonErr := json.Marshal(input.Body)

		if marshalJsonErr != nil {
			log("Encoding Request Body Failed: ", marshalJsonErr)
			return marshalJsonErr
		}

		tempRequest, requestErr := http.NewRequest(input.Method, input.Url, bytes.NewBuffer(jsonBytes))
		if requestErr != nil {
			log("Creating Request With Body Failed: ", requestErr)
			return requestErr
		}

		request = tempRequest
	} else {

		tempRequest, requestErr := http.NewRequest(input.Method, input.Url, nil)
		if requestErr != nil {
			log("Creating Request Without Body Failed: ", requestErr)
			return requestErr
		}

		request = tempRequest
	}

	if input.Headers != nil {
		for key, value := range *input.Headers {
			request.Header.Add(key, value)
		}
	}

	response, responseErr := http.DefaultClient.Do(request)
	if responseErr != nil {
		log("Creating Response Failed: ", responseErr)
		return responseErr
	}

	defer response.Body.Close()
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		if input.SuccessObj != nil {
			jsonErr := json.NewDecoder(response.Body).Decode(input.SuccessObj)
			if jsonErr != nil {
				log("Decoding Response Failed For Success: ", jsonErr)
				return jsonErr
			}
		}
		return nil
	}

	if input.ErrorObj != nil {
		jsonErr := json.NewDecoder(response.Body).Decode(input.ErrorObj)
		if jsonErr != nil {
			log("Decoding Response Failed For Error: ", jsonErr)
			return jsonErr
		}
	}

	return nil
}
