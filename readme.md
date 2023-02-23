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

### APIs
### REGISTRATION
```
  POST /register
```

##### Parameter
| Parameter    | Tipe Data | Deskripsi                                                                                |
|--------------|-----------|------------------------------------------------------------------------------------------|
| username     | string    | register username untuk user/customer                                                    | 
| password     | string    | register password untuk user/customer                                                    |

##### Contoh Request
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
##### Parameter
| Parameter | Tipe Data | Deskripsi                                   |
|-----------|-----------|---------------------------------------------|
| username  | string    | menggunakan username yang telah didaftarkan | 
| password  | string    | menggunakan password yang telah didaftarkan |
##### Contoh Request
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
##### Parameter
| Parameter      | Tipe Data  | Deskripsi                                   |
|----------------|------------|---------------------------------------------|
| name           | string     | untuk menginput nama kategori               | 
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

### PRODUCT
##### Parameter
| Parameter     | Tipe Data | Deskripsi                                                            |
|---------------|-----------|----------------------------------------------------------------------|
| status        | string    | untuk input status. Antara In-Stock/Out of Stock                     |


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
##### Parameter
| Parameter     | Tipe Data   | Deskripsi                                                |
|---------------|-------------|----------------------------------------------------------|
| name          | string      | Nama event, contohnya: Konser Akbar All Star             |
| description   | string      | Deskripsi event, contoh: Lokasi event                    |
| price         | int         | Harga untuk sebuah produk                                |
| image_url     | string      | URL untuk gambar                                         |
| stock         | int         | stock untuk cek stock ada atau tidak                     |
| status_id     | int         | status_id, contohnya id 1 = In-Stock                     |
| category_id   | int         | category_id, contohnya id 1 = mouse                      |

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
##### Contoh Response Gagal
```
<h1> Waktunya gak sempet buat bikin error handling </h1>
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
```
{
  "data": [
    {
      "id": 2,
      "name": "Blackpink World Tour 2023",
      "description": "Concert Blackpink ICE BSD",
      "start_date": "2023-05-05T00:00:00Z",
      "end_date": "2023-05-05T00:00:00Z",
      "category_id": 2,
      "category_name": "Concert"
    }
  ]
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
### Beli Barang
##### Parameter
