# cmd to run minIO

```bash
podman run -p 9000:9000 -p 9001:9001 -e MINIO_ROOT_USER=admin -e MINIO_ROOT_PASSWORD=Asd424941_a quay.io/minio/minio server /data --console-address ":9001"
```

note: user and pass is the access_key + secret