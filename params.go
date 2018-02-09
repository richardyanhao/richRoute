package richRoute

import "errors"

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) GetValByKey(Key string) (string, error) {
	for _, v := range ps {
		if  v.Key == Key {
			return v.Value, nil
		}
	}
	return "", errors.New("not existed")
}