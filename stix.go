package main

import (
    "encoding/json"
    "regexp"
    "strings"
    "errors"
)

type Bundle struct {
    Object []Object `json:"objects"`
}

type Object struct {
    Type string `json:"type"`
    Pattern string `json:"pattern"`
    Name string `json:"name"`
}

func (b Bundle) ToYara() Yara {
    out := []Indicator{}
    for _, obj := range b.Object {
        if obj.Pattern == "[]" {
            continue
        }

        obj.Pattern = strings.TrimLeft(obj.Pattern, "[")
        obj.Pattern = strings.TrimRight(obj.Pattern, "]")

        re := regexp.MustCompile(` AND | OR `)
        strSplit := re.Split(obj.Pattern, -1)

        for _, s := range strSplit {
            s = strings.TrimLeft(s, "(")
            s = strings.TrimRight(s, ")")
            type_, value, data := parsePattern(s)
            out = append(out, Indicator{
                Name: obj.Name,
                Type: type_,
                Value: value,
                Data: data,
            })
        }
    }

    return Yara{
        Indicator: out,
    }
} 

func parsePattern(pattern string) (string, string, string) {
    re := regexp.MustCompile(`:| = | MATCHES | ISSUBSET `)
    strs := re.Split(pattern, -1)

    trimmedData := strings.Trim(strs[len(strs) - 1], "'")
    return strs[0], strs[1], trimmedData
}

func Unmarshall(input string) (Bundle, error) {
    out := Bundle{}
    err := json.Unmarshal([]byte(input), &out)
    if err != nil {
        return out, errors.New("Invalid JSON, Input valid STIX Bundle") 
    }

    filteredObjects := []Object{}
    for _, o := range out.Object {
        if o.Type == "indicator" {
            filteredObjects = append(filteredObjects, o)
        }
    }
    out.Object = filteredObjects
    return out, nil
}
