import './style.css';
import './app.css';

import { GenerateKey, Encrypt, Decrypt } from '../wailsjs/go/main/App';

// Generate Key Button
document.getElementById('genKeyBtn').addEventListener('click', async () => {
    try {
        clearError();
        const key = await GenerateKey();
        const keyOutput = document.getElementById('generatedKey');
        keyOutput.innerHTML = `<strong>Generated Key:</strong><br><code>${key}</code>`;
        document.getElementById('keyInput').value = key;
    } catch (err) {
        showError('Error generating key: ' + err);
    }
});

// Encrypt Button
document.getElementById('encryptBtn').addEventListener('click', async () => {
    try {
        clearError();
        const key = document.getElementById('keyInput').value;
        const plaintext = document.getElementById('plaintext').value;

        if (!key) {
            showError('Please enter an encryption key');
            return;
        }
        if (!plaintext) {
            showError('Please enter text to encrypt');
            return;
        }

        const encrypted = await Encrypt(key, plaintext);
        const output = document.getElementById('encryptedOutput');
        output.innerHTML = `<strong>Encrypted (Base64):</strong><br><code>${encrypted}</code>`;

        // Auto-fill the decrypt input
        document.getElementById('encryptedInput').value = encrypted;
    } catch (err) {
        showError('Encryption error: ' + err);
    }
});

// Decrypt Button
document.getElementById('decryptBtn').addEventListener('click', async () => {
    try {
        clearError();
        const key = document.getElementById('keyInput').value;
        const encrypted = document.getElementById('encryptedInput').value;

        if (!key) {
            showError('Please enter a decryption key');
            return;
        }
        if (!encrypted) {
            showError('Please enter encrypted text to decrypt');
            return;
        }

        const decrypted = await Decrypt(key, encrypted);
        const output = document.getElementById('decryptedOutput');
        output.innerHTML = `<strong>Decrypted:</strong><br><code>${decrypted}</code>`;
    } catch (err) {
        showError('Decryption error: ' + err);
    }
});

function showError(message) {
    const errorDiv = document.getElementById('errorMsg');
    errorDiv.textContent = message;
    errorDiv.style.display = 'block';
}

function clearError() {
    const errorDiv = document.getElementById('errorMsg');
    errorDiv.textContent = '';
    errorDiv.style.display = 'none';
}
