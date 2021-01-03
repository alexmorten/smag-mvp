# Social Record

> Distributed scraping and analysis pipeline for a range of social media platforms

This is a remodelled version of [this](https://github.com/codeuniversity/smag-mvp)

## TODOS

- [x] use yaml files to configure services
- [x] use default values for all topics and group id
- [ ] Dockerfile:
  - [ ] use go 1.15
  - [ ] use experimental syntax
  - [ ] use mount type cache
- [ ] create kafka topics and partitions automatically
- [ ] use go migrate in go code to migrate db tables
- [ ] use tor for scraping (tricky, seems like login is necessary for when coming from a tor ip address :/ )
- [ ] remove cropping from image processing pipeline
- [ ] use targeted scraping instead of limitless loop
- [ ] dont download pictures if internal picture url is set
- [ ] errgroup in face detection worker
- [ ] develop admin panel


ideas:

- [ ] use wire for dependency injection ?
- [ ] compress pictures in storage?
- [ ] write kafka messages as protobuf?
- [ ] single error queue? or no error queue?


not possible right now
- [ ] use cockroach db as postgres replacement (change data capture is enterprise feature)
- [ ] use changestream handling from cockroach db
