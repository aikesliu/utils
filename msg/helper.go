package msg

import (
	"github.com/golang/protobuf/proto"
	"utils/logger"
)

func GetData(msg proto.Message) []byte {
	if msg == nil {
		return nil
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		logger.E("marshal msg failed, err: %v", err)
		return nil
	}
	return data
}
