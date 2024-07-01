package main

import (
	"encoding/json"
)

type Bundle struct {
    Object []Object `json:"objects"`
}

type Object struct {
    Type string `json:"type"`
    Pattern string `json:"pattern"`
}

func Unmarshall(input string) Bundle {
    out := Bundle{}
    json.Unmarshal([]byte(input), &out)

    filteredObjects := []Object{}
    for _, o := range out.Object {
        if o.Type == "indicator" {
            filteredObjects = append(filteredObjects, o)
        }
    }
    out.Object = filteredObjects
    return out
}
