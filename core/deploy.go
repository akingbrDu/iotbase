package core

import "encoding/json"

type DeployCommunication struct {
	Type  string          `json:"type"`
	Param json.RawMessage `json:"param"`
}

type DeployProto struct {
	Type  string          `json:"type"`
	Param json.RawMessage `json:"param"`
}

type DeployData struct {
	Props    json.RawMessage `json:"props"`
	Commands json.RawMessage `json:"commands"`
	Events   json.RawMessage `json:"events"`
	Setting  json.RawMessage `json:"setting"`
}

type DeployAriot struct {
	Communication DeployCommunication `json:"commu"`
	Runtime       json.RawMessage     `json:"runtime"`
	Proto         DeployProto         `json:"proto"`
	Data          DeployData          `json:"data"`
}
