# PDFTK as a service

## Image

`docker pull katsick/pdftkgolangservice:1.0.2`

## Hot to use

Invoke HTTP request to the service. Example below:

```
POST /fill-pdf HTTP/1.1
Content-Type: multipart/form-data;

--1
Content-Disposition: form-data; name="file"; filename="form.pdf"
Content-Type: application/pdf

<BYTES HERE>
--1

--2
Content-Disposition: form-data; name="json"
{"field1":"hello","field2":"world"}
--2
```

## Development

```
docker-compose up --build
```

Use Postman or Paw to invoke requests agains HTTP server
