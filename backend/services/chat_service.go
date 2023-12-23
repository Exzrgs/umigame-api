package services

import (
	"umigame-api/models"
	"umigame-api/repositories"
)

func (s *Service) GetChatroomListService() ([]models.Chatroom, error) {
	chatrooms, err := repositories.SelectChatroomList(s.db, s.userID)
	if err != nil {
		return nil, err
	}

	return chatrooms, nil
}
