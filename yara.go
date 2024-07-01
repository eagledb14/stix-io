package main

type Yara struct {
    Name string
    Indicator []Indicator
}

type Indicator struct {
    Type string
    Data string
}

