## php demo
```
$uri = 'mongodb://172.16.100.94:27017/';
$collection = (new \MongoDB\Client($uri))->test->users;

/*
// example1 (inserts a document into the users collection in the test database)
// The MongoDB\Collection::insertOne() method inserts a single document into MongoDB and returns an instance of MongoDB\InsertOneResult, which you can use to access the ID of the inserted document.
$insertOneResult = $collection->insertOne([
'username' => 'admin',
'email' => 'admin@example.com',
'name' => 'Admin User',
]);

printf("Inserted %d document(s)\n", $insertOneResult->getInsertedCount());

var_dump($insertOneResult->getInsertedId());
*/


/*
// example2 (specifying the value for the _id)
// If you include an _id value when inserting a document, MongoDB checks to ensure that the _id value is unique for the collection. If the _id value is not unique, the insert operation fails due to a duplicate key error.
$insertOneResult = $collection->insertOne(['_id' => 1, 'name' => 'Alice']);

printf("Inserted %d document(s)\n", $insertOneResult->getInsertedCount());

var_dump($insertOneResult->getInsertedId());
*/


/*
// example3 (inserts two documents into the users collection in the test database)
// The MongoDB\Collection::insertMany() method allows you to insert multiple documents in one write operation and returns an instance of MongoDB\InsertManyResult, which you can use to access the IDs of the inserted documents.
$insertManyResult = $collection->insertMany([
[
'username' => 'admin',
'email' => 'admin@example.com',
'name' => 'Admin User',
],
[
'username' => 'test',
'email' => 'test@example.com',
'name' => 'Test User',
],
]);

printf("Inserted %d document(s)\n", $insertManyResult->getInsertedCount());

var_dump($insertManyResult->getInsertedIds());
*/


/*
// example4
// MongoDB\Collection::findOne() returns the first document that matches the query or null if no document matches the query.
$document = $collection->findOne(['_id' => 1]);
var_dump($document);

$mongoId = '606fb6164ac31f521735b122';
$collection->findOne(['_id'=> new \MongoDB\BSON\ObjectId($mongoId)]); // new \MongoId($mongoId) not work in php7
var_dump($document);

$document = $collection->findOne(['username' => 'admin']);
var_dump($document);
*/


/*
// example5
// MongoDB\Collection::find() returns a MongoDB\Driver\Cursor object, which you can iterate upon to access all matched documents.
$cursor = $collection->find(['username' => 'admin', 'name' => 'Admin User']);
foreach ($cursor as $document) {
// var_dump($document);
echo $document['_id'], "\n";
}
*/


/*
// example6 (projection)
// By default, queries in MongoDB return all fields in matching documents. To limit the amount of data that MongoDB sends to applications, you can include a projection document in the query operation.
$cursor = $collection->find(
// where
[
'username' => 'admin',
'name' => 'Admin User',
],
// result
[
'projection' => [
'name' => 1,
'email' => 1,
],
'limit' => 4,
'sort' => ['_id' => -1], // -1 降序
]
);

foreach($cursor as $restaurant) {
var_dump($restaurant);
};
*/


/*
// example7 (Regular Expressions)
// Filter criteria may include regular expressions, either by using the MongoDB\BSON\Regex class directory or the $regex operator.
$cursor = $collection->find([
'name' => new \MongoDB\BSON\Regex('^admin', 'i'), // 'name' => ['$regex' => '^admin', '$options' => 'i'],
'username' => 'admin',
]);

foreach ($cursor as $document) {
printf("%s: %s, %s\n", $document['_id'], $document['username'], $document['name']);
}
*/


/*
// example8 (Complex Queries with Aggregation)
// MongoDB’s Aggregation Framework allows you to issue complex queries that filter, transform, and group collection data. The MongoDB PHP Library’s MongoDB\Collection::aggregate() method returns a Traversable object, 
// which you can iterate upon to access the results of the aggregation operation. Refer to the MongoDB\Collection::aggregate() method’s behavior reference for more about the method’s output.
$cursor = $collection->aggregate([
['$group' => ['_id' => '$name', 'count' => ['$sum' => 1]]],
['$sort' => ['_id' => -1]],
['$limit' => 5],
]);

foreach ($cursor as $row) {
printf("%s has %s\n", $row['_id'], $row['count']);
}
*/


/*
// example9 (Update One Document)
// Use the MongoDB\Collection::updateOne() method to update a single document matching a filter. MongoDB\Collection::updateOne() returns a MongoDB\UpdateResult object, which you can use to access statistics about the update operation.
// Update methods have two required parameters: the query filter that identifies the document or documents to update, and an update document that specifies what updates to perform. The MongoDB\Collection::updateOne() reference describes each parameter in detail.
// The following example inserts two documents into an empty users collection in the test database using the MongoDB\Collection::insertOne() method, and then updates the documents where the value for the state field is "ny" to include a country field set to "us":
$collection->drop();
$collection->insertOne(['name' => 'Bob', 'state' => 'ny']);
$collection->insertOne(['name' => 'Alice', 'state' => 'ny']);
$updateResult = $collection->updateOne(
['state' => 'ny'],
['$set' => ['country' => 'us']]
);
printf("Matched %d document(s)\n", $updateResult->getMatchedCount());
printf("Modified %d document(s)\n", $updateResult->getModifiedCount());

$collection->drop();
$collection->insertOne(['name' => 'Bob', 'state' => 'ny']);
$updateResult = $collection->updateOne(
['name' => 'Bob'],
['$set' => ['state' => 'ny']]
);
printf("Matched %d document(s)\n", $updateResult->getMatchedCount());
printf("Modified %d document(s)\n", $updateResult->getModifiedCount());
*/


/*
// example10 (Update Many Documents)
// MongoDB\Collection::updateMany() updates one or more documents matching the filter criteria and returns a MongoDB\UpdateResult object, which you can use to access statistics about the update operation.
// Update methods have two required parameters: the query filter that identifies the document or documents to update, and an update document that specifies what updates to perform. The MongoDB\Collection::updateMany() reference describes each parameter in detail.
$collection->drop();
$collection->insertOne(['name' => 'Bob', 'state' => 'ny', 'country' => 'us']);
$collection->insertOne(['name' => 'Alice', 'state' => 'ny']);
$collection->insertOne(['name' => 'Sam', 'state' => 'ny']);
$updateResult = $collection->updateMany(
['state' => 'ny'],
['$set' => ['country' => 'us']]
);

printf("Matched %d document(s)\n", $updateResult->getMatchedCount());
printf("Modified %d document(s)\n", $updateResult->getModifiedCount());
*/


/*
// example11 (Replace Documents)
// Replacement operations are similar to update operations, but instead of updating a document to include new fields or new field values, a replacement operation replaces the entire document with a new document, but retains the original document’s _id value.
// The MongoDB\Collection::replaceOne() method replaces a single document that matches the filter criteria and returns an instance of MongoDB\UpdateResult, which you can use to access statistics about the replacement operation.
// MongoDB\Collection::replaceOne() has two required parameters: the query filter that identifies the document or documents to replace, and a replacement document that will replace the original document in MongoDB. The MongoDB\Collection::replaceOne() reference describes each parameter in detail.
$collection->drop();
$collection->insertOne(['name' => 'Bob', 'state' => 'ny']);
$updateResult = $collection->replaceOne(
['name' => 'Bob'],
['name' => 'Robert', 'state' => 'ca']
);

printf("Matched %d document(s)\n", $updateResult->getMatchedCount());
printf("Modified %d document(s)\n", $updateResult->getModifiedCount());
*/


/*
// example12 (Upsert)
$collection->drop();
$updateResult = $collection->updateOne(
['name' => 'Bob'],
['$set' => ['state' => 'ny']],
['upsert' => true]
);

printf("Matched %d document(s)\n", $updateResult->getMatchedCount());
printf("Modified %d document(s)\n", $updateResult->getModifiedCount());
printf("Upserted %d document(s)\n", $updateResult->getUpsertedCount());

$upsertedDocument = $collection->findOne([
'_id' => $updateResult->getUpsertedId(),
]);

var_dump($upsertedDocument);
*/


/*
// example13 (Delete Documents)
// Delete One Document
$collection->drop();
$collection->insertOne(['name' => 'Bob', 'state' => 'ny']);
$collection->insertOne(['name' => 'Alice', 'state' => 'ny']);
$deleteResult = $collection->deleteOne(['state' => 'ny']);
printf("Deleted %d document(s)\n", $deleteResult->getDeletedCount());

// Delete Many Documents
$collection->drop();
$collection->insertOne(['name' => 'Bob', 'state' => 'ny']);
$collection->insertOne(['name' => 'Alice', 'state' => 'ny']);
$deleteResult = $collection->deleteMany(['state' => 'ny']);
printf("Deleted %d document(s)\n", $deleteResult->getDeletedCount());
*/
```
