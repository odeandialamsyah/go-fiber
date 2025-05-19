const API_URL = 'http://localhost:3000/api';

// Check if user is logged in
const token = localStorage.getItem('token');
if (!token) {
    window.location.href = '/index.html';
}

// Get user data from localStorage
const userData = JSON.parse(localStorage.getItem('userData'));
if (!userData) {
    logout();
}

// Set up headers for API requests
const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${token}`
};
