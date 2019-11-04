package worker

import (
	"gorobot/engine"
)

type RobotService struct{}

func (RobotService) Process(
	req Request,
	result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
