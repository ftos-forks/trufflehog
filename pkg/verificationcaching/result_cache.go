package verificationcaching

import (
	"github.com/trufflesecurity/trufflehog/v3/pkg/cache"
	"github.com/trufflesecurity/trufflehog/v3/pkg/detectors"
)

// ResultCache is a cache that holds individual detector results.
type ResultCache cache.Cache[detectors.Result]
