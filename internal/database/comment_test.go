package database

import (
	"context"
	"github.com/Permitted/go-rest-api/internal/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		database, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := database.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := database.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		database, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := database.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		err = database.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = database.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})

	t.Run("test update comment", func(t *testing.T) {
		database, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := database.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		cmt.Slug = "new-slug"
		cmt, err = database.UpdateComment(context.Background(), cmt.ID, cmt)
		assert.NoError(t, err)

		newCmt, err := database.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "new-slug", newCmt.Slug)
	})

	t.Run("test get comment", func(t *testing.T) {
		database, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := database.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := database.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)

	})
}
