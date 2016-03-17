package irc

type MessageForwarder struct {
	MsgRcvrs map[string][]MsgRcvr
}

type ReadStringer interface {
	ReadString(delim byte) (line string, err error)
}

type Decoder interface {
	Decode(msgStr string) (*Message, error)
}

func (f *MessageForwarder) StartForwarding(rsr ReadStringer, dcdr Decoder) {
	go f.forward(rsr, dcdr)
}

func (f *MessageForwarder) forward(rsr ReadStringer, dcdr Decoder) {
	for {
		msgStr, ok := rsr.ReadString('\n')
		if ok != nil {
			return
		}

		msg, _ := dcdr.Decode(msgStr)
		rcvrs := f.MsgRcvrs[msg.Command]
		for _, rcp := range rcvrs {
			rcp.RcvMsg(
				msg.NickOrSrvname,
				msg.FirstParams,
				msg.Params,
			)
		}
	}
}

type MsgRcvr interface {
	RcvMsg(nick, fprms, prms string)
}

//func (f *MessageForwarder) EnrollForMsgs(mrc MsgRcvr, cmd string) {
//f.MsgRcvrs[cmd] = appendIfNew(f.MsgRcvrs[cmd], mrc)
//}

func (f *MessageForwarder) EnrollForMsgs(rcvCh chan map[string]string, cmd string) {}

func appendIfNew(mrcs []MsgRcvr, mrc MsgRcvr) []MsgRcvr {
	for _, addedMrc := range mrcs {
		if mrc == addedMrc {
			return mrcs
		}
	}

	return append(mrcs, mrc)
}
