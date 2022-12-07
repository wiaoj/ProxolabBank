// -------- kredi türü id'leri ve faiz vade id'leri -----------
// Konut = 1
// Tüketici = 2
// Mevduat = 3
// 3 ay = 1
// 6 ay = 2
// 12 ay = 3
// 24 ay = 4
// 36 ay = 5
// 5 yıl = 6
// 10 yıl = 7

// ------------------------


// Hazır verdiğimiz kullanıcı adı ve şifre ile de işlem yapabilirsiniz yada o kullanıcı ile login olduktan sonra başka kullanıcı register edebilirsiniz

// login olmak için

```js
POST http://localhost/api/auth/login
```
```json
{
    "username": "proxolab", 
    "password": "proxolab"
}
```

//yeni kullanıcı kayıt etmek için (proxolab kullanıcısı sadece bu yetkiye sahip)

```js
POST http://localhost/api/auth/register
```

```json
{
    {
    "first_name": "bertan",
    "last_name": "tokgöz",
    "username": "bertan123",
    "age": 19,
    "email": "bertan@gmail.com",
    "phone_number": "000-000-00-00",
    "password": "password"
}
```

// ------------------------------

// banka eklerken gereken parametler 
// http://localhost/api/banks  (POST method)
{
    "bank_name":  // string
}

// --------------------


// tüm bankaları çekerken url
// http://localhost/api/banks    (GET method)


// -----------------------


// tek banka çekerken url 
// http://localhost/api/banks/id  (GET method)   buradaki id kısmına banka id gelecek 1 2 3 artık eklediğiniz bankalarınızın id si ne ise bir adet banka için bu


// -----------------------------

// banka silerken gereken parametre ve url
// http://localhost/api/banks  (DELETE method)
{
    "id":    // integer
}

// -----------------------------

// interest (faiz) eklerken gereken parametreler
// http://localhost/api/interests (POST method)
{
    "bank_id" : , //integer
    "interest" : , //float
    "time_option" : , //integer
    "credit_type" :  //integer 
}

// --------------------------

// interest silerken gereken parametreler
// http://localhost/api/interests (DELETE method)
{
    "id" : , // integer
    "bank_id" : // integer
}

// ----------------------


```js
GET http://localhost:3000/api/interests/q?bankId=2&creditTypeId=1&timeOptionId=7&interestOrderType=asc
```

```json
{
    "message": "faizler asc şeklinde sıralanıp getirildi",
    "items": [
        {
            "bank_id": 2,
            "bank_name": "test name3",
            "interest": 0.1,
            "time_option_id": 7,
            "time_option_description": "10 yıl",
            "credit_type_id": 1,
            "credit_type_description": "Konut Kredisi"
        },
        {
            "bank_id": 2,
            "bank_name": "test name3",
            "interest": 1.1,
            "time_option_id": 7,
            "time_option_description": "10 yıl",
            "credit_type_id": 1,
            "credit_type_description": "Konut Kredisi"
        }
    ]
}
```


/* credit_type (kredi türü) ve time_options (vade) ile alakalı açıklama;
--------------------
---id'ler---
-Kredi Türleri-
-Konut = 1  5 yıl ve 10 yıl vadeleri mevcut
-Tüketici = 2 12 ay 24 ay ve 36 ay vadeleri mevcut
-Mevduat = 3 3 ay 6 ay ve 12 ay vadeleri mevcut

-Vadeler-
3 ay = 1
6 ay = 2
12 ay = 3
24 ay = 4
36 ay = 5
5 yıl = 6
10 yıl = 7
--------------------
Konut kredi türünün id'si -> 1 . İki çeşit vadesi var (5 yıl -> id'si 6 | 10 yıl -> id'si : 7)
Tüketici kredi türünün id'si -> 2 . Üç çeşit vadesi var (12 ay -> id'si 3 | 24 ay -> id'si : 4 | 36 ay -> id'si : 5)
Mevduat faizi türünün id'si -> 3 . Üç çeşit vadesi var (3 ay -> id'si 1 | 6 ay -> id'si : 2 | 12 ay -> id'si : 3)
--------------------
Örnek olarak konut kredisine 12 ay vade girilemez. api'den hata döner. */
