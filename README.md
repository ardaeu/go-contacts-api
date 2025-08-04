Go Contacts API

Proje Hakkında
Bu proje Go dili ile geliştirilmiş basit bir RESTful API'dir. PostgreSQL veri tabanı kullanılarak kişiler üzerinde CRUD işlemleri yapılabilmektedir.

Özellikler

- Yeni kişi oluşturma (POST /contacts)
- Tüm kişileri listeleme (GET /contacts)
- ID ile kişi bilgisi getirme (GET /contacts/{id})
- Kişi güncelleme (PUT /contacts/{id})
- Kişi silme (DELETE /contacts/{id})

Teknolojiler

- Go 1.20+
- PostgreSQL
- Chi Router
- go.mod modülleri: github.com/go-chi/chi, github.com/lib/pq, github.com/google/uuid (eğer kullanıldı)

Kurulum ve Çalıştırma

1. PostgreSQL kur ve çalıştır.
2. Veritabanı oluştur:
   CREATE DATABASE go_contacts;
3. Tabloyu oluştur:
   CREATE TABLE contacts (
   id SERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   email TEXT NOT NULL UNIQUE,
   phone TEXT NOT NULL
   );
4. Projeyi klonla:
   git clone <repository-url>
   cd go-contacts-api
5. Modülleri indir:
   go mod tidy
6. Sunucuyu başlat:
   go run cmd/api/main.go
7. API http://localhost:8080 adresinde çalışır.

Test Etme (curl Örnekleri)
Yeni kayıt ekle:
curl -X POST http://localhost:8080/contacts \
-H "Content-Type: application/json" \
-d '{"name":"Arda Eymen Ulus","email":"arda@example.com","phone":"+90 555 123 4567"}'

Tüm kayıtları listele:
curl http://localhost:8080/contacts

ID ile kayıt getir:
curl http://localhost:8080/contacts/1

Kayıt güncelle:
curl -X PUT http://localhost:8080/contacts/1 \
-H "Content-Type: application/json" \
-d '{"name":"Arda U.","email":"arda.u@example.com","phone":"+90 555 765 4321"}'

Kayıt sil:
curl -X DELETE http://localhost:8080/contacts/1

Lisans
MIT License

İletişim
Arda Eymen Ulus - arda@example.com
