exports = async function ({ query, headers, body }, response) {  
    if (!query || !query.user_id || !body){
      return {
        status: "bad request",
        response: "missing user_id and text fields"
      }
    }
    // 1. Get a data source client
    const mongodb = context.services.get("mongodb-atlas");
    // 2. Get a database & collection
    const db = mongodb.db("cms")
    const collection = db.collection("posts")
    // 3. Read and write data with MongoDB queries
    
    // var obj = JSON.parse(body);
    await collection.insertOne({
      user_id: query.user_id,
      created_at: Date.now(),
      updated_at: Date.now(),
      text: body.text(),
    })
    
    return {
      status: "ok",
      response: "inserted"
    }
  };