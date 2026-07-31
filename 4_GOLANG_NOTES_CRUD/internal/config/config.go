package config

import(
	"fmt"
	"os"

	"github.com/joho/godotenv"
) 

type Config struct {
	MongoURI 	string
	MongoDB		string
	ServerPort	string
}

func Load() (Config, error){

	// gotdotenv.Load reads the .env and sets them into the process env
	// os.getenv -> will reads those values
	
	if err := godotenv.Load(); err != nil{
		return Config{}, fmt.Errorf("Failed to load .env")
	}

	mongoURI, err := extractEnv("MONGO_URI")

	if err != nil{
		return Config{}, err
	}

	mongoDb, err := extractEnv("MONGO_DB_NAME")

	if err != nil{
		return Config{}, err
	}

	port, err := extractEnv("PORT")

	if err != nil{
		return Config{}, err
	}

	return Config{
		MongoURI: mongoURI,
		MongoDB: mongoDb,
		ServerPort: port,
	}, nil
	
}

func extractEnv(key string) (string, error){
	val := os.Getenv(key)

	if val == ""{
		return "", fmt.Errorf("missing req env")
	}

	return val, nil
}