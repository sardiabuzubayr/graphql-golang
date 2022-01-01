package graphql

import (
	"mygraphql/model"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var (
	level = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Level",
			Fields: graphql.Fields{
				"id_level": &graphql.Field{Type: graphql.ID},
				"nama":     &graphql.Field{Type: graphql.String},
			},
			Description: "Data level",
		},
	)
	user = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"email":        &graphql.Field{Type: graphql.ID},
				"nama":         &graphql.Field{Type: graphql.String},
				"no_handphone": &graphql.Field{Type: graphql.String},
				"alamat":       &graphql.Field{Type: graphql.String},
				"ktp":          &graphql.Field{Type: graphql.String},
			},
			Description: "Data users",
		},
	)
)

func levelData() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(level),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			idLevel, _ := p.Args["id_level"].(int)
			listLevel, err := model.GetLevel(idLevel)
			return listLevel, err
		},
		Description: "List Level",
		Args: graphql.FieldConfigArgument{
			"id_level": &graphql.ArgumentConfig{Type: graphql.Int},
		},
	}
}

func usersData(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(user),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			keywords, _ := p.Args["keywords"].(string)
			listUser, err := model.GetUsers(keywords)
			return listUser, err
		},
		Description: "List User",
		Args: graphql.FieldConfigArgument{
			"keywords": &graphql.ArgumentConfig{Type: graphql.String},
		},
	}
}

func newQuery(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": usersData(db),
			"level": levelData(),
		},
	})
}
