package main

import (
	"encoding/json"
    "regexp"
    "fmt"
)

type Bundle struct {
    Object []Object `json:"objects"`
}

type Object struct {
    Type string `json:"type"`
    Pattern string `json:"pattern"`
}

func (b Bundle) ToYara() Yara {
    out := []Indicator{}
    for _, o := range b.Object {
        fmt.Println(parsePattern(o.Pattern))

        // out = append(out, Indicator{
        //     
        // })
    }

    return Yara{Indicator: out}
} 

func parsePattern(pattern string) (string, string) {
    re := regexp.MustCompile(`[:| = ]`)
    strs := re.Split(pattern, -1)

    return strs[0], strs[len(strs) - 1]
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
