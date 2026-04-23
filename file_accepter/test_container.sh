docker rmi -f file_accepter
docker rm file_accepter_con
docker build -t file_accepter .
docker run -it --name file_accepter_con -e BASE_FOLDER=user1 file_accepter