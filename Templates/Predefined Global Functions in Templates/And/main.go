package main

import (
	"log"
	"os"
	"text/template"
)

type user struct{
	Name string
	Motto string
	Admin bool
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main()  {
	u1 := user{
		Name : "Buddha",
		Motto : "Be Enlightened",
		Admin : false,	
	}	
		
	u2 := user{
		Name : "Gandhi",
		Motto : "Be Straightforward",
		Admin : true,	
	}	
		
	u3 := user{
		Name : "",
		Motto : "Nobody",
		Admin : true,	
	}	

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil{
		log.Fatalln(err)
	}		
}