package main



import	(
	"storie/config"
	"storie/infrastructure/api"		
)

func init() {
	config.ConfigEnv()
}

func main() { 
	api.Start()

}