const API_BASE_URL = window.location.hostname === 'localhost' ? 'http://localhost:8080' : 'http://localhost:8080';

function formatResult(data) {
    if (!data.open_ports || data.open_ports.length === 0) {
        return `Хост: ${data.host}\nОткрытые порты: не найдены`;
    }
    
    return `Хост: ${data.host}\nОткрытые порты (${data.open_ports.length}): ${data.open_ports.join(', ')}`;
}

function displayPortsAsChips(data, resultElem) {
    resultElem.innerHTML = '';
    
    const hostInfo = document.createElement('div');
    hostInfo.className = 'host-info';
    hostInfo.textContent = `Хост: ${data.host}`;
    resultElem.appendChild(hostInfo);
    
    if (!data.open_ports || data.open_ports.length === 0) {
        const noPortsMsg = document.createElement('div');
        noPortsMsg.textContent = 'Открытые порты не найдены';
        noPortsMsg.style.color = '#888';
        resultElem.appendChild(noPortsMsg);
        return;
    }
    
    const portsCountMsg = document.createElement('div');
    portsCountMsg.textContent = `Найдено открытых портов: ${data.open_ports.length}`;
    portsCountMsg.style.marginBottom = '10px';
    portsCountMsg.style.fontWeight = 'bold';
    resultElem.appendChild(portsCountMsg);
    
    const portsList = document.createElement('div');
    portsList.className = 'port-list';
    
    data.open_ports.forEach(port => {
        const portItem = document.createElement('span');
        portItem.className = 'port-item';
        portItem.textContent = port;
        portsList.appendChild(portItem);
    });
    
    resultElem.appendChild(portsList);
}

function validateHost(host) {
    if (!host || host.trim() === '') {
        return 'Хост обязателен для заполнения';
    }
    
    // Удаляем протокол если есть
    host = host.replace(/^https?:\/\//, '');
    
    // Проверяем IP адрес
    const ipRegex = /^(\d{1,3}\.){3}\d{1,3}$/;
    if (ipRegex.test(host)) {
        const parts = host.split('.');
        for (let part of parts) {
            if (parseInt(part) > 255) {
                return 'Неверный IP адрес';
            }
        }
        return null;
    }
    
    // Проверяем доменное имя
    const domainRegex = /^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
    if (!domainRegex.test(host)) {
        return 'Неверный формат хоста';
    }
    
    return null;
}

async function scanPorts() {
    const hostInput = document.getElementById("host");
    const resultElem = document.getElementById("result");
    const spinner = document.getElementById("spinner");
    const scanButton = document.querySelector("button");
    
    const host = hostInput.value.trim();
    
    // Валидация
    const validationError = validateHost(host);
    if (validationError) {
        resultElem.textContent = `Ошибка: ${validationError}`;
        resultElem.className = 'result error';
        return;
    }
    
    // Очищаем результат и показываем спиннер
    resultElem.textContent = '';
    resultElem.className = 'result';
    spinner.style.display = 'block';
    scanButton.disabled = true;
    scanButton.textContent = 'Сканирование...';
    
    try {
        const cleanHost = host.replace(/^https?:\/\//, '');
        
        const response = await fetch(`${API_BASE_URL}/scan`, {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: JSON.stringify({ host: cleanHost })
        });
        
        if (!response.ok) {
            let errorMessage = `HTTP ошибка! Статус: ${response.status}`;
            try {
                const errorData = await response.json();
                if (errorData.error) {
                    errorMessage = errorData.error;
                }
            } catch (e) {
                // Игнорируем ошибку парсинга JSON для error response
            }
            throw new Error(errorMessage);
        }
        
        const data = await response.json();
        displayPortsAsChips(data, resultElem);
        resultElem.className = 'result success';
        
    } catch (error) {
        console.error('Scan error:', error);
        resultElem.textContent = `Ошибка: ${error.message}`;
        resultElem.className = 'result error';
    } finally {
        spinner.style.display = 'none';
        scanButton.disabled = false;
        scanButton.textContent = 'Сканировать порты';
    }
}

// Обработчик Enter в поле ввода
document.addEventListener('DOMContentLoaded', function() {
    const hostInput = document.getElementById("host");
    if (hostInput) {
        hostInput.addEventListener('keypress', function(e) {
            if (e.key === 'Enter') {
                scanPorts();
            }
        });
    }
});

// Проверка доступности API при загрузке страницы
async function checkApiHealth() {
    try {
        const response = await fetch(`${API_BASE_URL}/health`, {
            method: 'GET',
            timeout: 5000
        });
        console.log('API доступен');
    } catch (error) {
        console.warn('API может быть недоступен:', error.message);
        const resultElem = document.getElementById("result");
        if (resultElem) {
            resultElem.textContent = 'Предупреждение: API сервер может быть недоступен';
            resultElem.className = 'result error';
        }
    }
}

// Проверяем API при загрузке страницы
document.addEventListener('DOMContentLoaded', checkApiHealth);
