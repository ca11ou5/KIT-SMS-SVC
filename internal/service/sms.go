package service

import (
	"SMS_Service/internal/pb"
	"SMS_Service/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"net/http"
)

type Server struct {
	ApiToken   string
	RestClient *resty.Client
}

type SendCodeResponseBody struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func (s *Server) SendCode(ctx context.Context, req *pb.SendCodeRequest) (*pb.SendCodeResponse, error) {
	fmt.Println("Дошли до сюда")
	code := utils.GenerateCode()
	url := fmt.Sprintf("https://smsgateway24.com/getdata/addsms?token=%s&sendto=%s&body=%s&device_id=11173", s.ApiToken, req.PhoneNumber, code)
	response, err := s.RestClient.R().Get(url)
	if err != nil {
		return &pb.SendCodeResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
			Code:   "",
		}, nil
	}

	b := SendCodeResponseBody{}
	if response.StatusCode() == 200 {
		err = json.Unmarshal([]byte(response.String()), &b)
		if err != nil {
			return &pb.SendCodeResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
				Code:   "",
			}, nil
		}
	} else {
		return &pb.SendCodeResponse{
			Status: http.StatusBadGateway,
			Error:  "sms api doesn't answer",
			Code:   "",
		}, nil
	}

	if b.Error == 1 {
		return &pb.SendCodeResponse{
			Status: http.StatusBadRequest,
			Error:  b.Message,
			Code:   "",
		}, nil
	}

	return &pb.SendCodeResponse{
		Status: http.StatusOK,
		Error:  "",
		Code:   code,
	}, nil

}
