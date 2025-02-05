package reponse

import (
	"fmt"
	"time"
)

// type alias
// cannot define new methods on non-local type int, so we use type alias
type JsonTime time.Time

// implement `json.Marshaler` interface in 'encoding/json' package
func (j JsonTime) MarshalJSON() ([]byte, error){
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id uint64 `json:"id"`
	NickName string `json:"name"`
	//Birthday string `json:"birthday"`
	Birthday JsonTime `json:"birthday"` // auto convertion
	Gender string `json:"gender"`
	Mobile string `json:"mobile"`
}
