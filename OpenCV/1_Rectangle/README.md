## Build for Linux
```bash
docker build -t gocv .
```

## Run in Linux
```bash
xhost +
docker run -it -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix gocv
```

## Dev Run in Linux
```
xhost +
docker run -it --net=host --ipc=host -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix gocv /bin/bash
```

## VSCode run
* See the `.devcontainer/devcontainer.json`

## References
* <https://hub.docker.com/r/gocv/opencv>
* <https://marcosnietoblog.wordpress.com/2017/04/30/docker-image-with-opencv-with-x11-forwarding-for-gui/>
* <https://github.com/microsoft/vscode-remote-try-go/blob/main/.devcontainer/devcontainer.json>
* <https://bitfieldconsulting.com/golang/docker-image>
* <https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/>
* <https://www.docker.com/blog/containerize-your-go-developer-environment-part-2/>
* <https://www.docker.com/blog/containerize-your-go-developer-environment-part-3/>


