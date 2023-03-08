package handlers

import (
	"errors"

	"github.com/abinashphulkonwar/dist-cache/storage"
)

type ResData struct {
	Key   string
	Value string
}

func WriteDoc(data *Body, db *storage.BadgerStorage) (*ResData, error) {
	if data.Data.Key == "" {
		return nil, errors.New("key is empty")
	}
	if data.Data.Value == "" {
		return nil, errors.New("value is empty")
	}
	if data.Data.Commend == "" {
		return nil, errors.New("commends is empty")
	}

	if !IsValidCommend(data.Data) {
		return nil, errors.New("invalid commend")
	}

	if data.Data.Commend == DELETE {
		err := db.Delete(data.Data.Key)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	if data.Data.Commend == GET {

		res, err := db.Get(data.Data.Key)
		if err != nil {
			return nil, err
		}

		if res[1] == "" {
			return nil, errors.New("key not found")
		}

		return &ResData{
			Key:   res[0],
			Value: res[1],
		}, nil
	}

	if data.Data.Commend == SET {
		err := db.Add(data.Data.Key, data.Data.Value)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
