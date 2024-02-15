package systems

import (
	"mono.thienhang.com/pkg/service"
)

func (s *System) InitPlugin(srv service.List) {
	route := s.App.Group(s.GroupName)
	//
	route.GET("/ping", s.infoHandler)
	route.GET("/info", s.infoHandler)
	route.GET("/metrics", s.prometheusHandler)
	route.GET("/version", s.versionandler)
}
