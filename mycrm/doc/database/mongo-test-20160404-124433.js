
/** user indexes **/
db.getCollection("user").ensureIndex({
  "_id": NumberInt(1)
},[
  
]);

/** user records **/
db.getCollection("user").insert({
  "_id": ObjectId("5702616914104f580c00002a"),
  "username": "admin",
  "password": "admin"
});
