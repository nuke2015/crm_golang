	// "gopkg.in/mgo.v2/bson"
	// ."crm/models"


	var m Mongobase

	// //添加数据;
	// person := Person{}
	// person.Name="hello person"
	// err:=m.Use("person").Insert(person)
	// this.dump(err);
	
	// //条件查询;
	//  mid := m.MongoId("54b11de847b5b00b9751fff0")
	//  where:=bson.M{"_id":mid}
	//  result:=m.Use("person").All(where)
	//  this.dump(result)

	// //查询全部;
	//  result:=m.Use("person").All(nil)
	//  this.dump(result)

	// //查询单条;
	//  result:=m.Use("person").One(nil)
	//  this.dump(result)

	// //查询条数;
	// count,_ :=m.Use("person").Count(nil)
	// this.dump(count)

	// //skip1,limit3跳过;
	//  result:=m.Use("person").Query(nil,1,3);
	//  this.dump(result)	

	// //skip1,limit3跳过;
	//  result:=m.Use("person").Query(nil,3,1);
	//  this.dump(result)

	// //条件查询;
	//  result:=m.Use("person").Query(bson.M{"title":"hello"},3,1);
	//  this.dump(result)		

	//条件查询;
	 result:=m.Use("person").Query(nil,0,10);
	 this.dump(result)	
