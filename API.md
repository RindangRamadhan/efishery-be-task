# Project: eFishery BE Task
# 📁 Auth App 


## End-point: Register
### Method: POST
>```
>https://efisheryapi.infomediaku.com/api/auth/register
>```
### Body (**raw**)

```json
{
    "username": "rama",
    "name": "Rindang Ramadhan",
    "phone": "0811252497",
    "role": "admin"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Login
### Method: POST
>```
>https://efisheryapi.infomediaku.com/api/auth/login
>```
### Body (**raw**)

```json
{
    "phone": "0811252497",
    "password": "Vidx"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Login Check
### Method: GET
>```
>https://efisheryapi.infomediaku.com/api/auth/login-check
>```
### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUxMTQ3NjIsIm5iZiI6MTY0NTExMTE2Miwic3ViIjp7Im5hbWUiOiJSaW5kYW5nIFJhbWFkaGFuIiwicGhvbmUiOiIwODExMjUyNDk3Iiwicm9sZSI6ImFkbWluIiwiY3JlYXRlZF9hdCI6IjIwMjItMDItMTdUMjI6MTY6MzgrMDcwMCJ9fQ.cYaX0onkdzYQ6rSoiS1IzqnkwT72UkWoY-727IuNasw|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Fetch App 


## End-point: Claims Token
### Method: GET
>```
>https://efisheryapi1.infomediaku.com/fetch/claims
>```
### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUxMjU0MTMsIm5iZiI6MTY0NTEyMTgxMywic3ViIjp7Im5hbWUiOiJSYW1hZGhhbiBSaW5kYW5nIiwicGhvbmUiOiIwODExMjUyNDk1Iiwicm9sZSI6ImJpYXNhIiwiY3JlYXRlZF9hdCI6IjIwMjItMDItMThUMDA6MDI6MTcrMDcwMCJ9fQ.Hc_kK7W6lN65ufBKk8v7rP3BAsqAm8NuIqbKuNzpRmA|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Resource (Role Bebas)
### Method: GET
>```
>https://efisheryapi1.infomediaku.com/fetch/resource-bebas
>```
### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUxMjU0MTMsIm5iZiI6MTY0NTEyMTgxMywic3ViIjp7Im5hbWUiOiJSYW1hZGhhbiBSaW5kYW5nIiwicGhvbmUiOiIwODExMjUyNDk1Iiwicm9sZSI6ImJpYXNhIiwiY3JlYXRlZF9hdCI6IjIwMjItMDItMThUMDA6MDI6MTcrMDcwMCJ9fQ.Hc_kK7W6lN65ufBKk8v7rP3BAsqAm8NuIqbKuNzpRmA|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Resource (Role Admin)
### Method: GET
>```
>https://efisheryapi1.infomediaku.com/fetch/resource-admin
>```
### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUxMjU0MTMsIm5iZiI6MTY0NTEyMTgxMywic3ViIjp7Im5hbWUiOiJSYW1hZGhhbiBSaW5kYW5nIiwicGhvbmUiOiIwODExMjUyNDk1Iiwicm9sZSI6ImJpYXNhIiwiY3JlYXRlZF9hdCI6IjIwMjItMDItMThUMDA6MDI6MTcrMDcwMCJ9fQ.Hc_kK7W6lN65ufBKk8v7rP3BAsqAm8NuIqbKuNzpRmA|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________