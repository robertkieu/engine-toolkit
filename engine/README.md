# `engine` tool

This tool is bundled into engines and becomes the entry point. It is more fully explained in [the documentation](https://machinebox.io/experiments/engine-toolkit#the-engine-executable).

## Development

### Kafka integration test

Start Kafka with:

```
zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties & kafka-server-start /usr/local/etc/kafka/server.properties
```

Run the tests with:

```
KAFKATEST=true go test -v ./...
```

## TODO

- [ ] Strict response checking?
- [ ] Kick off a test job?
- [ ] CI 

## Resources

* https://steel-ventures.atlassian.net/wiki/spaces/VT/pages/522453767/Message+Topics+Formats+and+Schema
* https://steel-ventures.atlassian.net/wiki/spaces/VT/pages/718700686/Edge+Best+Practices+Common+Troubleshootings
