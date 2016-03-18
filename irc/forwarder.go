package irc

import "github.com/crawsible/crawsibot/irc/models"

type MessageForwarder struct {
	MsgChs map[string][]chan *models.Message
}

type ReadStringer interface {
	ReadString(delim byte) (line string, err error)
}

type Decoder interface {
	Decode(msgStr string) (*models.Message, error)
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
		for _, msgCh := range f.MsgChs[msg.Command] {
			msgCh <- msg
		}
	}
}

func (f *MessageForwarder) EnrollForMsgs(cmd string) (ch chan *models.Message) {
	ch = make(chan *models.Message, 1)
	f.MsgChs[cmd] = append(f.MsgChs[cmd], ch)
	return
}
