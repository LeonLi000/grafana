package sqlstore

import (
	"github.com/LeonLi000/grafana/pkg/bus"
	m "github.com/LeonLi000/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
