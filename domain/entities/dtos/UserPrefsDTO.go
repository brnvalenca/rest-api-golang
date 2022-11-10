package dtos

type UserPrefsDTO struct {
	UserID       int `json:"user_id"`
	GoodWithKids int `json:"good_with_kids"`
	GoodWithDogs int `json:"good_with_dogs"`
	Shedding     int `json:"shedding"`
	Grooming     int `json:"grooming"`
	Energy       int `json:"energy"`
}

type UserPrefsDTOBuilder struct {
	uprefsdto *UserPrefsDTO
}

type UserPrefsDTOAttrBuilder struct {
	UserPrefsDTOBuilder
}

func NewUserPrefsDTOBuilder() *UserPrefsDTOBuilder {
	return &UserPrefsDTOBuilder{uprefsdto: &UserPrefsDTO{}}
}

func (updto *UserPrefsDTOBuilder) Has() *UserPrefsDTOAttrBuilder {
	return &UserPrefsDTOAttrBuilder{*updto}
}

func (updto *UserPrefsDTOAttrBuilder) UserID(userid int) *UserPrefsDTOAttrBuilder {
	updto.uprefsdto.UserID = userid
	return updto
}

func (updto *UserPrefsDTOAttrBuilder) GoodWithKids(gwithkids int) *UserPrefsDTOAttrBuilder {
	updto.uprefsdto.GoodWithKids = gwithkids
	return updto
}

func (updto *UserPrefsDTOAttrBuilder) GoodWithDogs(gwithdogs int) *UserPrefsDTOAttrBuilder {
	updto.uprefsdto.GoodWithDogs = gwithdogs
	return updto
}

func (updto *UserPrefsDTOAttrBuilder) SheddAndGroom(shedding, grooming int) *UserPrefsDTOAttrBuilder {
	updto.uprefsdto.Shedding = shedding
	updto.uprefsdto.Grooming = grooming
	return updto
}

func (updto *UserPrefsDTOAttrBuilder) Energy(energy int) *UserPrefsDTOAttrBuilder {
	updto.uprefsdto.Energy = energy
	return updto
}

func (updto *UserPrefsDTOBuilder) BuildUserPrefsDTO() *UserPrefsDTO {
	return updto.uprefsdto
}
