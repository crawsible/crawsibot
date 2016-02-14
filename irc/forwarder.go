package irc

type MessageForwarder struct {
	PINGRcvrs          []MsgRcvr
	RPL_ENDOFMOTDRcvrs []MsgRcvr
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

		var rcvrs []MsgRcvr
		switch msg.Command {
		case "PING":
			rcvrs = f.PINGRcvrs
		case "RPL_ENDOFMOTD":
			rcvrs = f.RPL_ENDOFMOTDRcvrs
		default:
			rcvrs = []MsgRcvr{}
		}

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

func (f *MessageForwarder) EnrollForMsgs(mrc MsgRcvr, cmd string) {
	switch cmd {
	case "PING":
		f.PINGRcvrs = appendIfNew(f.PINGRcvrs, mrc)
	case "RPL_ENDOFMOTD":
		f.RPL_ENDOFMOTDRcvrs = appendIfNew(f.RPL_ENDOFMOTDRcvrs, mrc)
	}
}

func appendIfNew(mrcs []MsgRcvr, mrc MsgRcvr) []MsgRcvr {
	for _, addedMrc := range mrcs {
		if mrc == addedMrc {
			return mrcs
		}
	}

	return append(mrcs, mrc)
}
