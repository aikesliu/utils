package proto

import (
	"github.com/aikesliu/utils/log"
	"github.com/golang/protobuf/proto"
)

func GetData(msg proto.Message) []byte {
	if msg == nil {
		return nil
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.E("marshal failed: %v", err)
		return nil
	}
	return data
}
