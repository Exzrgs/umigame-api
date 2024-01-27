package services

import (
	"umigame-api/models"
	"umigame-api/repositories"
)

func (s *Service) ChangeLikedService(activity models.Activity) (models.ActivityLiked, error) {
	if err := repositories.ChangeLiked(s.db, s.userID, activity); err != nil {
		return models.ActivityLiked{}, err
	}

	var response models.ActivityLiked
	response.ProblemID, response.UserID, response.IsLiked = activity.ProblemID, s.userID, !activity.IsLiked

	return response, nil
}
