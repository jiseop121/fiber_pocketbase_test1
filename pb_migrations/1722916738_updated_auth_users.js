/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("jc3c1ljdsabba4a")

  collection.viewRule = "@request.auth.email = email"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("jc3c1ljdsabba4a")

  collection.viewRule = null

  return dao.saveCollection(collection)
})
