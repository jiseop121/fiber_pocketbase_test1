/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.listRule = "@request.auth.id != \"\"\n&& @request.auth.id = user \n&& (title ~ \"f%\" || content:length >= 14)\n&& @request.headers.authorization != \"\""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.listRule = "@request.auth.id != \"\"\n&& @request.auth.id = user \n&& (title ~ \"f%\" || content:length >= 14)\n&& @request.headers.Authorization != \"\""

  return dao.saveCollection(collection)
})
