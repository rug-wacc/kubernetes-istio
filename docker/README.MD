Here are the commands used in the Docker Demo part of the lecture. 
```sh
docker run hello-world
docker run -it ubuntu bash
    apt-get update
    apt-get install curl
docker ps -a
docker commit [container name] [tag=ubuntu-curl]
docker run -it --name myubuntu ubuntu-curl bash
docker volume create data-v
docker run -it --name myubuntu -v data-v:/home/data ubuntu-curl
    (execute mkdir in /home and /home/data to show persistance)
docker run -it --name myubuntu2 -v data-v:/home/data2 ubuntu-curl
```