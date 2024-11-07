package jetfy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/rest-api/configs"
	"github.com/rest-api/infra/logger"
)

type JetfyPayload struct {
	SenderId  string `json:"sender_id"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendSms(message, recipient string) error {
	var response Response

	payload := JetfyPayload{
		SenderId:  configs.JetfySetting.SenderId,
		Recipient: recipient,
		Message:   message,
	}

	resp, err := resty.
		New().
		R().
		SetHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", configs.JetfySetting.Tooken),
			"Content-Type":  "application/json",
		}).
		SetBody(payload).
		Post(fmt.Sprintf("%s/%s", configs.JetfySetting.BaseUrl, "api/v1/sms/send"))

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNoContent {
		return nil
	}

	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		logger.Error(err)
		return err
	}

	logger.Error(response.Message)
	return errors.New(response.Message)
}
