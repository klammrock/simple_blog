package db

type Config struct {
	Host          string `json:"host"`
	Port          int    `json:"port"`
	DbName        string `json:"db_name"`
	Collection    string `json:"collection"`
	LogCollection string `json:"log_collection"`
}
