## Certificate
![sertifikat](https://github.com/firdausalif/challenge-todolist/blob/main/cert.png?raw=true)

## Description
DevCode Challenge as Backend Engineer : Api To Do List
https://devcode.gethired.id/challenge/api-todo

API Documentation
https://www.getpostman.com/collections/03cd6360184be37eb5b6

## Challange

###Section 1.
Goals pada section ini terdiri dari:
1. Engineer mampu membuat API yang disimpan berupa image Docker
2. Engineer mampu membuat API yang dapat terkoneksi ke database MySQL
3. Engineer mampu membuat routing API sesuai dengan dokumen perancangan pada Postman

###Detail Test Scenario

#### Create activity
- Berhasil menambahkan data activity baru.
- Negative case : Menampilkan response status 400, gagal menambahkan activity baru jika title tidak diisi.

#### Update activity
- Berhasil melakukan update data activity.
- Negative case : Menampilkan response status 400, gagal melakkukan update activity data jika id pada parameter tidak ditemukan.

##### Delete activity
- Berhasil menghapus data activity.
- Negative case : Menampilkan response status 400, gagal hapus data activity jika id pada parameter tidak ditemukan.

#### Get Detail Activity
- Berhasil menampilkan response get detail activity.
- Negative case : Menampilkan response status 400, gagal get detail pada parameter id yang tidak ditemukan.

##### Get List Activity
- Berhasil menampilkan response get list activity.

##### Create todo
- Berhasil menambahkan data Todo baru.
- Negative case : Menampilkan response status 400, gagal menambahkan data jika title tidak diisi.
- Negative case : Menampilkan response status 400, gagal menambahkan data jika activity_group_id tidak diisi.

#### Update todo
- Berhasil melakukan update data title.
- Berhasil melakukan update data status.
- Negative case : Menampilkan response status 400, gagal update data todo jika id pada parameter tidak ditemukan.

##### Delete todo
- Berhasil melakukan hapus data todo.
- Negative case : Menampilkan response status 400, Gagal hapus todo jika parameter id tidak ditemukan.

#### Get Detail todo
- Berhasil menampilkan response get detail todo.
- Negative case : Menampilkan response status 400, gagal get detail todo pada parameter id yang tidak ditemukan.

#### Get List todo
- Berhasil menampilkan response get list todo.
- Negative case : Menampilkan data kosong ketika id pada parameter tidak ditemukan.

### SECTION 2 - PERFORMANCE
Goals pada section ini terdiri dari 
1. Engineer mampu melakukan optimasi pada API

### Detail Test Scenario

#### Stress Test
- Berhasil menjalankan 1000 request dalam 1 count currency.
Performance
- Average response berada pada nilai kurang dari 100ms dari stress test yang dijalankan.
Optimation
- Optimasi file size image docker harus kurang dari 300MB.
