package client

import (
	"log"
	"os"
)

func UploadFile(filePaths []string){
	for _,path := range filePaths {
		if checkFile(path){
			data,err := os.ReadFile(path)
			if err != nil{
				log.Fatalf("Can't read %s", path)
			}
			
		}
	} 
}

func checkFile(path string) bool{
	fileStats,err := os.Stat(path)
	if (os.IsNotExist(err)) {
		
		log.Printf("%s File does not exitst", path)
		return false

	}else if(err != nil){
		
		log.Fatal("INTERNAL ERROR")
	}
	
	log.Println(fileStats.Name())
	return true
}