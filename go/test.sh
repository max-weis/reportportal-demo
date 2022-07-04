go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml
zip report.zip report.xml

curl -X POST "http://localhost:8080/api/v1/default_personal/launch/import" \
    -H "accept: */*" -H "Content-Type: multipart/form-data" \
    -H "Authorization: bearer b98a4a01-4b54-4383-8fae-d2d6db0f336e" \
    -F "file=@report.zip;type=application/zip"