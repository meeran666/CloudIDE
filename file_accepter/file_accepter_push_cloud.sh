echo -e  "docker rmi -f meeran666/file_accepter"
docker rmi -f meeran666/file_accepter
echo -e  "docker build -t meeran666/file_accepter ."
docker build -t meeran666/file_accepter .
echo -e  "docker push meeran666/file_accepter"
docker push meeran666/file_accepter