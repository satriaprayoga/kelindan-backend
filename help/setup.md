# install docker

# install redis with docker
docker pull redis
docker images -a
docker run -d -p 6379:6379 --name redistest redis
docker ps
docker exec -it redistest sh
redis-cli