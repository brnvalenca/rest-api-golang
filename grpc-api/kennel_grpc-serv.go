package apiservice

import (
	"context"
	"rest-api/golang/exercise/domain/dtos"
	"rest-api/golang/exercise/proto/pb"
	"rest-api/golang/exercise/services"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KennelService struct {
	pb.UnimplementedKennelServiceServer
	kennelService services.IKennelService
	dogService    services.IDogService
}

func NewKennelGrpcService(kennelService services.IKennelService, dogService services.IDogService) *KennelService {
	return &KennelService{kennelService: kennelService, dogService: dogService}
}

func (kennelserv *KennelService) GetAllKennels(ctx context.Context, req *pb.EmptyRequest) (*pb.GetAllKennelsResponse, error) {
	kennelListDto, err := kennelserv.kennelService.FindAllKennels()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no kennel found")
	}

	result := pb.GetAllKennelsResponse{}

	var kennelSlice []*pb.Kennel

	for i := 0; i < len(kennelListDto); i++ {
		kennel := &pb.Kennel{
			KennelID:      int32(kennelListDto[i].ID),
			ContactNumber: kennelListDto[i].ContactNumber,
			Name:          kennelListDto[i].Name,
			Address: &pb.Address{
				Numero: kennelListDto[i].Numero,
				Rua:    kennelListDto[i].Rua,
				Bairro: kennelListDto[i].Bairro,
				CEP:    kennelListDto[i].CEP,
				Cidade: kennelListDto[i].Cidade,
			},
		}
		kennelSlice = append(kennelSlice, kennel)
	}

	result.KennelList = kennelSlice

	return &result, nil
}

func (kennelserv *KennelService) GetKennelById(ctx context.Context, req *pb.KennelID) (*pb.GetKennelByIdResponse, error) {

	kennelDto, err := kennelserv.kennelService.FindKennelByIdServ(strconv.Itoa(int(req.GetKennelID())))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error finding kennel", err)
	}

	var dogsArrPb []*pb.DogsInTheKennel
	for i := 0; i < len(kennelDto.Dogs); i++ {
		dogpb := &pb.DogsInTheKennel{
			DogName:   kennelDto.Dogs[i].DogName,
			DogSex:    kennelDto.Dogs[i].Sex,
			BreedName: kennelDto.Dogs[i].BreedName,
			BreedID:   int32(kennelDto.Dogs[i].BreedID),
		}
		dogsArrPb = append(dogsArrPb, dogpb)
	}

	kennel := &pb.GetKennelByIdResponse{
		Kennel: &pb.Kennel{
			KennelID:      int32(kennelDto.ID),
			ContactNumber: kennelDto.ContactNumber,
			Name:          kennelDto.Name,
			Address: &pb.Address{
				Numero: kennelDto.Numero,
				Rua:    kennelDto.Rua,
				Bairro: kennelDto.Bairro,
				CEP:    kennelDto.CEP,
				Cidade: kennelDto.Cidade,
			},
		},
	}
	kennel.Dogs = dogsArrPb
	return kennel, nil
}

func (kennelserv *KennelService) CreateKennel(ctx context.Context, req *pb.CreateKennelRequest) (*pb.Kennel, error) {

	kennelBuilder := dtos.NewKennelBuilderDTO()
	kennelBuilder.Has().
		ContactNumber(req.GetContactNumber()).
		Name(req.GetName()).
		Numero(req.GetAddress().Numero).
		Rua(req.GetAddress().Rua).
		Bairro(req.GetAddress().Bairro).
		CEP(req.GetAddress().CEP).
		Cidade(req.GetAddress().Cidade)
	kennelDto := kennelBuilder.BuildKennel()

	id, err := kennelserv.kennelService.SaveKennel(kennelDto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save kennel: ", err)
	}

	result := &pb.Kennel{
		KennelID:      int32(id),
		ContactNumber: kennelDto.ContactNumber,
		Name:          kennelDto.Name,
		Address: &pb.Address{
			Numero: kennelDto.Numero,
			Rua:    kennelDto.Rua,
			Bairro: kennelDto.Bairro,
			CEP:    kennelDto.CEP,
			Cidade: kennelDto.Cidade,
		},
	}

	return result, nil
}

func (kennelserv *KennelService) DeleteKennel(ctx context.Context, req *pb.KennelID) (*pb.Kennel, error) {

	kennelDto, err := kennelserv.kennelService.DeleteKennelServ(strconv.Itoa(int(req.GetKennelID())))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "error deleting kennel: ", err)
	} else {
		kennel := &pb.Kennel{
			KennelID:      int32(kennelDto.ID),
			ContactNumber: kennelDto.ContactNumber,
			Name:          kennelDto.Name,
			Address: &pb.Address{
				Numero: kennelDto.Numero,
				Rua:    kennelDto.Rua,
				Bairro: kennelDto.Bairro,
				CEP:    kennelDto.CEP,
				Cidade: kennelDto.Cidade,
			},
		}
		return kennel, nil
	}
}

func (kennelserv *KennelService) UpdateKennel(ctx context.Context, req *pb.Kennel) (*pb.Kennel, error) {

	kennelBuilder := dtos.NewKennelBuilderDTO()
	kennelBuilder.Has().
		ID(int(req.GetKennelID())).
		ContactNumber(req.GetContactNumber()).
		Name(req.GetName()).
		Numero(req.GetAddress().Numero).
		Rua(req.GetAddress().Rua).
		Bairro(req.GetAddress().Bairro).
		CEP(req.GetAddress().CEP).
		Cidade(req.GetAddress().Cidade)
	kennelDto := kennelBuilder.BuildKennel()

	err := kennelserv.kennelService.UpdateKennelServ(kennelDto, strconv.Itoa(kennelDto.ID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating kennel: ", err)
	}

	result := &pb.Kennel{
		KennelID:      int32(kennelDto.ID),
		ContactNumber: kennelDto.ContactNumber,
		Name:          kennelDto.Name,
		Address: &pb.Address{
			Numero: kennelDto.Numero,
			Rua:    kennelDto.Rua,
			Bairro: kennelDto.Bairro,
			CEP:    kennelDto.CEP,
			Cidade: kennelDto.Cidade,
		},
	}

	return result, nil

}
