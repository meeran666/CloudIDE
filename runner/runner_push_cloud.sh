# echo -e  "docker rmi -f meeran666/ide_image_go"
# docker rmi -f meeran666/ide_image_go
# echo -e  "docker build -t meeran666/ide_image_go ."
# docker build -f Dockerfile.go_image -t meeran666/ide_image_go .
# echo -e  "docker push meeran666/ide_image_go"
# docker push meeran666/ide_image_go

echo -e  "docker rmi -f meeran666/ide_image_node"
docker rmi -f meeran666/ide_image_node
echo -e  "docker build -t meeran666/ide_image_node ."
docker build -f Dockerfile.node_image -t meeran666/ide_image_node .
echo -e  "docker push meeran666/ide_image_node"
docker push meeran666/ide_image_node
