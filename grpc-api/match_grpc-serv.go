package apiservice

import (
	"context"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/proto/pb"
	"rest-api/golang/exercise/services"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MatchService struct {
	pb.UnimplementedMatchServiceServer
	userService services.IUserService
	dogService  services.IDogService
}

func NewMatchGrpcService(UserService services.IUserService, DogService services.IDogService) *MatchService {
	return &MatchService{userService: UserService, dogService: DogService}
}

func getUserPrefs(userDto *dtos.UserDTOSignUp) *dtos.UserPrefsDTO {
	userPrefBuilder := dtos.NewUserPrefsDTOBuilder()
	userPrefBuilder.Has().
		UserID(userDto.ID).
		GoodWithKids(userDto.UserPrefs.GoodWithKids).
		GoodWithDogs(userDto.UserPrefs.GoodWithDogs).
		SheddAndGroom(userDto.UserPrefs.Shedding, userDto.UserPrefs.Grooming).
		Energy(userDto.UserPrefs.Energy)

	userPrefs := userPrefBuilder.BuildUserPrefsDTO()

	return userPrefs
}

func (matchserv *MatchService) MatchUserWithDog(ctx context.Context, req *pb.UserID) (*pb.Dog, error) {

	userDto, err := matchserv.userService.FindById(strconv.Itoa(int(req.GetUserID())))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %w", err)
	}

	dogsDtoList, err := matchserv.dogService.FindDogs()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "dogs not found: %w", err)
	}

	userPrefs := getUserPrefs(userDto)

	tempScore := 0
	score := 0
	idx := 0

	for i := 0; i < len(dogsDtoList); i++ {
		if userPrefs.GoodWithKids == dogsDtoList[i].GoodWithKids {
			tempScore++
		}
		if userPrefs.GoodWithDogs == dogsDtoList[i].GoodWithDogs {
			tempScore++
		}
		if userPrefs.Energy == dogsDtoList[i].Energy {
			tempScore++
		}
		if userPrefs.Grooming == dogsDtoList[i].Grooming {
			tempScore++
		}
		if userPrefs.Shedding == dogsDtoList[i].Grooming {
			tempScore++
		}

		if tempScore > score {
			score = tempScore
			idx = i
		}
		tempScore = 0
	}

	dogMatched := dogsDtoList[idx]
	resp := &pb.Dog{
		KennelID: int32(dogMatched.KennelID),
		DogID:    int32(dogMatched.DogID),
		DogName:  dogMatched.DogName,
		Sex:      dogMatched.Sex,
		Breed: &pb.DogBreed{
			BreedID:      int32(dogMatched.BreedID),
			GoodWithKids: int32(dogMatched.GoodWithKids),
			GoodWithDogs: int32(dogMatched.GoodWithDogs),
			Shedding:     int32(dogMatched.Shedding),
			Grooming:     int32(dogMatched.Grooming),
			Energy:       int32(dogMatched.Energy),
			BreedName:    dogMatched.BreedName,
			BreedImg:     dogMatched.BreedImg,
		},
	}

	return resp, nil
}
