function getRosePrice() {
    const date = document.getElementById('date').value;
    const resultDiv = document.getElementById('result');
    const loadingDiv = document.getElementById('loading');

    // 验证日期
    if (!date) {
        alert("請選擇日期！");
        return;
    }

    // 显示加载提示
    loadingDiv.style.display = "block";
    resultDiv.innerHTML = '';

    // 发送请求
    fetch(`http://localhost:8080/api/rose-price?date=${date}`)
        .then(response => {
            console.log("收到回應:", response); // 添加日誌
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log("解析 JSON 成功:", data); // 添加日誌
            loadingDiv.style.display = "none";
            if (data.success) {
                resultDiv.innerHTML = `玫瑰價格：${data.price}`;
            } else {
                resultDiv.innerHTML = `沒有該日期的玫瑰價格。`;
            }
        })
        .catch(error => {
            console.error("錯誤:", error); // 添加日誌
            loadingDiv.style.display = "none";
            resultDiv.innerHTML = `發生錯誤，請稍後再試。 (${error.message})`;
        });
}
