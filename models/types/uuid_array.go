package types

import (
	"errors"
	"fmt"
	"strings"
	
	"github.com/google/uuid"
	"database/sql/driver"
)

type UUIDArray []uuid.UUID
func (a *UUIDArray) Scan(value interface{}) error {

	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("failed to parse UUIDArray: unsupport datatype")
	}

	str = strings.TrimPrefix(str,"{")
	str = strings.TrimSuffix(str,"}")
	parts := strings.Split(str,",")

	*a = make(UUIDArray, 0, len(parts))
	for _, s := range parts{
		s = strings.TrimSpace(strings.Trim(s,`"`))
		if s == "" {
			continue
		}
		u, err := uuid.Parse(s)
		if err != nil {
			return fmt.Errorf("invalid UUID in array : %v", err)
		}
		*a = append(*a, u)
	}
	return nil
}

func (a UUIDArray) Value()(driver.Value, error)  {
	if len(a) == 0 {
		return "{}",nil
	}

	postgresFormat := make([]string,0, len(a))

	for _, value := range a{
		postgresFormat = append(postgresFormat, fmt.Sprintf(`"%s"`, value.String()))
	}
	return "{"+ strings.Join(postgresFormat,",") +"}", nil
}

func (UUIDArray) GormDataType() string{
	return "uuid[]"
}