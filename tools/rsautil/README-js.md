```js
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>RSA Encryption and Decryption</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jsencrypt/3.0.0/jsencrypt.min.js"></script>
</head>
<body>

<script>
    // 使用已有的RSA公钥加密
    function encryptWithPublicKey(publicKey, plaintext) {
        var jsEncrypt = new JSEncrypt({default_key_size: 2048});
        jsEncrypt.setPublicKey(publicKey);
        return jsEncrypt.encrypt(plaintext);
    }

    // 使用已有的RSA私钥解密
    function decryptWithPrivateKey(privateKey, ciphertext) {
        var jsEncrypt = new JSEncrypt({default_key_size: 2048});
        jsEncrypt.setPrivateKey(privateKey);
        return jsEncrypt.decrypt(ciphertext);
    }

    // 示例
    var publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC4yVFguN2rPQ9Uw3p6KivP+xCw
7p+RR/In0Lqohf6iPD7+zbuwrMHxypxjcWcFQaWBCcP2yKogT6q7LWfC4wtOwgq1
tqREqHbfkn+i3G/8W14LDCp97kq4WtxKtVdM9F5SaD/spXX/ThgVhOskq8YjcI4Y
4gTQ83E+fEbaQvCfKwIDAQAB
-----END PUBLIC KEY-----`;
    var privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC4yVFguN2rPQ9Uw3p6KivP+xCw7p+RR/In0Lqohf6iPD7+zbuw
rMHxypxjcWcFQaWBCcP2yKogT6q7LWfC4wtOwgq1tqREqHbfkn+i3G/8W14LDCp9
7kq4WtxKtVdM9F5SaD/spXX/ThgVhOskq8YjcI4Y4gTQ83E+fEbaQvCfKwIDAQAB
AoGBALatZ0LcX2ALBB4DBhDBogCBLpYLwTnpy05rPiyeEwYl0w0pLDTUBQPZDlQM
5xC+PjTcB5vv8qfwulNC5wI2XJS34l8JmYo/XHKvUq72Hk8cduCnT87FYcsIX41D
ihGOIgG0dHQUUEtqzxENI2qCjmglgGHP/+qLMab3EOhNDK7BAkEAzTkOzNAaYdnO
U6O9Azu7V/X8IjIjHv72vfF0Vmv+OiZTJ3WMBwudPWWM1ZGX9cu4pwCGoGyrtZxO
LpcGpJp7CwJBAOaBzPz1aG4ThdlVR8EWnUvpo+mNrXI8ph2Dfi17uQg2A6EZPjZA
DtS4COAzoWpsB5o58exduTmIMjZJrS44AGECQQC8Oo82j9EC2tEBqfbdFlY40WeW
vcG01knd4a7A7YBaOXifgpMSizaHb7MC1+03Bsmwcy0Hy2SayGh1FxSCuSYNAkBY
Qr9A5J7V9ze7HgJZltUn6hBPL2aIZVyd1GmN9N/GmxgMqWO+1gxXuxf68QoPe8n1
bdaKUODJfLLtQozDM8JBAkBBkJtY/h5oBfDYnA0yznDUU2vei3IYd2jVba6zg2Ve
yhK5ju1AZwsRmD/JLmD+Asq9ifi8LR1rgpGj2DEpyEHQ
-----END RSA PRIVATE KEY-----`;

    var message = 'Hello, RSA encryption and decryption! 你好，RSA 加解密！';

    // 加密
    var encryptedData = encryptWithPublicKey(publicKey, message);
    console.log('Encrypted data:', encryptedData);

    // 解密
    var decryptedData = decryptWithPrivateKey(privateKey, encryptedData);
    console.log('Decrypted data:', decryptedData);
</script>

</body>
</html>


```