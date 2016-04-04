
sessionon = true




tmp:=this.GetSession("hello")
if(tmp!=nil){
	this.dump(tmp);
}else{
	this.SetSession("hello",time.Now())
}


	

this.SetSession("hello","world")
tmp:=this.GetSession("hello")
this.dump(tmp);
this.SetSession("hello","world2")
tmp2:=this.GetSession("hello")
this.dump(tmp2);
this.DelSession("hello");

