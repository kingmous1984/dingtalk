package model

// JSON 统一JSON处理
type JSON string

func (j *JSON) FromDB(b []byte) error {
	if len(b) > 0 {
		*j = JSON(string(b))
		return nil
	}
	return nil
}
func (j *JSON) ToDB() ([]byte, error) {
	// log.Debug().Msgf("[JSON] %+v", *j)
	str := string(*j)
	// log.Debug().Msgf("[JSON] str=%+v", str)
	if str == "" {
		return []byte("{}"), nil
	}
	return []byte(str), nil
	// return []byte(string(*j)), nil
}
