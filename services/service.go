package services

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/domain/models/engine"
	"io/ioutil"
	"net"
	"net/http"
)

type Service interface {
	Call(context context.Context, request *adapters.Request) (*engine.Response, error)
}

type apiService struct {
	address string
}

func (a apiService) Call(context context.Context, request *adapters.Request) (*engine.Response, error) {
	jsonReq, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(a.address, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	defer resp.Body.Close()
	bodyBytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	bodyString := string(bodyBytes)
	response := engine.Response{}
	eer := json.Unmarshal([]byte(bodyString), response)
	if eer != nil {
		return nil, eer
	}
	return &response, nil
}

func NewApiService() Service {
	return &apiService{}
}

type socketService struct {
	address string
	service string
}

func (s socketService) Call(context context.Context, request *adapters.Request) (*engine.Response, error) {
	conn, err := net.Dial("tcp", s.address)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return nil, err
	}
	message, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	conn.Write(message)
	res, error := bufio.NewReader(conn).ReadString('\n')
	if error != nil {
		return nil, error
	}
	response := engine.Response{}
	e := json.Unmarshal([]byte(res), response)
	if e != nil {
		return nil, e
	}
	return &response, nil

}

func NewSocketService(address, service string) Service {
	return &socketService{
		address: address,
		service: service,
	}
}
