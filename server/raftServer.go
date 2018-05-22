package server

import (
    "fmt"
    "strings"
    "strconv"
)

type RaftServer struct {
    replicatedState uint64
}

var _ State = new(RaftServer)

func NewServer() *RaftServer{
    return &RaftServer{0}
}

func(self *RaftServer) RunCmd(cmd string, state *uint64) error {
    splitCmd := strings.Split(cmd, " ")
    op := splitCmd[0]
    val, e := strconv.ParseUint(splitCmd[1], 10, 64)
    if e != nil {
        return fmt.Errorf("Invalid value")
    }

    var err error = nil

    switch op {
    case "add":
        self.replicatedState += val
        *state = self.replicatedState
    case "mul":
        self.replicatedState *= val
        *state = self.replicatedState
    case "set":
        self.replicatedState = val
        *state = self.replicatedState
    default:
        err = fmt.Errorf("Invalid Operation. Supported Ops - add, mul, set")
    }

    return err
}
