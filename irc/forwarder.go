package irc

type MessageForwarder struct {
	PINGRcvrs []MsgRcvr
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
		for _, rcp := range f.PINGRcvrs {
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

func (f *MessageForwarder) EnrollForMsgs(mrc MsgRcvr, cmd string) {
	for _, addedMrc := range f.PINGRcvrs {
		if mrc == addedMrc {
			return
		}
	}

	f.PINGRcvrs = append(f.PINGRcvrs, mrc)
}
