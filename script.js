function getRosePrice() {
    const date = document.getElementById('date').value;
    const resultDiv = document.getElementById('result');
    const loadingDiv = document.getElementById('loading');

    // Validate the date
    if (!date) {
        alert("請選擇日期！");
        return;
    }

    // Show loading indicator
    loadingDiv.style.display = "block";
    resultDiv.innerHTML = '';

    fetch('https://cbf6-2407-4d00-1c03-752a-4878-922e-97bf-a371.ngrok-free.app/api/rose-price?date=2025-01-08', {
        headers: {
            'Accept': 'application/json',
            'ngrok-skip-browser-warning': 'true'
        }
    })
      //.then(response => response.json())
      //.then(data => console.log(data));
    .then(response => {
        console.log("收到回應:", response); // 調試輸出請求狀態
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
    })
    .then(data => {
        console.log("收到的數據:", data); // 調試輸出後端回應的數據
        loadingDiv.style.display = "none";
        if (data.success) {
            resultDiv.innerHTML = `玫瑰價格：${data.price}`;
        } else {
            resultDiv.innerHTML = `沒有該日期的玫瑰價格。`;
        }
    })
    .catch(error => {
        console.error("錯誤:", error); // 捕獲並打印錯誤
        loadingDiv.style.display = "none";
        resultDiv.innerHTML = `發生錯誤，請稍後再試。 (${error.message})`;
        });
    }
