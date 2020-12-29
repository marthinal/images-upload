Golang multipart file upload example.

Some JS to upload a file:
```javascript
        let data = new FormData();
        data.append("images", image[0]);
        fetch('http://127.0.0.1:6090/images', {method: "POST", body: data})
            .then(async response => {
                console.log(response);
        });
```