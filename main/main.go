package main

import (
    "Qianshou/dao"
    "Qianshou/router"
)

func main() {
    dao.Init()
    r := router.InitRouter()
    r.Run(":80")
}
