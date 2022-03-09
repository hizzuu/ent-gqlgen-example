// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "contents", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "image_posts", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "user_posts", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_images_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "post_title",
				Unique:  false,
				Columns: []*schema.Column{PostsColumns[1]},
			},
			{
				Name:    "post_contents",
				Unique:  false,
				Columns: []*schema.Column{PostsColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeString, Unique: true},
		{Name: "nickname", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"GENERAL", "OFFICIAL"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "image_users", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_images_users",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_nickname",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[2]},
			},
			{
				Name:    "user_uid",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ImagesTable,
		PostsTable,
		UsersTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = ImagesTable
	PostsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = ImagesTable
}
