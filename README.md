# cfztna-tls-beacon

## How

1. Generate a self-signed certificate

    ```shell
    openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout tls.key -out tls.crt -subj "/CN=cfztna-tls-beacon" -addext "subjectAltName=DNS:cfztna-tls-beacon"
    ```
   
1. Get SHA256 fingerprint of the certificate

    ```shell
    openssl x509 -noout -fingerprint -sha256 -inform pem -in tls.crt | tr -d :
    ```

1. Start beacon

    ```shell
    docker run -d --name cfztna-tls-beacon -p 4443:4443 \
      -e TLS_CERT=xxx \
      -e TLS_KEY=xxx \
      -e PORT=4443 \
      gcr.io/moonrhythm-containers/cfztna-tls-beacon
    ```
