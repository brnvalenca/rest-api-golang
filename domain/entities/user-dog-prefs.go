package entities

type UserDogPreferences struct {
	UserID       int `json:"user_id"`
	GoodWithKids int `json:"good_with_kids"`
	GoodWithDogs int `json:"good_with_dogs"`
	Shedding     int `json:"shedding"`
	Grooming     int `json:"grooming"`
	Energy       int `json:"energy"`
}

type UserDogPrefsBuilder struct {
	userprefs *UserDogPreferences
}

type UserDogPrefsAttrBuilder struct {
	UserDogPrefsBuilder
}

func NewUserDogPrefsBuilder() *UserDogPrefsBuilder {
	return &UserDogPrefsBuilder{userprefs: &UserDogPreferences{}}
}

func (upref *UserDogPrefsBuilder) Has() *UserDogPrefsAttrBuilder {
	return &UserDogPrefsAttrBuilder{*upref}
}

func (upref *UserDogPrefsAttrBuilder) UserID(id int) *UserDogPrefsAttrBuilder {
	upref.userprefs.UserID = id
	return upref
}

func (upref *UserDogPrefsAttrBuilder) GoodWithKidsAndDogs(gwithkids, gwithdogs int) *UserDogPrefsAttrBuilder {
	upref.userprefs.GoodWithKids = gwithkids
	upref.userprefs.GoodWithDogs = gwithdogs
	return upref
}

func (upref *UserDogPrefsAttrBuilder) SheddGroomAndEnergy(shedding, grooming, energy int) *UserDogPrefsAttrBuilder {
	upref.userprefs.Shedding = shedding
	upref.userprefs.Grooming = grooming
	upref.userprefs.Energy = energy
	return upref
}

func (uprefbuilder *UserDogPrefsBuilder) BuildUserPref() *UserDogPreferences {
	return uprefbuilder.userprefs
}

func BuildUserDogPreferences(id, gdwithkids, gdwithdogs, shedd, groom, energy int) UserDogPreferences {

	udogpref := UserDogPreferences{
		UserID:       id,
		GoodWithKids: gdwithkids,
		GoodWithDogs: gdwithdogs,
		Shedding:     shedd,
		Grooming:     groom,
		Energy:       energy,
	}
	return udogpref
}
