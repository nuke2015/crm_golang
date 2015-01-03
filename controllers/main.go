package controllers

import(
	// "gopkg.in/mgo.v2/bson"
	."crm/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	var mg1,mg2 Mongo

	mg1.Use("person")
	person := Person{}
	person.Name="hello person"
	err:=mg1.Insert(person)
	this.dump(err);
	
	mg2.Use("people")
	people := People{}
	people.Name="hello people"
	err2:=mg2.Insert(people)
	this.dump(err2)
}
