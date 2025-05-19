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

// Load admin dashboard data
async function loadAdminData() {
    try {
        // Load statistics
        const statsResponse = await fetch(`${API_URL}/protected/admin/stats`, {
            headers
        });
        const stats = await statsResponse.json();
        
        document.getElementById('totalUsers').textContent = stats.totalUsers;
        document.getElementById('activeUsers').textContent = stats.activeUsers;
        document.getElementById('totalRoles').textContent = stats.totalRoles;

        // Load users
        const usersResponse = await fetch(`${API_URL}/protected/admin/users`, {
            headers
        });
        const users = await usersResponse.json();
        
        const usersTableBody = document.getElementById('usersTableBody');
        usersTableBody.innerHTML = users.map(user => `
            <tr>
                <td>${user.username}</td>
                <td>${user.email}</td>
                <td>${user.role.name}</td>
                <td>
                    <div class="action-buttons">
                        <button onclick="editUser('${user.id}')">Edit</button>
                        <button class="delete-button" onclick="deleteUser('${user.id}')">Delete</button>
                    </div>
                </td>
            </tr>
        `).join('');

        // Load roles
        const rolesResponse = await fetch(`${API_URL}/protected/admin/roles`, {
            headers
        });
        const roles = await rolesResponse.json();
        
        const rolesTableBody = document.getElementById('rolesTableBody');
        rolesTableBody.innerHTML = roles.map(role => `
            <tr>
                <td>${role.name}</td>
                <td>${role.usersCount}</td>
                <td>
                    <div class="action-buttons">
                        <button onclick="editRole('${role.id}')">Edit</button>
                        <button class="delete-button" onclick="deleteRole('${role.id}')">Delete</button>
                    </div>
                </td>
            </tr>
        `).join('');
    } catch (error) {
        console.error('Error loading admin data:', error);
        alert('Error loading dashboard data');
    }
}

// Load user dashboard data
function loadUserData() {
    document.getElementById('profileUsername').textContent = userData.username;
    document.getElementById('profileEmail').textContent = userData.email;
    document.getElementById('profileRole').textContent = userData.role;
}

// Handle password change
async function handleChangePassword(event) {
    event.preventDefault();
    
    const currentPassword = document.getElementById('currentPassword').value;
    const newPassword = document.getElementById('newPassword').value;
    const confirmPassword = document.getElementById('confirmPassword').value;

    if (newPassword !== confirmPassword) {
        alert('New passwords do not match');
        return;
    }

    try {
        const response = await fetch(`${API_URL}/protected/user/change-password`, {
            method: 'POST',
            headers,
            body: JSON.stringify({
                currentPassword,
                newPassword
            })
        });

        const data = await response.json();
        if (response.ok) {
            alert('Password changed successfully');
            document.getElementById('changePasswordForm').reset();
        } else {
            alert(data.error || 'Failed to change password');
        }
    } catch (error) {
        console.error('Error changing password:', error);
        alert('Error changing password');
    }
}

// Admin functions
async function editUser(userId) {
    // Implement user editing functionality
}

async function deleteUser(userId) {
    if (!confirm('Are you sure you want to delete this user?')) return;

    try {
        const response = await fetch(`${API_URL}/protected/admin/users/${userId}`, {
            method: 'DELETE',
            headers
        });

        if (response.ok) {
            loadAdminData(); // Refresh the data
            alert('User deleted successfully');
        } else {
            const data = await response.json();
            alert(data.error || 'Failed to delete user');
        }
    } catch (error) {
        console.error('Error deleting user:', error);
        alert('Error deleting user');
    }
}

async function editRole(roleId) {
    // Implement role editing functionality
}

async function deleteRole(roleId) {
    if (!confirm('Are you sure you want to delete this role?')) return;

    try {
        const response = await fetch(`${API_URL}/protected/admin/roles/${roleId}`, {
            method: 'DELETE',
            headers
        });

        if (response.ok) {
            loadAdminData(); // Refresh the data
            alert('Role deleted successfully');
        } else {
            const data = await response.json();
            alert(data.error || 'Failed to delete role');
        }
    } catch (error) {
        console.error('Error deleting role:', error);
        alert('Error deleting role');
    }
}

