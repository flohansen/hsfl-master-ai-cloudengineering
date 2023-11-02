// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package database

import (
	"database/sql"
)

type Ingredient struct {
	IngredientName   string
	IngredientAmount int64
	IngredientUnit   string
	RecipeID         int64
}

type Profile struct {
	ProfileID      int64
	Username       string
	Password       string
	ProfilePicture []byte
	Bio            sql.NullString
	Friends        sql.NullInt64
	Weekplan       sql.NullInt64
}

type Recipe struct {
	RecipeID      int64
	RecipeName    string
	RecipePicture []byte
	TimeEstimate  sql.NullInt64
	Difficulty    sql.NullString
	FeedsPeople   sql.NullInt64
	Directions    string
	Author        string
}

type RecipeCollection struct {
	RecipeCollectionID   int64
	RecipeCollectionName string
	RecipeID             int64
	OwnerID              int64
	Date                 sql.NullString
	SubscriberID         sql.NullInt64
}
