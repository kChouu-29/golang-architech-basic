package main

import (
	"database/sql"
	"log"
	"mvc/internal/controller"
	"mvc/internal/database"
	"mvc/internal/handler"
	"mvc/internal/repository"
	"mvc/internal/server" // Import package 'server' bạn vừa tạo
	"mvc/internal/service"
	"net/http"
)

func main() {
	// 1. Gọi "nhà máy" để lấy router đã được lắp ráp hoàn chỉnh
	dbConn := database.New()
	sqlDB, ok := dbConn.(*sql.DB)
	if !ok {
		log.Fatal("Failed to assert dbConn to *sql.DB")
	}
	// 2. Khởi tạo Repository (Inject DB)
	userRepo := repository.NewUserRepo(sqlDB)

	// 3. Khởi tạo Service (Inject Repo)
	userService := service.NewUserService(userRepo)

	// 4. Khởi tạo Controller (Inject Service)
	// (Dùng *struct* vì bạn bỏ mock)
	userController := controller.NewUserController(userService)

	// 5. Khởi tạo Handler (Inject Controller)
	userHandler := handler.NewUserHandler(userController)

	log.Println("Setting up application routes...")
	mux := server.NewRouter(userHandler)

	// 2. Cấu hình server
	log.Println("Starting server on http://localhost:8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux, // Gán router của bạn cho server
	}

	// 3. Chạy server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
