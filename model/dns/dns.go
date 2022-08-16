package dns

import (
	"demo/core/errors"
	"strconv"
)

type GetDNSRequest struct {
	X   string `form:"x" binding:"required"`
	Y   string `form:"y" binding:"required"`
	Z   string `form:"z" binding:"required"`
	Vel string `form:"vel" binding:"required"`
}

func (q *GetDNSRequest) Default() {
	// TODO: Default values go here

}

func (q *GetDNSRequest) Validate() *errors.RestError {
	// TODO: Validation goes here
	// I don't like playground validator cause of bad returned field names
	// Prefer building an internal validation library to return meaningful validation errors
	// Opinionated choice, can discuss further
	return nil
}

type FindDNS struct {
	X   float64
	Y   float64
	Z   float64
	Vel float64
}

func (q *GetDNSRequest) Parse() (*FindDNS, error) {
	parsedDNS := FindDNS{}

	var err error
	if parsedDNS.X, err = strconv.ParseFloat(q.X, 64); err != nil {
		return nil, err
	}

	if parsedDNS.Y, err = strconv.ParseFloat(q.Y, 64); err != nil {
		return nil, err
	}

	if parsedDNS.Z, err = strconv.ParseFloat(q.Z, 64); err != nil {
		return nil, err
	}

	if parsedDNS.Vel, err = strconv.ParseFloat(q.Vel, 64); err != nil {
		return nil, err
	}

	return &parsedDNS, nil
}

type GetDNSResponse struct {
	Loc float64 `json:"loc"`
}
