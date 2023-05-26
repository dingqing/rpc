package global

import (
	"github.com/dingqing/rpc/codec"
	"github.com/dingqing/rpc/protocol"
)

var Codecs = map[protocol.SerializeType]codec.Codec{
	protocol.JSON: &codec.JSONCodec{},
	protocol.Gob:  &codec.GobCodec{},
}
