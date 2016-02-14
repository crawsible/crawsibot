package irc

type Forwarder struct {
	PINGRecipients []PINGRecipient
}

type ReadStringer interface {
	ReadString(delim byte) (line string, err error)
}

type Decoder interface {
	Decode(msgStr string) (*Message, error)
}

func (f *Forwarder) StartForwarding(rsr ReadStringer, dcdr Decoder) {
	go f.forward(rsr, dcdr)
}

func (f *Forwarder) forward(rsr ReadStringer, dcdr Decoder) {
	for {
		msgStr, ok := rsr.ReadString('\n')
		if ok != nil {
			return
		}

		msg, _ := dcdr.Decode(msgStr)
		for _, rcp := range f.PINGRecipients {
			rcp.RcvPING(
				msg.NickOrSrvname,
				msg.FirstParams,
				msg.Params,
			)
		}
	}
}

type PINGRecipient interface {
	RcvPING(nick, fprms, prms string)
}

func (f *Forwarder) EnrollForPING(rcp PINGRecipient) {
	for _, addedRcp := range f.PINGRecipients {
		if rcp == addedRcp {
			return
		}
	}

	f.PINGRecipients = append(f.PINGRecipients, rcp)
}
