package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/graph"
	"github.com/marcustut/fyp/backend/internal/adapter/handler"
	"github.com/marcustut/fyp/backend/internal/util/crypto/argon2"
	"github.com/marcustut/fyp/backend/internal/util/crypto/jwt"
)

func (r *mutationResolver) SignInWithUsername(ctx context.Context, input graph.SignInWithUsername) (*graph.UserWithAuth, error) {
	// get user with username
	u, err := r.controller.User.GetByUsername(ctx, input.Username)
	if err != nil {
		log.Println(err)
		return nil, handler.HandleError(ctx, err)
	}

	// compare given password and hash stored in db
	match, err := argon2.ComparePassword(input.Password, u.PasswordHash)
	// return err if password not match
	if !match || err != nil {
		return nil, handler.HandleError(ctx, fmt.Errorf("password incorrect"))
	}

	// generate new JWT claims
	jwt, err := jwt.NewJWTClaims(&jwt.NewJWTClaimsInput{
		ID:       string(u.ID),
		Username: u.Username,
		Email:    u.Email,
	})
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return &graph.UserWithAuth{AccessToken: jwt.Token, ExpiredAt: jwt.ExpiredAt, User: u}, nil
}

func (r *mutationResolver) SignInWithEmail(ctx context.Context, input graph.SignInWithEmail) (*graph.UserWithAuth, error) {
	// get user with email
	u, err := r.controller.User.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	// compare given password and hash stored in db
	match, err := argon2.ComparePassword(input.Password, u.PasswordHash)
	// return err if password not match
	if !match || err != nil {
		return nil, handler.HandleError(ctx, fmt.Errorf("password incorrect"))
	}

	// generate new JWT claims
	jwt, err := jwt.NewJWTClaims(&jwt.NewJWTClaimsInput{
		ID:       string(u.ID),
		Username: u.Username,
		Email:    u.Email,
	})
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return &graph.UserWithAuth{AccessToken: jwt.Token, ExpiredAt: jwt.ExpiredAt, User: u}, nil
}

func (r *mutationResolver) SignInWithGithub(ctx context.Context, githubAccessToken string) (*graph.UserWithAuth, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignUp(ctx context.Context, input ent.CreateUserInput) (*graph.UserWithAuth, error) {
	// generate argon2 hash
	hash, err := argon2.GeneratePassword(argon2.DefaultConfig, input.PasswordHash)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	input.PasswordHash = hash

	// store user with password_hash into db
	u, err := r.controller.User.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	// genereate new JWT claims
	jwt, err := jwt.NewJWTClaims(&jwt.NewJWTClaimsInput{
		ID:       string(u.ID),
		Username: u.Username,
		Email:    u.Email,
	})
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return &graph.UserWithAuth{AccessToken: jwt.Token, ExpiredAt: jwt.ExpiredAt, User: u}, nil
}

func (r *queryResolver) ValidateAccessToken(ctx context.Context, token string) (bool, error) {
	valid, err := jwt.ValidateJWTToken(token)
	if err != nil {
		return false, handler.HandleError(ctx, err)
	}
	return valid, nil
}
