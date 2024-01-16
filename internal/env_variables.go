package internal

import "os"

func InitVariables() {
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("APP_PORT", "3000")
	os.Setenv("NGINX_PORT", "80")
	os.Setenv("ADDRESS", "localhost:"+os.Getenv("NGINX_PORT"))
	os.Setenv("LINK_EXP_TIME", "100") //minutes
}
