/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("jc3c1ljdsabba4a")

  collection.name = "superusers"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("jc3c1ljdsabba4a")

  collection.name = "auth_users"

  return dao.saveCollection(collection)
})
