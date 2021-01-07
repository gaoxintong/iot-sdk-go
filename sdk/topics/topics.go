package topics

import (
	"fmt"

	"github.com/imdario/mergo"
)

// Topics 主题
type Topics struct {
	Register     string
	Login        string
	PostProperty string
	SetProperty  string
	PostEvent    string
	OnCommand    string
}

// DefaultTopics 默认主题列表
var DefaultTopics = Topics{
	Register:     "/v1/devices/registration",
	Login:        "/v1/devices/authentication",
	PostProperty: "s",
	SetProperty:  "",
	PostEvent:    "e",
	OnCommand:    "c",
}

// Override 合并默认主题列表
func Override(n Topics) Topics {
	if err := mergo.Merge(&DefaultTopics, n, mergo.WithOverride); err != nil {
		fmt.Println(err)
	}
	return DefaultTopics
}
