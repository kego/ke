package assets

import (
	"net/http"

	"frizz.io/edit/assets/assets"
)

// Assets contains project assets.
var Assets http.FileSystem = assets.Assets
