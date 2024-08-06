/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.createRule = "@request.data.user.email = user.email\n&& @request.data.user.role = \"gold\""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ulebprpm7fjc58b")

  collection.createRule = "@request.data.user.email = user.email"

  return dao.saveCollection(collection)
})
