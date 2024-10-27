# Data for test
```js
const userData = {
    username: 'testuser',
    email: 'test@example.com',
    password: 'password123',
  };
```
## Run server
```bash
$ node index.js 
Register response: { success: true, message: 'Пользователь успешно зарегистрврован' }
Login response: {
  token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyNC0xMC0yN1QwNToxNDo1OFoiLCJlbWFpbCI6InRlc3Q
zQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyNTg3MzAwLCJ1c2VyX2lkIjo5LCJ1c2VybmFtZSI6InRlc3QzdXNlciJ9.ueTy0CSiTAlJQ1wRQyywD4Vn
Y7ECNIgvDNXXX5JL1kc',
  message: 'Успешная авторизация'
}
```
# JWT encoded format
```
{
  "created_at": "2024-10-27T05:14:58Z",
  "email": "test3@example.com",
  "exp": 1732587300,
  "user_id": 9,
  "username": "test3user"
}
```