/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("v9fglyvmal6p1bb")

  collection.listRule = "@request.auth.id = \"js@naver.com\""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("v9fglyvmal6p1bb")

  collection.listRule = "@request.auth.email ~ \"%js%\""

  return dao.saveCollection(collection)
})
