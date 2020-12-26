# Social Record

> Distributed scraping and analysis pipeline for a range of social media platforms

This is a remodelled version of [this](https://github.com/codeuniversity/smag-mvp)

## TODOS

- [ ] use yaml files to configure services
- [ ] use default values for all topics and group id
- [ ] use kafka client to create all topics in code
- [ ] single error queue? or no error queue?
- [ ] create kafka topics and partitions automatically
- [ ] use go 1.15
- [ ] use tor for scraping
- [ ] use wire for dependency injection
- [ ] use cockroach db as postgres replacement
- [ ] use go migrate in go code to migrate db tables
- [ ] use changestream handling from cockroach db
- [ ] write kafka messages as protobuf?
- [ ] use targeted scraping instead of limitless loop
- [ ] dont download pictures if internal picture url is set
- [ ] compress pictures in storage?
- [ ] remove opencv from imageprocessing pipeline
- [ ] remove old frontend, develp admin panel
- [ ] errgroup in face detection worker
