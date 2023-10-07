package args

import (
	"github.com/graphql-go/graphql"
)

type ConfigNameEnum int

const (
	MONTHLY_PERIOD ConfigNameEnum = iota
	DAILY_NOTIFY
)

var ProfileConfigEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "ProfileConfigEnum",
	Values: graphql.EnumValueConfigMap{
		"MONTHLY_PERIOD": &graphql.EnumValueConfig{
			Value: "MONTHLY_PERIOD",
		},
		"DAILY_NOTIFY": &graphql.EnumValueConfig{
			Value: "DAILY_NOTIFY",
		},
		// Tambahkan bidang lain yang Anda butuhkan di sini
	},
})
