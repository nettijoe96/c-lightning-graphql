package lightning

import (
	"github.com/niftynei/glightning/glightning"
)

var lightning *glightning.Lightning = glightning.NewLightning()


func GetGlobalLightning() *glightning.Lightning {
	return lightning
}



