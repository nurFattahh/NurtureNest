package di

import (
	http "NurtureNest/pkg/api"
	"NurtureNest/pkg/api/handler"
	"NurtureNest/pkg/config"
	"NurtureNest/pkg/db"
	"NurtureNest/pkg/repository"
	"NurtureNest/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	tutorRepository := repository.NewTutorRepository(gormDB)
	tutorUseCase := usecase.NewTutorUseCase(tutorRepository)
	tutorHandler := handler.NewTutorHandler(tutorUseCase)

	chatBotRepository := repository.NewChatBotRepository(gormDB)
	chatBotUseCase := usecase.NewChatBotUseCase(chatBotRepository)
	chatBotHandler := handler.NewChatBotHandler(chatBotUseCase)

	goalTrackingRepository := repository.NewGoalTrackingRepository(gormDB)
	goalTrackingUseCase := usecase.NewGoalTrackingUseCase(goalTrackingRepository)
	goalTrackingHandler := handler.NewGoalTrackingHandler(goalTrackingUseCase)

	forumRepository := repository.NewForumRepository(gormDB)
	forumUseCase := usecase.NewForumUseCase(forumRepository)
	forumHandler := handler.NewForumHandler(forumUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, tutorHandler, chatBotHandler, goalTrackingHandler, forumHandler)

	return serverHTTP, nil
}
