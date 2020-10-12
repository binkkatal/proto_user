package store

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/binkkatal/proto_user/src/model"
	"github.com/binkkatal/proto_user/src/service"
	pb "github.com/golang/protobuf/ptypes/timestamp"
)

type ActivityStore struct {
	service.UnimplementedActivityServiceServer
	db *gorm.DB
}

func NewActivityStore(db *gorm.DB) *ActivityStore {
	return &ActivityStore{db: db}
}

func (st *ActivityStore) Add(ctx context.Context, in *model.Activity) (*service.ActivityResponse, error) {
	in.Status = model.Activity_CREATED
	err := st.db.Create(in).Error
	if err != nil {
		return nil, err
	}
	return &service.ActivityResponse{Activity: in}, nil
}

func (st *ActivityStore) Get(ctx context.Context, in *service.GetActivityRequest) (*service.ActivityResponse, error) {

	activity, err := st.get(in.Id)
	return &service.ActivityResponse{Activity: activity}, err
}
func (st *ActivityStore) Start(ctx context.Context, in *service.GetActivityRequest) (*service.ActivityResponse, error) {
	activity, err := st.get(in.Id)
	if err != nil {
		return nil, err
	}
	ts := time.Now()
	err = st.db.Model(activity).Updates(model.Activity{Status: model.Activity_IN_PROGRESS, StartedAt: ts.Unix()}).Error
	return &service.ActivityResponse{Activity: activity}, err
}
func (st *ActivityStore) Stop(ctx context.Context, in *service.GetActivityRequest) (*service.ActivityResponse, error) {
	activity, err := st.get(in.Id)
	if err != nil {
		return nil, err
	}
	ts := time.Now()
	err = st.db.Model(activity).Updates(model.Activity{Status: model.Activity_DONE, FinishedAt: ts.Unix()}).Error
	return &service.ActivityResponse{Activity: activity}, err
}
func (st *ActivityStore) IsValid(ctx context.Context, in *service.GetActivityRequest) (*service.BoolResult, error) {
	activity, err := st.get(in.Id)
	if err != nil {
		return nil, err
	}
	result := &service.BoolResult{}
	if activity.Status != model.Activity_DONE {
		return result, err
	}
	startedAt := time.Unix(activity.StartedAt, 0)
	finishedAt := time.Unix(activity.FinishedAt, 0)

	switch activity.ActivityType {
	case model.Activity_EAT:
		result.Result = startedAt.Add(1*time.Hour).Sub(finishedAt) > 0
	case model.Activity_SLEEP:
		result.Result = startedAt.Add(6*time.Hour).Sub(finishedAt) > 0
	case model.Activity_PLAY:
		result.Result = startedAt.Add(2*time.Hour).Sub(finishedAt) > 0
	case model.Activity_READ:
		result.Result = startedAt.Add(3*time.Hour).Sub(finishedAt) > 0
	default:
		result.Result = false

	}
	result.Result = activity.Status == model.Activity_DONE
	return result, nil
}
func (st *ActivityStore) IsDone(ctx context.Context, in *service.GetActivityRequest) (*service.BoolResult, error) {
	activity, err := st.get(in.Id)
	if err != nil {
		return nil, err
	}
	result := &service.BoolResult{}
	result.Result = activity.Status == model.Activity_DONE
	return result, nil
}

func (st *ActivityStore) get(uid int32) (*model.Activity, error) {
	activity := &model.Activity{Id: uid}
	err := st.db.First(activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func newTimeStamp() *pb.Timestamp {
	now := time.Now()
	s := int64(now.Second())
	n := int32(now.Nanosecond())

	ts := &pb.Timestamp{Seconds: s, Nanos: n}
	return ts
}
