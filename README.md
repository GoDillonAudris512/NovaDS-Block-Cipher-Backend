# 🔐 NovaDS Modern Block Cipher
> Back-end side of NovaDS Modern Block Cipher Web using Go languange. Provides endpoints to encrypt and decrypt block in various modes using NovaDS algorithm

## Project Structure
```bash
.
├─── algorithms
│   ├─── feistelNetwork.go
│   ├─── helper.go
│   ├─── novaDSCipher.go
│   ├─── roundFunction.go
│   └─── roundKeyGenerator.go
├─── constants
│   └─── constants.go
├─── handlers
│   ├─── cbcHandlers.go
│   ├─── cfbHandlers.go
│   ├─── counterHandlers.go
│   ├─── ecbHandlers.go
│   ├─── generalHandlers.go
│   └─── ofbHandlers.go
├─── middlewares
│   └─── corsMiddleware.go
├─── model
│   ├─── cbc.go
│   ├─── cfb.go
│   ├─── counter.go
│   ├─── ecb.go
│   └─── ofb.go
├─── router
│   └─── router.go
├─── .env.example
├─── .gitignore
├─── go.mod
├─── go.sum
├─── main.go
└─── README.md
```

## User Interfaces
User Interface is designed and implemented on the front-end side. Further implementation stated on [this repository](https://github.com/Salomo309/NovaDS-Block-Cipher-Frontend)

## ⚙️ &nbsp;How to Run the Program

Clone this repository from terminal with this command
``` bash
$ git clone https://github.com/GoDillonAudris512/NovaDS-Block-Cipher-Backend.git
```

### Run the application on development server
1. Create a .env file inside the repository directory using .env.example file as the template. You can keep the variables blank. The server should automatically use port 8080 as the default port 
2. Run the server using this following command
    ``` bash
    go run main.go
    ```

If you do it correctly, the back-end development server should be running. You can also check the server by opening http://localhost:8080/api. To use back-end side functionalities, don't forget to also run the front-end side. Further explanation on how to run the front-end development server stated on [this repository](https://github.com/Salomo309/NovaDS-Block-Cipher-Frontend)


## 🔑 &nbsp;Endpoints
| Endpoint                             |  Method  |   Usage  |
| ------------------------------------ | :------: | -------- |
| /api/cbc                        | POST     | Users can perform encryption and decryption with NovaDS using CBC mode
| /api/cfb                      | POST     | Users can perform encryption and decryption uwith NovaDS using CFB mode
| /api/counter               | POST     | Users can perform encryption and decryption with NovaDS using Counter mode
| /api/ecb                       | POST     | Users can perform encryption and decryption with NovaDS using ECB mode
| /api/ofb                          | POST     | Users can perform encryption and decryption uwith NovaDS using OFB mode

## Authors
| NIM      | Name                           |
| -------- | ------------------------------ |
| 13521062 | Go Dillon Audris               |
| 13521063 | Salomo Reinhart Gregory Manalu |