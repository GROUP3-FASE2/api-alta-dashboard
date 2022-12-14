package repository

import (
	_log "api-alta-dashboard/features/log"
	"api-alta-dashboard/features/user"
	_user "api-alta-dashboard/features/user/repository"

	"gorm.io/gorm"
)

// struct gorm model
type Log struct {
	gorm.Model
	Title    string
	Feedback string
	Status   string
	UserID   uint
	MenteeID uint
	User     _user.User
}

// DTO
// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _log.Core) Log {
	logGorm := Log{
		Title:    dataCore.Title,
		Feedback: dataCore.Feedback,
		Status:   dataCore.Status,
		UserID:   dataCore.UserID,
		MenteeID: dataCore.MenteeID,
	}
	return logGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Log) toCore() _log.Core {
	return _log.Core{
		ID:        dataModel.ID,
		Title:     dataModel.Title,
		Feedback:  dataModel.Feedback,
		CreatedAt: dataModel.CreatedAt,
		Status:    dataModel.Status,
		UserID:    dataModel.UserID,
		User: user.Core{
			FullName: dataModel.User.FullName,
			Email:    dataModel.User.Email,
			Team:     dataModel.User.Team,
			Role:     dataModel.User.Role,
			Status:   dataModel.User.Status,
		},
		MenteeID: dataModel.MenteeID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Log) []_log.Core {
	var dataCore []_log.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
