# GraphQL with Golang #
This is an example how to use GraphQL with Golang <br/>
In this example i am using Echo Framework as http router and GORM as Database

## Table Structure ##
```
CREATE TABLE `level_user` (
 `id_level` int(2) NOT NULL AUTO_INCREMENT,
 `nama` varchar(100) NOT NULL,
 PRIMARY KEY (`id_level`)
);

CREATE TABLE `users` (
 `email` varchar(100) NOT NULL,
 `nama` varchar(50) NOT NULL,
 `no_handphone` varchar(15) DEFAULT NULL,
 `alamat` text DEFAULT NULL,
 `ktp` varchar(255) DEFAULT NULL,
 `id_level` int(2) NOT NULL,
 PRIMARY KEY (`email`) USING BTREE
);
```
You can copy this table structure to run this project. Change your database configuration in `config/database.go`

## Tes with Postman ##
You can test this graphql api with postman. Example of result is like that if it works : <br/>
![](../go-graphql-api/01.png)<br/>
<i>Request all field without filter</i><br/>
![](../go-graphql-api/02.png)<br/>
<i>Request with specified field</i><br/>
![](../go-graphql-api/03.png)<br/>
<i>Request with specified field and filter</i>