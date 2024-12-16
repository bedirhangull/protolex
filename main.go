package protolex

import (
	"github.com/bedirhangull/protolex/internal/adapter/config"
	"github.com/bedirhangull/protolex/internal/adapter/prescriptive"
	"github.com/bedirhangull/protolex/internal/core/parser"
	"github.com/bedirhangull/protolex/internal/core/proto"
)

func ReadProtoFile(protoFilePath string) (*proto.Proto, error) {
	cfg, err := config.NewConfig(protoFilePath)
	if err != nil {
		return nil, err
	}

	protoContent, err := config.ReadProtoFile(cfg.ProtoPath)
	if err != nil {
		return nil, err
	}

	/*
		Remove all comments, tabs, carriage returns, escape characters and spaces from the proto file
		This is a prescriptive approach to make the proto file more readable
		Also if you want to keep specific characters, you can remove the corresponding function
	*/
	pre := prescriptive.NewPrescriptive(protoContent)
	pre.CleanContent()

	parser := &parser.Parser{
		Prescriptive: pre,
	}

	protoInstance := proto.NewProto(parser)

	return protoInstance, nil
}
