package dto

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type EventParseFunc func(event *WSPayload, message []byte) error

type EventParse struct {
	funcMap map[OPCode]map[EventType]EventParseFunc
	intent  Intent
}

func NewEventParse() *EventParse {
	return &EventParse{
		funcMap: map[OPCode]map[EventType]EventParseFunc{
			WSDispatchEvent: {},
		},
	}
}

func (e *EventParse) FuncMap() map[OPCode]map[EventType]EventParseFunc {
	return e.funcMap
}

func (e *EventParse) AtMessage(handler ATMessageEventHandler) *EventParse {
	e.funcMap[WSDispatchEvent][EventAtMessageCreate] = func(event *WSPayload, message []byte) error {
		data := &WSATMessageData{}
		if err := parseData(message, data); err != nil {
			return err
		}
		return handler(event, data)
	}
	e.intent = e.intent | EventToIntent(EventAtMessageCreate)
	return e
}

func (e *EventParse) Intent() Intent {
	return e.intent
}

func parseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	return json.Unmarshal([]byte(data.String()), target)
}
