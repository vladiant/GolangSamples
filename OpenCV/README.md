docker build -t gocv .
docker run -it gocv /bin/bash

docker run -it -v /home/vladiant/Desktop/ToPublish/golang/OpenCV:/test gocv /bin/bash

https://hub.docker.com/r/gocv/opencv

https://marcosnietoblog.wordpress.com/2017/04/30/docker-image-with-opencv-with-x11-forwarding-for-gui/

xhost +
docker run -it --net=host --ipc=host -v /home/vladiant/Desktop/ToPublish/golang/OpenCV:/test -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix gocv /bin/bash

https://bitfieldconsulting.com/golang/docker-image
COPY main.go go.* /src/

https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/

https://hub.docker.com/r/gocv/opencv
