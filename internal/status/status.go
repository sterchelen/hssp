package status

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	_ "embed"
)

const (
	Informal = iota + 1
	Successful
	Redirection
	ClientError
	ServerError
)

type (
	Status struct {
		Code        int
		class       int
		Description string
		RFCLink     string
	}

	Statuses []*Status
)

//go:embed http-statuses.csv
var b []byte

//Initialize reads embedded http-statuses csv file to fill inner structure
//and returns an array of status
func Initialize() (Statuses, error) {
	s, err := fillStatuses()
	if err != nil {
		return nil, err
	}
	return s, nil
}

//GiveClassName returns class name based on the given status code
func (s Status) GiveClassName() string {
	switch s.class {
	case Informal:
		return "Informal"
	case Successful:
		return "Successful"
	case Redirection:
		return "Redirection"
	case ClientError:
		return "Client Error"
	case ServerError:
		return "Server Error"
	}

	return "Unassigned"
}

func CodeClassFromArg(arg string) (int, bool) {
	if class, err := strconv.Atoi(arg); err == nil {
		return class, true
	}

	switch strings.ToLower(arg) {
	case "1xx", "informal":
		return Informal, true
	case "2xx", "successful":
		return Successful, true
	case "3xx", "redirection":
		return Redirection, true
	case "4xx", "clienterror", "client error":
		return ClientError, true
	case "5xx", "servererror", "server error":
		return ServerError, true
	}

	return 0, false
}

//StatusesByClass returns all the statuses fulfilling the given class condition
func (s Statuses) StatusesByClass(class int) (Statuses, error) {
	if class < 1 || class > 5 {
		return nil, fmt.Errorf("Class undefined")
	}

	var codes Statuses
	for _, status := range s {
		if status.class == class {
			codes = append(codes, status)
		}
	}

	return codes, nil
}

//FindStatusByCode returns a list of statuses based on the given code
func (s Statuses) FindStatusesByCode(code int) (Statuses, error) {
	var statuses Statuses
	for _, status := range s {
		if status.Code == code {
			statuses = append(statuses, status)
		}
	}
	if len(statuses) > 0 {
		return statuses, nil
	}
	return nil, fmt.Errorf("Code undefined")
}

func fillStatuses() (Statuses, error) {
	r := bytes.NewReader(b)
	reader := csv.NewReader(r)
	var statuses Statuses

	//Read csv and fill data structure
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		code, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, fmt.Errorf("csv: %w", err)
		}

		s := &Status{
			Code:        code,
			class:       code / 100,
			Description: line[1],
			RFCLink:     line[2],
		}
		statuses = append(statuses, s)
	}

	return statuses, nil
}
