//go:build tools
// +build tools

package annict

// workaround for https://github.com/Khan/genqlient/issues/160#issuecomment-1112106276

import (
	_ "github.com/Khan/genqlient"
)
