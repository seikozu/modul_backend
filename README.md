# Dokumentasi REST API Students

| Metode | Endpoint | Parameter | Contoh Body | Status | Contoh Respons |
|---|---|---|---|---|---|
| **GET** | `/api/v1/students` | `page`, `limit`, `search`, `sort`, `order`, `is_active` | - | `200 OK` | `{"success":true, "data":[...], "meta":{...}}` |
| **GET** | `/api/v1/students/:id` | `id` (path) | - | `200 OK`, `404 Not Found` | `{"success":true, "data":{"id":1, "nim":"123", ...}}` |
| **POST** | `/api/v1/students` | - | `{"nim":"123", "name":"Jinhsi", "grade":90}` | `201 Created`, `422 Unprocessable`| `{"success":true, "message":"student berhasil dibuat", "data":{...}}` |
| **PUT** | `/api/v1/students/:id` | `id` (path) | `{"nim":"123", "name":"Baru", "grade":95, "is_active":true}` | `200 OK`, `422 Unprocessable` | `{"success":true, "message":"student berhasil diganti seluruhnya", "data":{...}}` |
| **PATCH** | `/api/v1/students/:id` | `id` (path) | `{"is_active":false}` | `200 OK`, `400 Bad Request` | `{"success":true, "message":"student berhasil diperbarui sebagian", "data":{...}}` |
| **DELETE**| `/api/v1/students/:id` | `id` (path) | - | `204 No Content`, `404 Not Found` | *(Tanpa body)* |