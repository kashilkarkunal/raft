package server

type AppendRequest struct {
    term uint64
    leaderId uint64
    prevLogIndex uint64
    prevLogTerm uint64
    entries []uint64
    leaderCommit uint64
}

type AppendResponse struct {
    term uint64
    success bool
}

type VoteRequest struct {
    term uint64
    candidateId uint64
    lastLogIndex uint64
    lastLogTerm uint64
}

type VoteResponse struct {
    term uint64
    voteGranted bool
}

type Raft interface {
    AppendEntries(request AppendRequest, response *AppendResponse) error

    RequestVote(request VoteRequest, response *VoteResponse) error
}

type State interface {
    RunCmd(cmd string, state *uint64) error
}
