document.addEventListener('DOMContentLoaded', () => {
  const btn = document.getElementById('convert-btn');

  btn.addEventListener('click', function() {
    const temperatureInput = document.getElementById('temperature-input');
    const convertType = document.getElementById('conversion-type');
    const temperature = parseFloat(temperatureInput.value);

    if (isNaN(temperature)) {
      document.getElementById('result').innerText = '❌ Please enter a valid number.';
      return;
    }

    let result;

    if (convertType.value === "c-to-f") {
      result = (temperature * 9 / 5) + 32;
      document.getElementById('result').innerText = `🌡️ ${temperature}°C = ${result.toFixed(2)}°F`;
    } else if (convertType.value === "f-to-c") {
      result = (temperature - 32) * 5 / 9;
      document.getElementById('result').innerText = `🌡️ ${temperature}°F = ${result.toFixed(2)}°C`;
    } else {
      document.getElementById('result').innerText = '❌ Invalid conversion type';
    }
  });
});

