package control


import (
	"net/http"
	"kboard/core"
)

type CtlIndex struct {
	Control
}

func NewCtlIndex(config *core.Config, w http.ResponseWriter, r *http.Request) *CtlIndex {
	return &CtlIndex{
		Control{
			Config: config,
			TplEngine: core.NewTplEngine(w, r),
			Control: "index",
			Actions: map[string]func(){},
		},
	}
}

func (this *CtlIndex) Index() {
	this.TplEngine.Display("index")
}