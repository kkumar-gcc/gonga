package bootstrap

// import (
// 	"fmt"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// )

// func SetupApplication() *gin.Engine {
// 	// Load the configuration files
// 	viper.SetConfigName("app")
// 	viper.AddConfigPath(".")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatalf("Error reading config file, %s", err)
// 	}

// 	viper.SetConfigName("database")
// 	viper.AddConfigPath(".")
// 	err = viper.MergeInConfig()
// 	if err != nil {
// 		log.Fatalf("Error reading config file, %s", err)
// 	}

// 	viper.SetConfigName("auth")
// 	viper.AddConfigPath(".")
// 	err = viper.MergeInConfig()
// 	if err != nil {
// 		log.Fatalf("Error reading config file, %s", err)
// 	}

// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Set the router mode based on the configuration
// 	if viper.GetString("app.mode") == "debug" {
// 		gin.SetMode(gin.DebugMode)
// 	} else {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	// Use the database configuration to open a database connection
// 	dbConfig := viper.GetStringMapString("database")
// 	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		dbConfig["username"],
// 		dbConfig["password"],
// 		dbConfig["host"],
// 		dbConfig["port"],
// 		dbConfig["database"],
// 	)
// 	db, err := gorm.Open("mysql", dbConnString)
// 	if err != nil {
// 		log.Fatalf("Error connecting to database, %s", err)
// 	}

// 	// Migrate the database based on the configuration
// 	if viper.GetBool("database.auto_migrate") {
// 		log.Println("Migrating database...")
// 		db.AutoMigrate(&User{})
// 	}

// 	// Use the auth configuration to set up authentication middleware
// 	authConfig := viper.GetStringMapString("auth")
// 	authMiddleware := NewAuthMiddleware(authConfig["secret"])

// 	// Set up the routes
// 	api := router.Group("/api")
// 	api.Use(authMiddleware)
// 	api.GET("/users", ListUsers)
// 	api.POST("/users", CreateUser)
// 	api.GET("/users/:id", GetUser)
// 	api.PUT("/users/:id", UpdateUser)
// 	api.DELETE("/users/:id", DeleteUser)

// 	return router
// }
