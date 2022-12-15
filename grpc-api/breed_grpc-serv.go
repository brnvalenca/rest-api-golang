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

type BreedService struct {
	pb.UnimplementedBreedServiceServer
	breedService services.IBreedService
}

func NewBreedGrpcService(breedServ services.IBreedService) *BreedService {
	return &BreedService{breedService: breedServ}
}

func (breedserv *BreedService) CreateBreed(ctx context.Context, req *pb.CreateBreedRequest) (*pb.Breed, error) {

	breedBuilder := dtos.NewBreedBuilderDTO()
	breedBuilder.Has().
		Name(req.GetBreedName()).
		Img(req.GetBreedImg()).
		GoodWithKidsAndDogs(int(req.GetGoodWithKids()), int(req.GetGoodWithDogs())).
		SheddGroomAndEnergy(int(req.GetShedding()), int(req.GetGrooming()), int(req.GetEnergy()))

	breedDto := breedBuilder.BuildBreedDTO()

	err := breedserv.breedService.ValidateBreed(breedDto)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "breed not accepted: ", err)
	}
	err = breedserv.breedService.CreateBreed(breedDto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during breed creation: ", err)
	} else {
		result := &pb.Breed{
			BreedID:      int32(breedDto.ID),
			GoodWithKids: int32(breedDto.GoodWithKids),
			GoodWithDogs: int32(breedDto.GoodWithDogs),
			Shedding:     int32(breedDto.Shedding),
			Grooming:     int32(breedDto.Grooming),
			Energy:       int32(breedDto.Energy),
			BreedName:    breedDto.Name,
			BreedImg:     breedDto.BreedImg,
		}
		return result, nil
	}
}

func (breedserv *BreedService) GetAllBreed(ctx context.Context, req *pb.EmptyRequest) (*pb.GetAllBreedResponse, error) {

	breedListDto, err := breedserv.breedService.FindBreeds()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no breed founded", err)
	}

	result := pb.GetAllBreedResponse{}

	var breedSlice []*pb.Breed

	for i := 0; i < len(breedListDto); i++ {
		breed := &pb.Breed{
			BreedID:      int32(breedListDto[i].ID),
			GoodWithKids: int32(breedListDto[i].GoodWithKids),
			GoodWithDogs: int32(breedListDto[i].GoodWithDogs),
			Shedding:     int32(breedListDto[i].Shedding),
			Grooming:     int32(breedListDto[i].Grooming),
			Energy:       int32(breedListDto[i].Energy),
			BreedName:    breedListDto[i].Name,
			BreedImg:     breedListDto[i].BreedImg,
		}
		breedSlice = append(breedSlice, breed)
	}

	result.BreedList = breedSlice

	return &result, nil

}

func (breedserv *BreedService) GetBreedById(ctx context.Context, req *pb.BreedID) (*pb.Breed, error) {

	breedDto, err := breedserv.breedService.FindBreedByID(strconv.Itoa(int(req.GetBreedID())))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "breed not found", err)
	}

	result := &pb.Breed{
		BreedID:      int32(breedDto.ID),
		GoodWithKids: int32(breedDto.GoodWithKids),
		GoodWithDogs: int32(breedDto.GoodWithDogs),
		Shedding:     int32(breedDto.Shedding),
		Grooming:     int32(breedDto.Grooming),
		Energy:       int32(breedDto.Energy),
		BreedName:    breedDto.Name,
		BreedImg:     breedDto.BreedImg,
	}

	return result, nil
}

func (breedserv *BreedService) UpdateBreed(ctx context.Context, req *pb.Breed) (*pb.Breed, error) {
	breedBuilder := dtos.NewBreedBuilderDTO()
	breedBuilder.Has().
		ID(int(req.GetBreedID())).
		Name(req.GetBreedName()).
		Img(req.GetBreedImg()).
		GoodWithKidsAndDogs(int(req.GetGoodWithKids()), int(req.GetGoodWithDogs())).
		SheddGroomAndEnergy(int(req.GetShedding()), int(req.GetGrooming()), int(req.GetEnergy()))

	breedDto := breedBuilder.BuildBreedDTO()

	err := breedserv.breedService.ValidateBreed(breedDto)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "breed not accepted: ", err)
	}

	err = breedserv.breedService.UpdateBreed(breedDto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update breed: ", err)
	} else {
		result := &pb.Breed{
			BreedID:      int32(breedDto.ID),
			GoodWithKids: int32(breedDto.GoodWithKids),
			GoodWithDogs: int32(breedDto.GoodWithDogs),
			Shedding:     int32(breedDto.Shedding),
			Grooming:     int32(breedDto.Grooming),
			Energy:       int32(breedDto.Energy),
			BreedName:    breedDto.Name,
			BreedImg:     breedDto.BreedImg,
		}
		return result, nil
	}

}

func (breedserv *BreedService) DeleteBreed(ctx context.Context, req *pb.BreedID) (*pb.Breed, error) {
	check := breedserv.breedService.CheckIfBreedExist(strconv.Itoa(int(req.GetBreedID())))
	if !check {
		return nil, status.Errorf(codes.NotFound, "breed not found")
	}

	breedDto, err := breedserv.breedService.DeleteBreed(strconv.Itoa(int(req.GetBreedID())))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting breed: ", err)
	} else {
		result := &pb.Breed{
			BreedID:      int32(breedDto.ID),
			GoodWithKids: int32(breedDto.GoodWithKids),
			GoodWithDogs: int32(breedDto.GoodWithDogs),
			Shedding:     int32(breedDto.Shedding),
			Grooming:     int32(breedDto.Grooming),
			Energy:       int32(breedDto.Energy),
			BreedName:    breedDto.Name,
			BreedImg:     breedDto.BreedImg,
		}
		return result, nil
	}
}
