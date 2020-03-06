package msg

import (
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
	"middleware-serivce/model"
)

type SneakerProduct struct{}

func (r *SneakerProduct) Decode(input []byte) (*model.SneakerProduct, error) {
	sneakerProduct := &model.SneakerProduct{}
	if err := msgpack.Unmarshal(input, sneakerProduct); err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Decode")
	}
	return sneakerProduct, nil
}

func (r *SneakerProduct) Encode(input *model.SneakerProduct) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.SneakerProduct.Encode")
	}
	return rawMsg, nil
}
