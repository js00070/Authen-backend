# 说明

## 接口说明
请求参数与返回参数均在http body中，且为json字符串

### 注册
- url: /user/register
- 方法: POST
- 请求参数:
  - username 用户名
  - passwd 密码
- 返回参数:
  - code 0表示成功，-1表示失败
  - msg 成功或错误信息

### 登陆
- url: /user/login
- 方法: POST
- 请求参数:
  - username 用户名
  - passwd 密码
- 返回参数:
  - code 0表示成功，-1表示失败
  - msg 成功或错误信息
  - token 登陆成功后返回的token字符串

### 上传文件hash
- url: /file/hash
- 方法: POST，header中带有token字段
- 请求参数:
  - filename 文件名
  - hash 文件hash串
  - timestamp 文件创建的unix秒时间戳
- 返回参数:
  - code 0表示成功，-1表示失败
  - msg 成功或错误信息

### 文件上传
- url: /file/upload
使用的是tus协议，详情见https://tus.io/
client使用方式见test目录下的main.go，需要在上传时，在Metadata中加入token字段

### 获取已上传文件列表
- url: /filelist
- 方法: GET，header中带有token字段
- 请求参数: 无
- 返回参数:
  - code 0表示成功，-1表示失败
  - msg 成功或错误信息
  - filelist: 列表，内含详细信息

## 使用例子

### 上传文件hash
```bash
curl -H 'token:4e1610bdab129d2143652093de01e15200000000000000000000000000000000' -X POST localhost:8080/file/hash -d '{"filename":"my-file.txt","hash":"123321123","timestamp":1233211233}'

{"code":0,"msg":"success"}
```

### 获取文件列表
```bash
curl -H 'token:4e1610bdab129d2143652093de01e15200000000000000000000000000000000' localhost:8080/filelist

{"code":0,"filelist":[{"ID":5,"CreatedAt":"2020-05-03T16:46:03+08:00","UpdatedAt":"2020-05-03T16:46:20+08:00","DeletedAt":null,"UserID":1,"FileName":"my-file.txt","Path":"/tmp/a9b7403fce75a7960886d8e14ab10446","Hash":"123321123","TimeStamp":"2009-01-29T14:40:33+08:00","IsUpload":true,"UploadID":"a9b7403fce75a7960886d8e14ab10446"}],"msg":"success"}
```