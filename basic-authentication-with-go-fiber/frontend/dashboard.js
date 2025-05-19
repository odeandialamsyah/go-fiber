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

// Initialize dashboard based on user role
function initializeDashboard() {
    // Set username in sidebar
    document.getElementById('username').textContent = userData.username;

    // Show appropriate sections based on role
    if (userData.role === 'admin') {
        document.getElementById('adminLinks').style.display = 'block';
        document.getElementById('adminSections').style.display = 'block';
        loadAdminData();
    } else {
        document.getElementById('userLinks').style.display = 'block';
        document.getElementById('userSections').style.display = 'block';
        loadUserData();
    }

    // Add click handlers for navigation
    document.querySelectorAll('.nav-links a[data-section]').forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const section = e.target.getAttribute('data-section');
            showSection(section);
        });
    });
}

// Show selected section and hide others
function showSection(sectionId) {
    // Get the container based on user role
    const container = userData.role === 'admin' ? 'adminSections' : 'userSections';
    
    // Hide all sections in the container
    document.querySelectorAll(`#${container} .dashboard-section`).forEach(section => {
        section.classList.remove('active');
    });
    
    // Show selected section
    document.getElementById(sectionId).classList.add('active');

    // Update active nav link
    const navContainer = userData.role === 'admin' ? 'adminLinks' : 'userLinks';
    document.querySelectorAll(`#${navContainer} li`).forEach(li => {
        li.classList.remove('active');
    });
    document.querySelector(`#${navContainer} a[data-section="${sectionId}"]`).parentElement.classList.add('active');
}
