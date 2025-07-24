// Helper function to render user list
export function renderUserList(users, container) {
    // Clear container
    container.innerHTML = '';

    if (!users || users.length === 0) {
        container.innerHTML = '<p>No users found.</p>';
        return;
    }

    // Add users
    users.forEach(user => {
        const userCard = document.createElement('div');
        userCard.className = 'user-card';
        userCard.innerHTML = `
        <h4>${user.name}</h4>
        <p><strong>Daily Calorie Target:</strong> ${user.calories}</p>
        <div class="actions">
        <button class="delete-btn" data-id="${user.id}">Delete</button>
        </div>
        `;

        // Add delete event listener
        const deleteBtn = userCard.querySelector('.delete-btn');
        deleteBtn.addEventListener('click', async () => {
            if (confirm(`Are you sure you want to delete ${user.name}?`)) {
                await deleteDieter(user);
                userCard.remove();
                // Remove from allUsers array
                const index = allUsers.findIndex(u => u.id === user.id);
                if (index !== -1) {
                    allUsers.splice(index, 1);
                }
            }
        });

        container.appendChild(userCard);
    });
}

// API Functions
export async function addDieter(name, calories) {
    const userListContainer = document.getElementById('user-list-container');
    try {
        // Add User
        const addUserResponse = await fetch(`${API_BASE_URL}/dieters`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name: name, calories: calories })
        });
        showSuccess('User added successfully!');
        allUsers.push(await addUserResponse.json());
        renderUserList(allUsers, userListContainer);
        return true;
    } catch (error) {
        console.error('Error adding user:', error);
        showError('Failed to add user. Please try again later.');
        return false;

    }
}

export async function deleteDieter(dieter) {
    try {
        const response = await fetch(`${API_BASE_URL}/dieters`, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(dieter)
        });

        if (!response.ok) throw new Error('Failed to delete user');

        showSuccess('User deleted successfully!');
        return true;
    } catch (error) {
        console.error('Error deleting user:', error);
        showError('Failed to delete user. Please try again later.');
        return false;
    }
}

export async function getDieterMealsToday() {
    return await fetch(`${API_BASE_URL}/dieter/meals`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name: currentUser.name })
    })
}

// Helper function to populate user select
export async function populateUserSelect(selectElement) {
    // Clear existing options
    selectElement.innerHTML = '';

    // Add users
    allUsers.forEach(user => {
        const option = document.createElement('option');
        option.value = user.id;
        option.textContent = user.name;
        selectElement.appendChild(option);
    });
}
