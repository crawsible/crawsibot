package irc

type Decoder interface {
	Decode(msgStr string) (*Message, error)
}

type Forwarder struct {
	Decoder        Decoder
	PINGRecipients []PINGRecipient
}

func NewForwarder() *Forwarder {
	return &Forwarder{
		Decoder: &MessageCipher{},
	}
}

type ReadStringer interface {
	ReadString(delim byte) (line string, err error)
}

func (f *Forwarder) StartForwarding(rsr ReadStringer) {
	go f.forward(rsr)
}

func (f *Forwarder) forward(rsr ReadStringer) {
	for {
		msgStr, ok := rsr.ReadString('\n')
		if ok != nil {
			return
		}

		msg, _ := f.Decoder.Decode(msgStr)
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
