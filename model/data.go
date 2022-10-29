package model

type Data struct {
    Code int64       `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}
