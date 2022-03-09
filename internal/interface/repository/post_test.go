package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/hizzuu/plate-backend/test/utils"
	"github.com/hizzuu/plate-backend/utils/pointer"
)

func TestPostGet(t *testing.T) {
	t.Parallel()

	type args struct {
		id int
	}
	tests := map[string]struct {
		args    args
		want    *domain.Post
		wantErr bool
		errStr  string
	}{
		"get post": {
			args: args{id: 1},
			want: &domain.Post{ID: 1},
		},
		"not found post": {
			args: args{
				id: 999,
			},
			wantErr: true,
			errStr:  "ent: post not found",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client, teardown := utils.NewDBClient(t)
			defer teardown()

			r := NewPostRepository(client)

			got, err := r.Get(context.Background(), tt.args.id)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.ID, got.ID)
			}

			if tt.wantErr {
				assert.EqualError(t, err, tt.errStr)
			}
		})
	}
}

func TestPostList(t *testing.T) {
	t.Parallel()

	type args struct {
		after  *domain.Cursor
		first  *int
		before *domain.Cursor
		last   *int
		where  *domain.PostWhereInput
	}
	tests := map[string]struct {
		args    args
		want    *domain.PostConnection
		wantErr bool
		errStr  string
	}{
		"get posts list (first)": {
			args: args{
				first: pointer.Int(5),
			},
			want: &domain.PostConnection{
				Edges: []*domain.PostEdge{
					{
						Node: &domain.Post{
							ID:       1,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       2,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       3,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       4,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       5,
							Title:    "test title",
							Contents: "test contents",
						},
					},
				},
			},
		},
		"get posts list (last)": {
			args: args{
				last: pointer.Int(5),
			},
			want: &domain.PostConnection{
				Edges: []*domain.PostEdge{
					{
						Node: &domain.Post{
							ID:       1,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       2,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       3,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       4,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       5,
							Title:    "test title",
							Contents: "test contents",
						},
					},
				},
			},
		},
		"get posts list (where)": {
			args: args{
				where: &domain.PostWhereInput{
					And: []*domain.PostWhereInput{
						{IDIn: []int{1, 2, 3, 4, 5}},
						{Title: pointer.String("test title")},
						{Contents: pointer.String("test contents")},
					},
				},
			},
			want: &domain.PostConnection{
				Edges: []*domain.PostEdge{
					{
						Node: &domain.Post{
							ID:       1,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       2,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       3,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       4,
							Title:    "test title",
							Contents: "test contents",
						},
					},
					{
						Node: &domain.Post{
							ID:       5,
							Title:    "test title",
							Contents: "test contents",
						},
					},
				},
			},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client, teardown := utils.NewDBClient(t)
			defer teardown()

			r := NewPostRepository(client)
			got, err := r.List(context.Background(), tt.args.after, tt.args.first, tt.args.before, tt.args.last, tt.args.where)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.want.Edges), len(got.Edges))
				for i := range got.Edges {
					assert.Equal(t, tt.want.Edges[i].Node.ID, got.Edges[i].Node.ID)
					assert.Equal(t, tt.want.Edges[i].Node.Title, got.Edges[i].Node.Title)
					assert.Equal(t, tt.want.Edges[i].Node.Contents, got.Edges[i].Node.Contents)
				}
			}

			if tt.wantErr {
				assert.EqualError(t, err, tt.errStr)
			}
		})
	}
}

func TestPostCreate(t *testing.T) {
	t.Parallel()

	type args struct {
		input domain.CreatePostInput
	}
	tests := map[string]struct {
		args    args
		want    *domain.Post
		wantErr bool
		errStr  string
	}{
		"create post": {
			args: args{
				input: domain.CreatePostInput{
					Title:    "create post title",
					Contents: "create post bio",
					UserID:   1,
					PhotoID:  11,
				},
			},
			want: &domain.Post{
				Title:    "create post title",
				Contents: "create post bio",
			},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client, teardown := utils.NewDBClient(t)
			defer teardown()

			r := NewPostRepository(client)
			got, err := r.Create(context.Background(), tt.args.input)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Title, got.Title)
				assert.Equal(t, tt.want.Contents, got.Contents)
			}

			if tt.wantErr {
				assert.EqualError(t, err, tt.errStr)
			}
		})
	}
}

func TestPostDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		id     int
		userID int
	}
	tests := map[string]struct {
		args    args
		wantErr bool
		errStr  string
	}{
		"delete post": {
			args: args{
				id:     1,
				userID: 1,
			},
		},
		"not found post": {
			args: args{
				id:     999,
				userID: 999,
			},
			wantErr: true,
			errStr:  "ent: post not found",
		},
		"cannot be post deleted": {
			args: args{
				id:     2,
				userID: 2,
			},
			wantErr: true,
			errStr:  "this post cannot be deleted",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client, teardown := utils.NewDBClient(t)
			defer teardown()

			r := NewPostRepository(client)
			err := r.Delete(context.Background(), tt.args.id, tt.args.userID)
			if !tt.wantErr {
				assert.NoError(t, err)
			}

			if tt.wantErr {
				assert.EqualError(t, err, tt.errStr)
			}
		})
	}
}
