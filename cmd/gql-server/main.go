package main

import (
    log "github.com/icadpratama/attendance/internal/logger"
    "github.com/icadpratama/attendance/internal/orm"
    "github.com/icadpratama/attendance/pkg/server"
)

func main() {
    mapping, err := orm.Factory()
    if err != nil {
        log.Panic(err)
    }

    server.Run(mapping)
}