package controller

import (
	"github.com/dpcat237/afs/internal/handler"
	"github.com/dpcat237/afs/internal/interfaces"
)

//Collector contains HTTP controllers.
type Collector struct {
	Controller interfaces.Controller
	Link       interfaces.LinkController
}

//NewCollector initializes a new Collector with HTTP controllers.
func NewCollector(lgr interfaces.Logger, hnds handler.Collector) Collector {
	cnt := newController(lgr)
	return Collector{
		Controller: cnt,
		Link:       newLink(cnt, lgr, hnds.Link),
	}
}
