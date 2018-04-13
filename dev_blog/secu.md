https://github.com/unrolled/secure

https://blog.rapid7.com/2016/07/13/quick-security-wins-in-golang/

https://blog.rapid7.com/2017/08/28/rsa-rivest-shamir-and-adleman/

https://blog.rapid7.com/2017/08/28/des-data-encryption-standard/

https://blog.rapid7.com/2017/07/28/exploring-sha-1-secure-hash-algorithm/

## 第一步

先在某个位置生私钥跟公钥

开发时可以放项目根目录 生产环境的位置：
the cert in something like /root/certs 
and keys in something like /etc/ssl/private.

~~~shell

# 生私钥：
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

# 生x509自验证公钥： （ generate the x509 self-signed public key: ） 
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

~~~