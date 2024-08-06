/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("jc3c1ljdsabba4a")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "cnueivza",
    "name": "asldkfj",
    "type": "text",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("jc3c1ljdsabba4a")

  // remove
  collection.schema.removeField("cnueivza")

  return dao.saveCollection(collection)
})
