package main

import (
	"log"
	"net/http"
	"time"
	"todo/config"
	"todo/internal/auth"
	"todo/internal/lists"
	"todo/internal/tasks"
	"todo/internal/users"
	"todo/pkg/db"
)

func main() {
	conf := config.MustLoadConfig()

	router := http.NewServeMux()

	//TODO defer db.Close()
	database := db.NewPostgreDb(conf)
	defer database.Close()

	//Repository
	usersRep := users.NewUserRepository(database)
	listsRep := lists.NewListRepository(database)
	tasksRep := tasks.NewTaskRepository(database)

	//Service
	authService := auth.NewAuthService(usersRep)
	listsService := lists.NewListsService(listsRep)
	tasksService := tasks.NewTasksService(tasksRep)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	lists.NewListsHandler(router, lists.ListsHandlerDeps{
		Config:       conf,
		ListsService: listsService,
	})

	tasks.NewTaskHandlers(router, tasks.TaskHandlersDeps{
		Config:       conf,
		TasksService: tasksService,
	})

	server := &http.Server{
		Addr:           conf.App.HTTPAddr,
		Handler:        router,
		MaxHeaderBytes: 1 << conf.App.MaxHeaderByte, //1<<20
		ReadTimeout:    time.Duration(conf.App.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.App.WriteTimeout) * time.Second,
	}

	log.Printf("server started on %s", conf.App.HTTPAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("the server is not working: %s", err.Error())
	}

}
