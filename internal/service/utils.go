package service

import (
	"bytes"
	"encoding/json"
	"fr/internal/repository/models"
	"net/http"
	"strconv"
)

const (
	jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA1NDM3MTgsImlzcyI6ImZhYnJpcXVlIiwibmFtZSI6Imh0dHBzOi8vdC5tZS92YW1wX3NoMDB0YSJ9.bTbqxi5ixRVg3WAVMEjgrBInQkq2cwMsLje_anY_yN8"
)

type ResponseBody struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func sendMsg(msg_user models.MessageFull) (*ResponseBody, error) {
	body := struct {
		Id    int
		Phone int
		Text  string
	}{
		Id:    msg_user.MessageId,
		Phone: msg_user.TelNumber,
		Text:  msg_user.Text,
	}
	postBody, _ := json.Marshal(body)
	responseBody := bytes.NewBuffer(postBody)

	url := "https://probe.fbrq.cloud/v1/send/"
	req, err := http.NewRequest("POST", url+strconv.Itoa(msg_user.MessageId), responseBody)
	req.Header.Add("Authorization", jwt)

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var response ResponseBody

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
