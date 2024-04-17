docker rm leminContainer
docker rmi lemin_image
docker build -t lemin_image .
docker run -it --name leminContainer -p 2030:2030 lemin_image