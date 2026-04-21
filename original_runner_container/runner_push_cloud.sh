echo -e  "docker rmi -f meeran666/ide_image"
docker rmi -f meeran666/ide_image
echo -e  "docker build -t meeran666/ide_image ."
docker build -t meeran666/ide_image .
echo -e  "docker push meeran666/ide_image"
docker push meeran666/ide_image