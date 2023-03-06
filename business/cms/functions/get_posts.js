exports = async function ({ query }) {  
    if (!query || !query.id){
      return {
        status: "bad request",
        response: "missing id field"
      }
    }
    // 1. Get a data source client
    const mongodb = context.services.get("mongodb-atlas");
    // 2. Get a database & collection
    const db = mongodb.db("cms")
    const collection = db.collection("posts")
    // 3. Read and write data with MongoDB queries
    return await collection.findOne({
      _id: BSON.ObjectId(query.id)
    })
  };