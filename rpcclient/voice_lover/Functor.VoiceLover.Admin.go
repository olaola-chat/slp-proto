package voice_lover

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"
)

// VoiceLoverAdmin *voiceLoverAdmin
var VoiceLoverAdmin = &voiceLoverAdmin{
	&base.Base{
		Name: "Functor.VoiceLover.Admin",
	},
}

type voiceLoverAdmin struct {
	*base.Base
}
