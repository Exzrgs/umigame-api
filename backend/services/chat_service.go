package services

// import (
// 	"umigame-api/models"
// 	"umigame-api/myerrors"
// 	"umigame-api/repositories"

// 	"github.com/jinzhu/copier"
// )

// func (s *Service) GetChatRoomListService() ([]models.ChatList, error) {
// 	chats, err := repositories.SelectChatList(s.db, s.userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	titleMap, err := repositories.SelectProblemsByIDs(s.db)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var response []models.ChatList
// 	for _, chat := range chats {
// 		var resPart models.ChatList
// 		resPart.ProblemTitle, ok = titleMap[chat.ProblemID]
// 		if !ok {
// 			/*
// 				write error handling
// 			*/
// 		}
// 		if err := copier.Copy(&resPart, &chat); err != nil {
// 			err = myerrors.TypeCastFailed.Wrap(err, "internal server error")
// 			return models.ProblemDetail{}, err
// 		}
// 		response = append(response, resPart)
// 	}

// 	return response, nil
// }
