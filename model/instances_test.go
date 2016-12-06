package model

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/axsh/openvdc/internal/unittest"
	"github.com/axsh/openvdc/model/backend"
	"github.com/stretchr/testify/assert"
)

func withConnect(t *testing.T, c func(context.Context)) error {

	ctx, err := Connect(context.Background(), []string{unittest.TestZkServer})
	if err != nil {
		t.Fatal(err)
	}
	defer Close(ctx)
	err = InstallSchemas(GetBackendCtx(ctx).(backend.ModelSchema))
	if err != nil {
		t.Fatal(err)
	}
	c(ctx)
	return err
}

func TestCreateInstance(t *testing.T) {
	assert := assert.New(t)
	n := &Instance{
		ExecutorId: "xxx",
	}

	var err error
	_, err = Instances(context.Background()).Create(n)
	assert.Equal(ErrInstanceMissingResource, err)
	n.ResourceId = "r-xxxx"
	_, err = Instances(context.Background()).Create(n)
	assert.Equal(ErrBackendNotInContext, err)

	withConnect(t, func(ctx context.Context) {
		got, err := Instances(ctx).Create(n)
		assert.NoError(err)
		assert.NotNil(got)
	})
}

func TestFindInstance(t *testing.T) {
	assert := assert.New(t)
	n := &Instance{
		ExecutorId: "xxx",
		ResourceId: "r-xxxx",
	}
	_, err := Instances(context.Background()).FindByID("i-xxxxx")
	assert.Equal(ErrBackendNotInContext, err)

	withConnect(t, func(ctx context.Context) {
		got, err := Instances(ctx).Create(n)
		assert.NoError(err)
		got2, err := Instances(ctx).FindByID(got.Id)
		assert.NoError(err)
		assert.NotNil(got2)
		assert.Equal(got.Id, got2.Id)
		_, err = Instances(ctx).FindByID("i-xxxxx")
		assert.Error(err)
	})
}

func TestUpdateStateInstance(t *testing.T) {
	assert := assert.New(t)
	err := Instances(context.Background()).UpdateState("i-xxxxx", InstanceState_REGISTERED)
	assert.Equal(ErrBackendNotInContext, err)

	withConnect(t, func(ctx context.Context) {
		n := &Instance{
			ExecutorId: "xxx",
			ResourceId: "r-xxxx",
		}
		got, err := Instances(ctx).Create(n)
		assert.NoError(err)
		assert.Equal(InstanceState_REGISTERED, got.GetLastState().State)
		assert.Error(Instances(ctx).UpdateState(got.GetId(), InstanceState_TERMINATED))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_QUEUED))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_STARTING))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_RUNNING))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_STOPPING))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_STOPPED))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_STARTING))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_RUNNING))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_SHUTTINGDOWN))
		assert.NoError(Instances(ctx).UpdateState(got.GetId(), InstanceState_TERMINATED))
	})
}

func TestUpdateInstance(t *testing.T) {
	assert := assert.New(t)
	err := Instances(context.Background()).Update(&Instance{Id: "i-xxxx"})
	assert.Equal(ErrBackendNotInContext, err)

	withConnect(t, func(ctx context.Context) {
		n := &Instance{
			ExecutorId: "xxx",
			ResourceId: "r-xxxx",
		}
		err = Instances(ctx).Update(n)
		assert.Error(err)
		assert.Equal(ErrInvalidID, err, "Empty ID object should get ErrInvalidID")

		got, err := Instances(ctx).Create(n)
		assert.NoError(err)
		got.ResourceId = "r-yyyy"
		err = Instances(ctx).Update(got)
		got2, err := Instances(ctx).FindByID(got.Id)
		assert.NoError(err)
		assert.Equal(got.ResourceId, got2.ResourceId)
	})
}
