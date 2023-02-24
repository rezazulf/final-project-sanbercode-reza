# final-project-sanbercode-42-reza

### Build With
    1. Go version 1.19
    2. Database Postgres
    3. Framework Go Gin 

### Get started
Download terlebih dahulu library-library yang akan digunakan:

    go get -u "github.com/gin-gonic/gin"
    go get -u "github.com/lib/pq"
    go get -u "github.com/rubenv/sql-migrate"
    go get -u "github.com/gobuffalo/packr/v2"
    go get -u "github.com/joho/godotenv"

<h1> Maaf, kurang lengkap. Saya kurang tidur. </h1>
<h2> Untuk view body dan param url lebih lengkap dapat dilihat di : <a href="https://www.postman.com/rezazulf/workspace/fp-sanbercode/collection/12473257-c3833057-9583-4014-8f11-d47123030f6d?ctx=documentation">disini</a></h2>
<h2> link diatas workspace untuk postman. lengkap dengan body json</h2>

### username|pass : admin|admin => role admin, user|user => role customer


### APIs
### REGISTRATION
```
  POST /register
```
#### Parameter Body
```json
{
    "username":"user1",
    "password":"user1"
}
```

##### Parameter Body
```json
{
    "username":"user",
    "password":"user"
}
```
##### Contoh Response Sukses
```json
{
    "message": "Success to Create User"
}
```
### LOGIN
```
  POST /login
```
##### Parameter Body
```json
{
    "username":"admin",
    "password":"admin"
}
```
##### Response berhasil login
```json
{
  "message": "Success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcwNDEyNTQsIkVtYWlsIjoibWVnYUBlbWFpbC5jb20iLCJQYXNzd29yZCI6ImFkbWluT0sxMjMiLCJSb2xlIjoiYWRtaW4ifQ.qfctmzyfBT97it3VMUQfPkRGoo8HKus_xX51vgI3j6Q"
}

```
### logout
```
  POST /logout
```
##### Parameter
```
none
```
##### Response berhasil logout
```json
{
    "message": "Berhasil logout"
}

```

### CATEGORIES
#### 1. Get All Categories
```
  GET /categories
```
##### Contoh Response Sukses
```json
{
    "result": [
        {
            "id": 1,
            "name": "Headset",
            "created_at": "2023-02-24T04:58:13.328464+07:00",
            "updated_at": "2023-02-24T04:58:13.328464+07:00"
        },
        {
            "id": 2,
            "name": "Mouse",
            "created_at": "2023-02-24T04:58:16.71952+07:00",
            "updated_at": "2023-02-24T04:58:16.71952+07:00"
        },
        {
            "id": 3,
            "name": "Keyboard",
            "created_at": "2023-02-24T04:58:20.439873+07:00",
            "updated_at": "2023-02-24T04:58:44+07:00"
        },
        {
            "id": 4,
            "name": "Mousepad",
            "created_at": "2023-02-24T04:58:34.686515+07:00",
            "updated_at": "2023-02-24T04:58:34.686515+07:00"
        },
        {
            "id": 5,
            "name": "Mousepad",
            "created_at": "2023-02-24T05:12:00.762588+07:00",
            "updated_at": "2023-02-24T05:12:00.762588+07:00"
        }
    ]
}
```
#### 2. Insert Categories
```
  POST /categories
```
##### Contoh Request
```json
{
    "name": "Mousepad"
}
```
##### Contoh Response Sukses
```json
{
    "result": "Success Insert Category"
}
```
#### 3. Update Categories
```
  PUT /categories/:id
```
##### Contoh Request
```json
{
    "name":"Keyboard"
}
```
##### Contoh Response Sukses
```json
{
    "result": "Success Update Category"
}
```

### Status


#### 1. Create Status
```
  POST /status
```
##### Contoh Request
```json
{
    "status":"Out of Stock"
}
```
##### Contoh Response Sukses
```json
{
    "result": "Sukses Menambahkan Status"
}
```
#### 4. Get All Status
```
  GET /status
```
##### Contoh Response Sukses
```json
{
    "result": [
        {
            "id": 1,
            "status": "In-Stock"
        },
        {
            "id": 2,
            "status": "Out of Stock"
        }
    ]
}
```


### Product
#### 1. Create Product (ADMIN AUTH)
```
  POST /product
```
##### Contoh Request
```json
{
    "name":"HaiperX Awan 2",
    "description":"Headset Gaming dari HaiperX",
    "price": 4,
    "image_url": "https://ssl-product-images.www8-hp.com/digmedialib/prodimg/lowres/c08149407.png",
    "stock": 10, 
    "status_id": 1,
    "category_id": 1
}
```
##### Contoh Response Sukses
```json
{
    "result": "Sukses Menambahkan Produk"
}
```
#### 2. Update
```
  PUT /product/:id
```
##### Contoh Request
```json
{
    "name": "Logitech Mouse GeminX",
    "description": "Mouse Gaming dari Logitech",
    "price": 5,
    "image_url": "https://asset.kompas.com/crops/uCsYc5Nazq8k_90zI5MPMvM9zgw=/63x0:720x438/750x500/data/photo/2022/02/14/6209b1e3f12c0.png",
    "stock": 16,
    "status_id": 1,
    "category_id": 2
}
```
##### Contoh Response Sukses
```json
{
    "result": "Sukses Update Produk"
}

```
#### 3. Delete
```
  DELETE /product/:id
```
##### Contoh Response Sukses
```json
{
  "result": "Sukses Menghapus Produk"
}
```
#### 5. Get All Product
```
GET /product/
```
##### Contoh Response Sukses
```json
{
    "result": [
        {
            "id": 1,
            "name": "Logitech Mouse GeminX",
            "description": "Mouse Gaming dari Logitech",
            "price": 5,
            "image_url": "https://asset.kompas.com/crops/uCsYc5Nazq8k_90zI5MPMvM9zgw=/63x0:720x438/750x500/data/photo/2022/02/14/6209b1e3f12c0.png",
            "stock": 16,
            "status_id": 1,
            "category_id": 2,
            "created_at": "2023-02-24T05:25:22.5157+07:00",
            "updated_at": "2023-02-24T05:27:33+07:00"
        }
    ]
}
```
#### Create Status
```
  POST /buy/:productid/customer/:customerid
```
##### Contoh Request
```json
{
    "username": "user",
    "password" : "user",
    "sum_item" : 1
}
```

