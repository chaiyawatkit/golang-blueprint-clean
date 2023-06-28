package users

import (
	"golang-blueprint-clean/app/entities"
	"golang-blueprint-clean/app/utils"
	"log"
)

func (u *useCase) LoginBySession(input *entities.UsersSignIn) (*string, error) {

	log.Printf(">>>%+v", input.SessionID)
	//sessionLogin := entities.UsersSignIn{
	//	SessionID: input.SessionID,
	//}
	//ตรงนี้คือขาที่ต่อกับทาง k+ ในที่นี้จะยกว่าได้ uid มาจาก k+

	//get data user
	//uid := uint(1)
	//filter := &entities.UsersFilter{
	//	Uid: &uid,
	//}
	//foundUser, err := u.UsersRepo.FindOneUser(filter)

	//mock user data
	foundUser := &entities.Users{
		Uuid:       1,
		UserStatus: "true",
		Segment:    "wisdom",
	}

	token, err := utils.GetJwtToken(foundUser)
	if err != nil {
		return nil, err
	}

	//ถ้ามีการเข้าสู้ระบบใหม่จะต้องไปอัพเดท token สำหรับ session ใหม่
	//foundUser.AccessToken = token
	//_, err = u.UsersRepo.UpdateUserAccessToken(foundUser)
	//if err != nil {
	//	return nil, errors.InternalError{Message: err.Error()}
	//}

	return token, nil
}
