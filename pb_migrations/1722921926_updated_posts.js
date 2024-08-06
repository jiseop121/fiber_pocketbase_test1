/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.viewRule = ""
  collection.createRule = "@request.auth.email = user.email"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.viewRule = "@request.auth.email = user.email"
  collection.createRule = ""

  return dao.saveCollection(collection)
})
