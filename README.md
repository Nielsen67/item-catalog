# Description
Item Catalog merupakan Website yang dikembangkan sebagai untuk mendaftarkan barang - barang dalam sebuah katalog dengan konsep minimalis
Item Catalog memungkinkan pengguna untuk memperoleh item, menambahkan item dan menghapus seluruh item. Item Catalog didesain dengan arsitektur serverless dengan mengandalkan Vercel dalam deployment. Website ini dapat diakses melalui domain berikut : 
https://item-catalog.vercel.app/

## Requirement
Sebelum menerapkan API ke Vercel, pastikan hal - hal berikut:
- **Go** (Versi 1.21 ke atas)
- **Vercel CLI**
- **Git** 
- **Node.js and npm** 
- **Akun Vercel** 


## Deployment to Vercel
Berikut adalah langkah - langkah dalam menjalankan API di Vercel : 

### 1. Instalasi Vercel CLI
Instal Vercel CLI secara global menggunakan npm:
```bash
npm install -g vercel
```


### 2. Login ke Vercel
Masuk ke Vercel dengan akun yang valid
```bash
vercel login
```

### 3. Pastikan Kelengkapan Proyek
Pastikan direktori proyek dilengkapi sepert go.mod, go.sum, main.go, vercel.json dan controller, model, service dan lainnya.
Semua keperluan di atas disesuaikan dengan kebutuhan pengembangan API.

### 4. Deployment API ke Preview (Staging)
Pastikan berada pada direktori yang tepat dan gunakan perintah : 
```bash
vercel
```
Lalu pilih penyesuaian deployment sesuai dengan keinginan seperti Integrasi dengan Proyek sebelumnya, Nama Proyek dan Lainnya.

### 5. Akses Preview Deployment
Pastikan Preview Deployment berjalan lancar dan normal

#### 6. Penerapan ke Production
Jika API sudah berjalan sesuai dengan ekspektasi melalui Preview Deployment, maka deployment dapat dilanjutkan ke stage Production dengan : 
```bash
vercel --prod
```


Dengan menerapkan langkah - langkah di atas, Anda sudah dapat mengakses API melalui Vercel. Namun, Anda juga bisa melakukan beberapa tindakan tambahan seperti : 
1. Testing API melalui cURL atau Postman
2. Pengaturan Domain yang lebih Advance
3. Pengaktifan Security

