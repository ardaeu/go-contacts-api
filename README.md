# Go Contacts API

## Proje Hakkında

Bu proje Go dili ile geliştirilmiş basit bir RESTful API'dir. PostgreSQL veri tabanı kullanılarak kişiler üzerinde CRUD işlemleri yapılabilmektedir.

## Özellikler

- Yeni kişi oluşturma (POST /contacts)
- Tüm kişileri listeleme (GET /contacts)
- ID ile kişi bilgisi getirme (GET /contacts/{id})
- Kişi güncelleme (PUT /contacts/{id})
- Kişi silme (DELETE /contacts/{id})

## Teknolojiler

- Go 1.20+
- PostgreSQL
- Chi Router
- go.mod modülleri: github.com/go-chi/chi, github.com/lib/pq, github.com/google/uuid (eğer kullanıldı)

## Kurulum ve Çalıştırma

1. PostgreSQL kur ve çalıştır.
2. Veritabanı oluştur:
   ```sql
   CREATE DATABASE go_contacts;
   ```
3. Tabloyu oluştur:
   ```sql
   CREATE TABLE contacts (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL,
       email TEXT NOT NULL UNIQUE,
       phone TEXT NOT NULL
   );
   ```
4. Projeyi klonla:
   ```bash
   git clone <repository-url>
   cd go-contacts-api
   ```
5. Modülleri indir:
   ```bash
   go mod tidy
   ```
6. Sunucuyu başlat:
   ```bash
   go run cmd/api/main.go
   ```
7. API `http://localhost:8088` adresinde çalışır.

## Test Etme (curl Örnekleri)

Yeni kayıt ekle:

```bash
curl -X POST http://localhost:8088/contacts \
-H "Content-Type: application/json" \
-d '{"name":"Arda Eymen Ulus","email":"arda@example.com","phone":"+90 555 123 4567"}'
```

Tüm kayıtları listele:

```bash
curl http://localhost:8088/contacts
```

ID ile kayıt getir:

```bash
curl http://localhost:8088/contacts/1
```

Kayıt güncelle:

```bash
curl -X PUT http://localhost:8088/contacts/1 \
-H "Content-Type: application/json" \
-d '{"name":"Arda U.","email":"arda.u@example.com","phone":"+90 555 765 4321"}'
```

Kayıt sil:

```bash
curl -X DELETE http://localhost:8088/contacts/1
```

## Lisans

MIT License

## İletişim

Arda Eymen Ulus - ulusarda9@gmail.com - https://www.linkedin.com/in/ardaeu/
