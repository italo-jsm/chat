package config

type DatabaseConfiguration struct{
	Host string `json:"host"`
	Port int32 `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Dbname string `json:"dbname"`
}