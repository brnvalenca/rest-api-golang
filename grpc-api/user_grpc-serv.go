package apiservice

import (
	"context"
	"log"
	"rest-api/golang/exercise/authentication"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/proto/pb"
	"rest-api/golang/exercise/security"
	"rest-api/golang/exercise/services"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userService     services.IUserService
	passwordService security.IPasswordHash
}

// TODO: No retorno da deleção de um usuário, o campo UserPrefs tá retornando null
// TODO: ID sendo passada desnecessariamente na funcao de update

func NewUserGrpcService(userService services.IUserService, passwordService security.IPasswordHash) *UserService {
	return &UserService{userService: userService, passwordService: passwordService}
}

func (userv *UserService) SignIn(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	check, userDB := userv.userService.CheckEmailServ(req.GetEmail())
	if !check {
		return nil, status.Errorf(codes.Unauthenticated, "email not registered")
	}

	checkPassword := userv.passwordService.CheckPassword(req.GetPassword(), userDB.PasswordDTO)
	if !checkPassword {
		return nil, status.Errorf(codes.Unauthenticated, "password incorrect")
	}

	token, err := authentication.GenerateJWT(userDB.ID)
	if err != nil {
		log.Println("internal error during token generation: ", err)
		return nil, status.Errorf(codes.Unauthenticated, "internal error generating jwt token: ", err)
	}
	resp := &pb.LoginResponse{
		Token: token,
	}
	return resp, nil
}

func (userv *UserService) SignUp(ctx context.Context, req *pb.User) (*pb.UserID, error) {
	check, _ := userv.userService.CheckEmailServ(req.GetEmail())
	if check {
		return nil, status.Errorf(codes.AlreadyExists, "email already registered on our system")
	}
	hashedPassword, err := userv.passwordService.GeneratePasswordHash(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash the password")
	}

	userPrefs := dtos.UserPrefsDTO{
		GoodWithKids: int(req.GetUserPrefs().GoodWithKids),
		GoodWithDogs: int(req.GetUserPrefs().GoodWithDogs),
		Shedding:     int(req.GetUserPrefs().Shedding),
		Grooming:     int(req.GetUserPrefs().Grooming),
		Energy:       int(req.GetUserPrefs().Energy),
	}

	userDTO := dtos.UserDTOSignUp{
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Password:  hashedPassword,
		UserPrefs: userPrefs,
	}

	userID, err := userv.userService.Create(&userDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user")
	}

	IDResponse := &pb.UserID{
		UserID: int32(userID),
	}

	return IDResponse, nil
}

func (userv *UserService) GetAllUsers(ctx context.Context, req *pb.GetEmptyRequest) (*pb.GetAllUsersResponse, error) {
	userDTO, err := userv.userService.FindAll()

	result := pb.GetAllUsersResponse{}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all users")
	}

	var usersWithoutPasswordSlice []*pb.UserWithoutPassword
	for i := 0; i < len(userDTO); i++ {
		userWithoutPassword := &pb.UserWithoutPassword{
			ID:    int32(userDTO[i].ID),
			Name:  userDTO[i].Name,
			Email: userDTO[i].Email,
			UserPrefs: &pb.UserPrefs{
				GoodWithKids: int32(userDTO[i].UserPrefs.GoodWithKids),
				GoodWithDogs: int32(userDTO[i].UserPrefs.GoodWithDogs),
				Shedding:     int32(userDTO[i].UserPrefs.Shedding),
				Grooming:     int32(userDTO[i].UserPrefs.Grooming),
				Energy:       int32(userDTO[i].UserPrefs.Energy),
			},
		}
		usersWithoutPasswordSlice = append(usersWithoutPasswordSlice, userWithoutPassword)
	}

	result.UsersList = usersWithoutPasswordSlice

	return &result, nil
}

func (userv *UserService) GetUserById(ctx context.Context, req *pb.UserID) (*pb.UserWithoutPassword, error) {
	id := req.GetUserID()
	idInt := int(id)
	idString := strconv.Itoa(idInt)
	userDTO, err := userv.userService.FindById(idString)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find user")
	}
	response := &pb.UserWithoutPassword{
		ID:    int32(userDTO.ID),
		Name:  userDTO.Name,
		Email: userDTO.Email,
		UserPrefs: &pb.UserPrefs{
			GoodWithKids: int32(userDTO.UserPrefs.GoodWithKids),
			GoodWithDogs: int32(userDTO.UserPrefs.GoodWithDogs),
			Shedding:     int32(userDTO.UserPrefs.Shedding),
			Grooming:     int32(userDTO.UserPrefs.Grooming),
			Energy:       int32(userDTO.UserPrefs.Energy),
		},
	}

	return response, nil
}

func (userv *UserService) DeleteUser(ctx context.Context, req *pb.UserID) (*pb.UserWithoutPassword, error) {
	idString := strconv.Itoa(int(req.GetUserID()))
	check := userv.userService.Check(idString)

	if !check {
		return nil, status.Errorf(codes.NotFound, "failed to find user")
	} else {
		user, err := userv.userService.Delete(idString)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to delete user")
		} else {
			response := &pb.UserWithoutPassword{
				ID:    int32(user.ID),
				Name:  user.Name,
				Email: user.Email,
			}
			return response, nil
		}
	}
}

func (userv *UserService) UpdateUser(ctx context.Context, req *pb.UserWithoutPassword) (*pb.UserWithoutPassword, error) {

	idString := strconv.Itoa(int(req.GetID()))

	check := userv.userService.Check(idString)
	if !check {
		return nil, status.Errorf(codes.NotFound, "failed to find user")
	}
	userPrefs := dtos.UserPrefsDTO{
		UserID:       int(req.GetID()),
		GoodWithKids: int(req.UserPrefs.GetGoodWithKids()),
		GoodWithDogs: int(req.UserPrefs.GetGoodWithDogs()),
		Shedding:     int(req.UserPrefs.GetShedding()),
		Grooming:     int(req.UserPrefs.GetGrooming()),
		Energy:       int(req.UserPrefs.GetEnergy()),
	}
	userDTO := dtos.UserDTOSignUp{
		ID:        int(req.GetID()),
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		UserPrefs: userPrefs,
	}

	err := userv.userService.UpdateUser(&userDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user", err)
	}
	result := &pb.UserWithoutPassword{
		ID:    int32(userDTO.ID),
		Name:  userDTO.Name,
		Email: userDTO.Email,
		UserPrefs: &pb.UserPrefs{
			GoodWithKids: int32(userDTO.UserPrefs.GoodWithKids),
			GoodWithDogs: int32(userDTO.UserPrefs.GoodWithDogs),
			Shedding:     int32(userDTO.UserPrefs.Shedding),
			Grooming:     int32(userDTO.UserPrefs.Grooming),
			Energy:       int32(userDTO.UserPrefs.Energy),
		},
	}
	return result, nil

}
