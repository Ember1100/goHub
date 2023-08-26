function uploadFile(file) {
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/presignedurl?objectName=' + file.name, true);

  xhr.onload = function() {
    if (xhr.status === 200) {
      var presignedURL = xhr.responseText;
      uploadToMinIO(file, presignedURL);
    } else {
      console.error('无法获取临时签名 URL');
    }
  };

  xhr.send();
}

function uploadToMinIO(file, presignedURL) {
  var xhr = new XMLHttpRequest();
  xhr.open('PUT', presignedURL, true);

  xhr.onload = function() {
    if (xhr.status === 200) {
      console.log('文件上传成功');
    } else {
      console.error('文件上传失败');
    }
  };

  xhr.send(file);
}