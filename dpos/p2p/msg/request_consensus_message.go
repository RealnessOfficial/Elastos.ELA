package msg

import (
	"io"

	"github.com/elastos/Elastos.ELA.Utility/common"
)

type RequestConsensusMessage struct {
	Height uint32
}

func (msg *RequestConsensusMessage) CMD() string {
	return RequestConsensus
}

func (msg *RequestConsensusMessage) MaxLength() uint32 {
	return 4
}

func (msg *RequestConsensusMessage) Serialize(w io.Writer) error {
	if err := common.WriteUint32(w, msg.Height); err != nil {
		return err
	}

	return nil
}

func (msg *RequestConsensusMessage) Deserialize(r io.Reader) error {
	var err error
	if msg.Height, err = common.ReadUint32(r); err != nil {
		return err
	}

	return nil
}