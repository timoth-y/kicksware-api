package service

import model "beta-service/core/model"

type BetaSerializer interface {
	Decode(input []byte) (*model.Beta, error)
	DecodeRange(input []byte) ([]*model.Beta, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
