package export

import "chpunk/internal/translation"

type Exporter interface {
	Export([]*translation.Content) error
}
