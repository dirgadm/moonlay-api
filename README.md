# Moonlay API Assessment
```
```
### Teknologi dan library digunakan:
| Teknologi   | Version | Link |
| ----------- | ---------------- | ------------------- |
| Golang      | v1.19 or later   | [Go Download](https://go.dev/dl)  |
| Go Echo Framework     | v4     | [Echo Installation](https://echo.labstack.com/guide/#installation) | 
| GORM | v2 | [GORM Installation](https://gorm.io/docs/#Install) |
| PostgreSQL | v13 or later | [PostgreSQL Download](https://www.postgresql.org/download/) |
| Docker | v24.0.6 or later |  |
| Docker compose| v12.21.0 or later |  |
<br>

## To Do
    install docker dan docker-compose
    Install postman
    Install git
    clone repo [https://github.com/dirgadm/moonlay-api]

## Ruuning Server
    1. command: *docker compose up -d*, Menjalankan file docker compose di sisi background, sekaligus melakukan migrasi sql ke database postgre, migrasi file terdapat di file *moonlay.sql*
    2. command: *go run main.go* , menjalankan server dengan port 9090
    3. command: *go test* , masuk terlebih dahulu ke directory yang akan dilakukan unit test, dan *go test -cover* untuk melihat coverage dari unit test
    ```

## Endpoint Testing 
    - available in `./doc/Moonlay.postman_collection.json` and ready to import to postman
    - base_url: http://localhost:9090/v1

## List Endpoint
1. [METHOD:GET] Menampilkan data all list ( include pagination, filter[Search By: title, description] ) dengan atau tanpa preload sub list (dynamic)
2. [METHOD:GET] Menampilkan data detail list by list id.
3. [METHOD:GET] Menampilkan data all sub list by list id ( include pagination, filter[Search By: title, description] )
4. [METHOD:GET] Menampilkan data detail sub list by sub list id.
5. [METHOD:POST] Menambahkan data list.
6. [METHOD:POST] Menambahkan data sub list untuk spesifik list.
7. [METHOD:POST/PUT] Mengubah data list/sub list dengan kritera input diatas. 
8. [METHOD:DELETE] Menghapus data list/sub list. 
9. [METHOD:POST] Upload file single/multi, dengan extension hanya txt dan pdf 