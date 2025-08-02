import * as utils from './utils.js'

// Helper function to render user list
export function renderUserList(app, API_BASE_URL, users, container) {
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
        await deleteDieter(app, API_BASE_URL, user);
        userCard.remove();
        // Remove from allUsers array
        const index = users.findIndex(u => u.id === user.id);
        if (index !== -1) {
          users.splice(index, 1);
        }
      }
    });

    container.appendChild(userCard);
  });
}

// API Functions
export async function addDieter(app, API_BASE_URL, allUsers, name, calories) {
  const userListContainer = document.getElementById('user-list-container');
  try {
    // Add User
    const addUserResponse = await fetch(`${API_BASE_URL}/dieters`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name: name, calories: calories })
    });
    utils.showSuccess(app, 'User added successfully!');
    allUsers.push(await addUserResponse.json());
    renderUserList(app, API_BASE_URL, allUsers, userListContainer);
    return true;
  } catch (error) {
    console.error('Error adding user:', error);
    utils.showError(app, 'Failed to add user. Please try again later.');
    return false;

  }
}

export async function deleteDieter(app, API_BASE_URL, dieter) {
  try {
    const response = await fetch(`${API_BASE_URL}/dieters`, {
      method: 'DELETE',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dieter)
    });

    if (!response.ok) throw new Error('Failed to delete user');

    utils.showSuccess(app, 'User deleted successfully!');
    return true;
  } catch (error) {
    console.error('Error deleting user:', error);
    utils.showError(app, 'Failed to delete user. Please try again later.');
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
export async function populateUserSelect(allUsers, selectElement) {
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
