# PDFTK as a service

## Image

`docker pull katsick/pdftkgolangservice:1.1.0`

## How to use

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

### Flatten option

By default, the service flattens the PDF file (making a form non-editable). If you do 
not want the service to flatten the file, provide the query param `flatten=false` to the `fill-pdf`
endpoint. i.e. `POST /fill-pdf?flatten=false`

## Development

```
docker-compose up --build
```

Use Postman or Paw to invoke requests agains HTTP server
