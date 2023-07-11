package models

type CustomString string

func (cs *CustomString) UnmarshalJSON(data []byte) error {

	str := string(data)
	*cs = CustomString(str)
	return nil
}
