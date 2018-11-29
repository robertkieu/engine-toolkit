# task-rt-test-engine-nodejs
This is test chunk engine using Nodejs.

### Summary

This engine is test chunk engine using Nodejs that belongs to the face detection category and the output is fixed.   


### Installation

`
    npm install
`
### Depedencies

* node v9.11.2

### Run local 
`
    node app.js
`
#### Sample kafka message with local test
``
{
"type":"media_chunk",
"mimeType":"video/mp4",
"cacheURI":"https://s3.amazonaws.com/test-chunk-engine/chunk2.mp4",
"TaskID": "taskid",
"ChunkUUID": "chunkuuid"
}
``

### Environment variables
Variables that can be passed in as environment variables i.e. `docker run -e KAFKA_CHUNK_TOPIC="CHUNK_ALL"`

| Variable              | Description                                            |
|-----------------------|--------------------------------------------------------|
| KAFKA_BROKERS         | Comma-seperated list of Kafka Broker addresses.        |
| KAFKA_CHUNK_TOPIC     | The Chunk Queue Kafka topic. Ex: "chunk_all"           |
| ENGINE_ID             | The engine ID                                          |
| ENGINE_INSTANCE_ID    | Unique instance ID for the engine instance             |
| KAFKA_INPUT_TOPIC     | The Kafka topic the engine should consume chunks from. |
| KAFKA_CONSUMER_GROUP  | The consumer group the engine must use.                | 
