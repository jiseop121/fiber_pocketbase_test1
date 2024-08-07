/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.listRule = "@request.auth.id ?= user.id\n&& user.role = \"gold\""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.listRule = "@request.auth.id = user.id\n&& user.role = \"gold\""

  return dao.saveCollection(collection)
})
