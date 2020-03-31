package tracker

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	GetTicket   = "v2/issues/"
	PatchTicket = "v2/issues/"
)

type Ticket map[string]interface{}

// Получить yandex пользователя автора тикета
func (t Ticket) CreatedBy() User {
	if createdBy, ok := t["createdBy"].(map[string]interface{}); ok {
		return User{
			Self:    interfaceToString(createdBy["self"]),
			ID:      interfaceToString(createdBy["id"]),
			Display: interfaceToString(createdBy["display"]),
		}
	}

	return User{}
}

// Получить yandex пользователя исполнителя тикета
func (t Ticket) Assignee() User {
	if assignee, ok := t["assignee"].(map[string]interface{}); ok {
		return User{
			Self:    interfaceToString(assignee["self"]),
			ID:      interfaceToString(assignee["id"]),
			Display: interfaceToString(assignee["display"]),
		}
	}

	return User{}
}

// получить массив наблюдателей в тикете
func (t Ticket) Followers() Users {
	var users Users
	if followers, ok := t["followers"].([]interface{}); ok {
		for _, user := range followers {
			users = append(users, User{
				Self:    interfaceToString(user.(map[string]interface{})["self"]),
				ID:      interfaceToString(user.(map[string]interface{})["id"]),
				Display: interfaceToString(user.(map[string]interface{})["display"]),
			})
		}
	}

	return users
}

// Получить заголовок тикета
func (t Ticket) Summary() string {
	if summary, ok := t["summary"]; ok {
		return interfaceToString(summary)
	}

	return ""
}

// Получить номер тикета
func (t Ticket) Key() string {
	if key, ok := t["key"]; ok {
		return interfaceToString(key)
	}

	return ""
}

// Получить описание тикета
func (t Ticket) Description() string {
	if description, ok := t["description"]; ok {
		return interfaceToString(description)
	}

	return ""
}

// Получить статус тикета
func (t Ticket) Status() string {
	if status, ok := t["status"].(map[string]interface{}); ok {
		if display, ok := status["display"]; ok {
			return interfaceToString(display)
		}
	}

	return ""
}

// получить информацию о созданном тикете
func (api *Client) GetTicket(tickerKey string) (*Ticket, error) {
	req := api.createRequest(http.MethodGet, map[string]string{})
	resp, err := api.do(req, GetTicket+tickerKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return unmarshalTicket(resp.Body)
}

// обновить поля в тиките
func (api *Client) PatchTicket(tickerKey string, body map[string]string) (*Ticket, error) {
	req := api.createRequest(http.MethodPatch, body)
	resp, err := api.do(req, PatchTicket+tickerKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return unmarshalTicket(resp.Body)
}

func unmarshalTicket(body io.ReadCloser) (*Ticket, error) {
	result := &Ticket{}

	respBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// создать новый тикет
func (api *Client) CreateTicket() {

}

// Получить время отправки сообщения тикета в слак
func (t Ticket) SlackMessageID() string {
	if description, ok := t["slackMessageID"]; ok {
		return interfaceToString(description)
	}

	return ""
}
