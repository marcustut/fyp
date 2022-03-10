package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/graph"
	"github.com/marcustut/fyp/backend/internal/adapter/handler"
	awsConstant "github.com/marcustut/fyp/backend/internal/const/aws"
	"github.com/marcustut/fyp/backend/internal/entity/model"
	"github.com/marcustut/fyp/backend/internal/util/helper"
)

func (r *mutationResolver) CreateSlide(ctx context.Context, input ent.CreateSlideInput) (*ent.Slide, error) {
	s, err := r.controller.Slide.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return s, nil
}

func (r *mutationResolver) CreateSlideWithText(ctx context.Context, input graph.CreateSlideWithTextInput, text string) (*ent.Slide, error) {
	// create slide in db
	s, err := r.CreateSlide(ctx, ent.CreateSlideInput{
		Name:        input.Name,
		AccessLevel: input.AccessLevel,
		SharedWith:  input.SharedWith,
		UserID:      input.UserID,
	})
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	// s3 path
	path := helper.ConstructPath("slides", string(input.UserID), string(s.ID)+".md")

	// upload the file to s3
	_, err = r.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(awsConstant.BucketName),
		Key:         aws.String(path),
		Body:        bytes.NewReader([]byte(text)),
		ContentType: aws.String("text/markdown"),
	})
	if err != nil {
		return nil, err
	}
	// get metadata of the file from s3
	o, err := r.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(awsConstant.BucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}

	// update the slide with data from s3
	us, err := r.UpdateSlide(ctx, ent.UpdateSlideInput{
		ID:        s.ID,
		PathToken: &[]string{"slides", string(input.UserID), string(s.ID) + ".md"},
		Size:      &o.ContentLength,
	})
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return us, nil
}

func (r *mutationResolver) UpdateSlide(ctx context.Context, input ent.UpdateSlideInput) (*ent.Slide, error) {
	s, err := r.controller.Slide.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return s, nil
}

func (r *mutationResolver) UpdateSlideWithText(ctx context.Context, id ulid.ID, text string) (*ent.Slide, error) {
	s, err := r.Query().Slide(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	userID, err := s.QueryUser().OnlyID(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	path := helper.ConstructPath("slides", string(userID), string(s.ID)+".md")

	// update the file in s3
	_, err = r.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(awsConstant.BucketName),
		Key:         aws.String(path),
		Body:        bytes.NewReader([]byte(text)),
		ContentType: aws.String("text/markdown"),
	})
	if err != nil {
		return nil, err
	}
	// get metadata of the file from s3
	o, err := r.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(awsConstant.BucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}

	// update the slide with data from s3
	us, err := r.UpdateSlide(ctx, ent.UpdateSlideInput{
		ID:        s.ID,
		PathToken: &[]string{"slides", string(userID), string(s.ID) + ".md"},
		Size:      &o.ContentLength,
	})
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return us, nil
}

func (r *mutationResolver) DeleteSlide(ctx context.Context, id ulid.ID, userID ulid.ID) (*ent.Slide, error) {
	path := helper.ConstructPath("slides", string(userID), string(id)+".md")

	// try to get file from s3
	_, err := r.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(awsConstant.BucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}

	// delete file from s3
	_, err = r.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(awsConstant.BucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}

	// delete slide from db
	s, err := r.controller.Slide.Delete(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return s, nil
}

func (r *mutationResolver) AddUsersToSlide(ctx context.Context, id ulid.ID, emails []string) (*ent.Slide, error) {
	s, err := r.controller.Slide.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	newSharedWith := helper.RemoveDuplicatesFromStrings(append(s.SharedWith, emails...))

	s, err = s.Update().SetSharedWith(newSharedWith).Save(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return s, nil
}

func (r *queryResolver) Slide(ctx context.Context, id ulid.ID) (*ent.Slide, error) {
	s, err := r.controller.Slide.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return s, nil
}

func (r *queryResolver) Slides(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.SlideWhereInput, orderBy *ent.SlideOrder) (*ent.SlideConnection, error) {
	ss, err := r.controller.Slide.List(ctx, after, first, before, last, where, orderBy)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return ss, nil
}
