package args

import (
	"github.com/graphql-go/graphql"
)

var OrderParamEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "OrderParamEnum",
	Values: graphql.EnumValueConfigMap{
		"DESC": &graphql.EnumValueConfig{
			Value: "desc",
		},
		"ASC": &graphql.EnumValueConfig{
			Value: "asc",
		},
	},
	Description: "",
})

var InfiniteScroll = graphql.FieldConfigArgument{
	"order": &graphql.ArgumentConfig{
		Type:         OrderParamEnum,
		DefaultValue: "desc",
		Description:  "Input order by data, by default DESC",
	},
	"cursor": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
		Description:  "",
	},
}

var OrderInfiniteScroll = InfiniteScroll["order"]
var CursorInfiniteScroll = InfiniteScroll["cursor"]
var PeriodeBool = &graphql.ArgumentConfig{
	Type:         graphql.Boolean,
	DefaultValue: false,
}
var StartTimeParam = &graphql.ArgumentConfig{
	Type:         graphql.String,
	DefaultValue: "",
}
var EndTimeParam = &graphql.ArgumentConfig{
	Type:         graphql.String,
	DefaultValue: "",
}
var TypeParam = graphql.NewEnum(graphql.EnumConfig{
	Name: "TypeParam",
	Values: graphql.EnumValueConfigMap{
		"bulan_ini": &graphql.EnumValueConfig{
			Value: "bulan-ini",
		},
		"minggu_ini": &graphql.EnumValueConfig{
			Value: "bulan-ini",
		},
		"hari_ini": &graphql.EnumValueConfig{
			Value: "bulan-ini",
		},
		"kemarin": &graphql.EnumValueConfig{
			Value: "bulan-ini",
		},
	},
})
