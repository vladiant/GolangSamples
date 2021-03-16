# Tamper-Proof Systems sample

Based on [Building Tamper-Proof Systems with ImmuDB and Go!](https://www.youtube.com/watch?v=01_wEMdOp1U) from [TutorialEdge](https://www.youtube.com/channel/UCwFl9Y49sWChrddQTD9QhRA)

## Package init
```
go mod init github.com/vladiant/GolangSamples/BuildingTamperProofSystems
```

## immudb site 

https://github.com/codenotary/immudb

### Run
```
docker pull codenotary/immudb:latest
docker run -it -d -p 3322:3322 -p 9497:9497 --name immudb codenotary/immudb:latest
```

### Check
```
docker logs immudb
```
