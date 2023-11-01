# Superindo API Assessment

## Requirment tools
    Docker versi 24.0.6 [https://docs.docker.com/desktop/install/windows-install/]
    Docker-Compose versi v2.21.0-desktop.1
    Postman
    Golang go1.19

## To Do
    install docker dan docker-compose
    Install postman
    Install git
    clone repo [https://github.com/dirgadm/superindo_api.git]

## Serve
    ```
    - make up = create container and image in docker, so it can be serve a server and running the application
    - make down = shut down image and container in docker which is running by system
    - make destroy = delete all container and image in docker
    - make log = watch your development while it's running
    ```

## Endpoint Testing 
    - available in `./doc/Superindo.postman_collection.json` and ready to import to postman