# Monica

Monica 默认使用本机的私钥来验证登录公钥

```bash
ssh-keygen
cat id_rsa.pub >> ~/.ssh/authorized_keys
```

在生成密钥对后, 将本机的公钥加入到已认证的keychain中来保证自动验证