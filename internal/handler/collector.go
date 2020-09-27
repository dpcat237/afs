package handler

import (
	"github.com/dpcat237/afs/internal/interfaces"
)

//Collector contains handlers.
type Collector struct {
	Link interfaces.LinkHandler
}

//NewCollector initializes handlers.
func NewCollector() Collector {
	return Collector{
		Link: newLink(),
	}
}
