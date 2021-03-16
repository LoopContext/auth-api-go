package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

const MaxBytes = 1048576

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (mr *Response) Error() string {
	return mr.Msg
}

// SendJSON sends the interface obj as JSON on the res provided
func SendJSON(res http.ResponseWriter, req *http.Request, dst interface{}, statusCodes ...int) error {
	res.Header().Set("Content-Type", "application/json")
	if len(statusCodes) == 1 {
		res.WriteHeader(statusCodes[0])
	}

	return json.NewEncoder(res).Encode(dst)
}

// ReadJSON reads the JSON into dst interface obj of the req provided, handlesmost JSON errors
func ReadJSON(res http.ResponseWriter, req *http.Request, dst interface{}) error {
	if req.Header.Get("Content-Type") != "application/json" {
		return &Response{Status: http.StatusUnsupportedMediaType, Msg: "Content-Type header is not application/json"}
	}
	req.Body = http.MaxBytesReader(res, req.Body, MaxBytes)
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxError):
			return &Response{
				Status: http.StatusBadRequest,
				Msg:    fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset),
			}
		case errors.Is(err, io.ErrUnexpectedEOF):
			return &Response{
				Status: http.StatusBadRequest,
				Msg:    "Request body contains badly-formed JSON",
			}
		case errors.As(err, &unmarshalTypeError):
			return &Response{
				Status: http.StatusBadRequest,
				Msg: fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)",
					unmarshalTypeError.Field, unmarshalTypeError.Offset),
			}
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			return &Response{
				Status: http.StatusBadRequest,
				Msg: fmt.Sprintf("Request body contains unknown field %s", strings.TrimPrefix(err.Error(),
					"json: unknown field ")),
			}
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &Response{Status: http.StatusBadRequest, Msg: msg}

		case err.Error() == "http: request body too large":
			return &Response{
				Status: http.StatusRequestEntityTooLarge,
				Msg:    fmt.Sprintf("Request body must not be larger than %d bytes", MaxBytes),
			}
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return &Response{Status: http.StatusBadRequest, Msg: "Request body must contain only a single JSON object"}
	}

	return nil
}

// HandleErr Handles errors and sends a JSON response
func HandleErr(res http.ResponseWriter, req *http.Request, err error, statusCode ...int) {
	if err != nil {
		log.Err(err).Send()
		code := http.StatusInternalServerError
		if len(statusCode) > 0 {
			code = statusCode[0]
		}
		var mr *Response
		if !errors.As(err, &mr) {
			mr = &Response{
				Status: code,
				Msg:    err.Error(),
			}
		}
		err = SendJSON(res, req, *mr, code)
		if err != nil {
			log.Err(err).Send()
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
}
