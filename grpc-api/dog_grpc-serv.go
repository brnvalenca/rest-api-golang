package apiservice

import (
	"context"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/middleware"
	"rest-api/golang/exercise/proto/pb"
	"rest-api/golang/exercise/services"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: Analisar por que no retorno do create dog o ID do cachorro vem 0
// TODO: Na funcao de update, o retorno tem o campo Breed com todos os valores zerados, consertar.

type DogService struct {
	pb.UnimplementedDogServiceServer
	dogService   services.IDogService
	breedService services.IBreedService
}

func NewDogService(dogserv services.IDogService, breedServ services.IBreedService) *DogService {
	return &DogService{dogService: dogserv, breedService: breedServ}
}

func (dogserv *DogService) CreateDog(ctx context.Context, req *pb.CreateDogRequest) (*pb.Dog, error) {

	dogBuilder := dtos.NewDogDTOBuilder()
	dogBuilder.Has().
		BreedID(int(req.GetBreedID())).
		KennelID(int(req.GetKennelID())).
		NameAndSex(req.GetDogName(), req.GetSex())

	dogDto := dogBuilder.BuildDogDTO()
	breedCheck := dogserv.dogService.CheckIfBreedExist(dogDto)
	if !breedCheck {
		return nil, status.Errorf(codes.NotFound, "breed not found")
	}
	kennelCheck := dogserv.dogService.CheckIfKennelExist(dogDto)
	if !kennelCheck {
		return nil, status.Errorf(codes.NotFound, "kennel not found")
	}

	dog, breed := middleware.PartitionDogDTO(*dogDto)
	dogserv.dogService.CreateDog(dog, breed)

	result := &pb.Dog{
		KennelID: int32(dogDto.KennelID),
		DogID:    int32(dogDto.DogID),
		DogName:  dogDto.DogName,
		Sex:      dogDto.Sex,
		Breed: &pb.DogBreed{
			BreedID:   int32(breed.ID),
			BreedName: breed.Name,
		},
	}
	return result, nil
}

func (dogserv *DogService) GetAllDogs(ctx context.Context, req *pb.EmptyRequest) (*pb.GetAllDogsResponse, error) {
	dogsListDto, err := dogserv.dogService.FindDogs()

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "dogs not found")
	}

	result := pb.GetAllDogsResponse{}

	var dogsSlicePb []*pb.Dog
	for i := 0; i < len(dogsListDto); i++ {
		dogs := &pb.Dog{
			KennelID: int32(dogsListDto[i].KennelID),
			DogID:    int32(dogsListDto[i].DogID),
			DogName:  dogsListDto[i].DogName,
			Sex:      dogsListDto[i].Sex,
			Breed: &pb.DogBreed{
				BreedID:      int32(dogsListDto[i].BreedID),
				GoodWithKids: int32(dogsListDto[i].GoodWithKids),
				GoodWithDogs: int32(dogsListDto[i].GoodWithDogs),
				Shedding:     int32(dogsListDto[i].Shedding),
				Grooming:     int32(dogsListDto[i].Grooming),
				Energy:       int32(dogsListDto[i].Energy),
				BreedName:    dogsListDto[i].BreedName,
				BreedImg:     dogsListDto[i].BreedImg,
			},
		}
		dogsSlicePb = append(dogsSlicePb, dogs)
	}
	result.DogList = dogsSlicePb

	return &result, nil
}

func (dogserv *DogService) GetDogById(ctx context.Context, req *pb.DogID) (*pb.Dog, error) {

	dogDto, err := dogserv.dogService.FindDogByID(req.GetDogID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "dog not found")
	}

	result := &pb.Dog{
		KennelID: int32(dogDto.KennelID),
		DogID:    int32(dogDto.DogID),
		DogName:  dogDto.DogName,
		Sex:      dogDto.Sex,
		Breed: &pb.DogBreed{
			BreedID:      int32(dogDto.BreedID),
			GoodWithKids: int32(dogDto.GoodWithKids),
			GoodWithDogs: int32(dogDto.GoodWithDogs),
			Shedding:     int32(dogDto.Shedding),
			Grooming:     int32(dogDto.Grooming),
			Energy:       int32(dogDto.Energy),
			BreedName:    dogDto.BreedName,
			BreedImg:     dogDto.BreedImg,
		},
	}
	return result, nil
}

func (dogserv *DogService) DeleteDog(ctx context.Context, req *pb.DogID) (*pb.Dog, error) {
	dogDto, err := dogserv.dogService.DeleteDog(req.GetDogID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "dog not found")
	}

	result := &pb.Dog{
		KennelID: int32(dogDto.KennelID),
		DogID:    int32(dogDto.DogID),
		DogName:  dogDto.DogName,
		Sex:      dogDto.Sex,
		Breed: &pb.DogBreed{
			BreedID:      int32(dogDto.BreedID),
			GoodWithKids: int32(dogDto.GoodWithKids),
			GoodWithDogs: int32(dogDto.GoodWithDogs),
			Shedding:     int32(dogDto.Shedding),
			Grooming:     int32(dogDto.Grooming),
			Energy:       int32(dogDto.Energy),
			BreedName:    dogDto.BreedName,
			BreedImg:     dogDto.BreedImg,
		},
	}
	return result, nil
}

func (dogserv *DogService) UpdateDog(ctx context.Context, req *pb.UpdateDogRequest) (*pb.Dog, error) {

	check := dogserv.dogService.CheckIfDogExist(strconv.Itoa((int(req.GetDogID()))))
	if !check {
		return nil, status.Errorf(codes.NotFound, "dog not found")
	}

	dogBuilder := dtos.NewDogDTOBuilder()
	dogBuilder.Has().
		BreedID(int(req.GetBreedID())).
		KennelID(int(req.GetKennelID())).
		DogID(int(req.GetDogID())).
		NameAndSex(req.GetDogName(), req.GetSex())

	dogDto := dogBuilder.BuildDogDTO()

	// TODO: This check should recieve only the kennelID as argument, and be called in the beginning of the function to spare unecessary processing spent
	kennelCheck := dogserv.dogService.CheckIfKennelExist(dogDto)
	if !kennelCheck {
		return nil, status.Errorf(codes.NotFound, "kennel not found")
	}

	dogserv.dogService.UpdateDog(dogDto, strconv.Itoa(int(req.GetDogID())))

	result := &pb.Dog{
		KennelID: int32(dogDto.KennelID),
		DogID:    int32(dogDto.DogID),
		DogName:  dogDto.DogName,
		Sex:      dogDto.Sex,
		Breed: &pb.DogBreed{
			BreedID: int32(dogDto.BreedID),
		},
	}
	return result, nil

}
