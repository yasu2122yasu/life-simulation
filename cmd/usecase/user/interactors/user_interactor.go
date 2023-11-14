package interactors

import (
	"github.com/yasu2122yasu/life-sim-go/cmd/domain/user/entities"
	"github.com/yasu2122yasu/life-sim-go/cmd/domain/user/ids"
	"github.com/yasu2122yasu/life-sim-go/cmd/domain/user/repositories"
	"github.com/yasu2122yasu/life-sim-go/cmd/domain/user/values"
	"github.com/yasu2122yasu/life-sim-go/cmd/usecase/user/boundaries"
	"github.com/yasu2122yasu/life-sim-go/cmd/usecase/user/requests"
	"github.com/yasu2122yasu/life-sim-go/cmd/usecase/user/responses"
)

type UserInteractor struct {
	userRepository repositories.UserRepository
}

var _ boundaries.UserInputBoundary = &UserInteractor{}

// NewUserInteractor コンストラクタ
func NewUserInteractor(userRepository repositories.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

// CreateUser ユーザを作成します．
func (ui *UserInteractor) CreateUser(userCreateRequest *requests.UserCreateRequest) (*responses.UserCreateResponse, error) {
	user := entities.NewUser(
		ids.UserId(userCreateRequest.UserId),
		values.NewUserName(userCreateRequest.UserLastName, userCreateRequest.UserFirstName, userCreateRequest.UserLastKanaName, userCreateRequest.UserFirstKanaName),
		values.UserGenderType(userCreateRequest.UserGenderType),
	)

	err := ui.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return &responses.UserCreateResponse{
		UserId:         user.Id().ToPrimitive(),
		UserName:       user.Name().FullName(),
		UserKanaName:   user.Name().FullKanaName(),
		UserGenderType: user.GenderType().String(),
	}, nil
}

// GetUser ユーザを取得します．
func (ui *UserInteractor) GetUser(userGetRequest *requests.UserGetRequest) (*responses.UserGetResponse, error) {
	user, err := ui.userRepository.FindById(ids.UserId(userGetRequest.UserId))

	if err != nil {
		return nil, err
	}

	return &responses.UserGetResponse{
		UserId:         user.Id().ToPrimitive(),
		UserName:       user.Name().FullName(),
		UserKanaName:   user.Name().FullKanaName(),
		UserGenderType: user.GenderType().String(),
	}, nil
}
