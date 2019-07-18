package handlers

import (
	"fmt"
	"encoding/json"
	"github.com/twitchtv/twirp"
)

// TODO define errors that return from the server
// Switch case them here and custom make the error code

// errFormater represents a twirp err formatter
type errFormater func(err error) []byte

// handleError is an errFormatter
// handleError is passed to the respond wrapper to format errors
func handleError(err error) []byte {
	twirpErr := twirp.NewError(twirp.NotFound, err.Error())
	errorMessage := map[string]string{
		"Message": twirpErr.Msg(),
		"Status": string(twirpErr.Code()),
	}
	errToReturn, _ := json.Marshal(errorMessage)
	if err != nil {
		fmt.Println(err)
	}
	return errToReturn
}