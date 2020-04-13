package export

import "chpunk/translation"

type Exporter interface {
	Export([]*translation.Content) error
}
