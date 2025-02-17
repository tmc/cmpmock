package _example_test

import (
	"context"
	"testing"
	"time"

	"github.com/tmc/cmpmock"
	"github.com/tmc/cmpmock/_example"
	"github.com/tmc/cmpmock/_example/mock_example"
	"go.uber.org/mock/gomock"
)

func TestUserUsecase_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)
	ctx := context.Background()
	name := "John Due"
	address := "Tokyo"
	wantUser := &_example.User{
		Name:     name,
		Address:  address,
		CreateAt: time.Now(),
	}

	mrepo := mock_example.NewMockUserRepo(ctrl)
	mrepo.EXPECT().Save(ctx, cmpmock.DiffEq(wantUser)).Return(nil)

	sut := _example.UserUsecase{Repo: mrepo}

	if err := sut.Register(ctx, name, address); err != nil {
		t.Fatal(err)
	}
}
