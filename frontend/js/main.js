// API Base URL
const API_BASE_URL = 'http://localhost:9090';

// Current state
let currentPage = 'dashboard';
let currentUser = null;
let allUsers = [];
let allFoods = [];

// DOM Elements
const app = document.getElementById('app');
const loading = document.getElementById('loading');
const navLinks = document.querySelectorAll('nav a');

// Initialize the application
document.addEventListener('DOMContentLoaded', () => {
    // Set up navigation
    navLinks.forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const page = link.getAttribute('data-page');
            navigateTo(page);
        });
    });

    // Load initial data and render dashboard
    loadInitialData();
});

// Load initial data from API
async function loadInitialData() {
    showLoading();
    try {
        // Load all users
        const usersResponse = await fetch(`${API_BASE_URL}/dieters/all`);
        if (!usersResponse.ok) throw new Error('Failed to load users');
        allUsers = await usersResponse.json();

        // Load all foods
        const foodsResponse = await fetch(`${API_BASE_URL}/food/all`);
        if (!foodsResponse.ok) throw new Error('Failed to load foods');
        allFoods = await foodsResponse.json();

        // Render initial page
        navigateTo(currentPage);
    } catch (error) {
        console.error('Error loading initial data:', error);
        showError('Failed to load initial data. Please try again later.');
    } finally {
        hideLoading();
    }
}

// Navigation function
function navigateTo(page) {
    // Update active nav link
    navLinks.forEach(link => {
        if (link.getAttribute('data-page') === page) {
            link.classList.add('active');
        } else {
            link.classList.remove('active');
        }
    });

    // Update current page
    currentPage = page;

    // Render the page
    renderPage(page);
}

// Render the current page
function renderPage(page) {
    showLoading();

    // Clear the app container
    while (app.firstChild) {
        if (app.firstChild !== loading) {
            app.removeChild(app.firstChild);
        }
    }

    // Get the template
    const template = document.getElementById(`${page}-template`);
    if (!template) {
        hideLoading();
        return showError(`Template for ${page} not found`);
    }

    // Clone the template content
    const content = template.content.cloneNode(true);
    app.appendChild(content);

    // Initialize page-specific functionality
    switch (page) {
        case 'dashboard':
            initDashboard();
            break;
        case 'meals':
            initMeals();
            break;
        case 'foods':
            initFoods();
            break;
        case 'users':
            initUsers();
            break;
    }

    hideLoading();
}

// Initialize Dashboard
function initDashboard() {
    const userSelect = document.getElementById('user-select');
    const calorieSummary = document.getElementById('calorie-summary');
    const todayMeals = document.getElementById('today-meals');
    const userNameSpan = document.getElementById('dashboard-user-name');
    const dailyTarget = document.getElementById('daily-target');
    const consumedToday = document.getElementById('consumed-today');
    const remaining = document.getElementById('remaining');
    const mealsList = document.getElementById('meals-list');

    // Populate user select
    populateUserSelect(userSelect);

    // Handle user selection
    userSelect.addEventListener('change', async () => {
        const selectedUserId = userSelect.value;
        if (!selectedUserId) {
            calorieSummary.style.display = 'none';
            todayMeals.style.display = 'none';
            return;
        }

        showLoading();
        try {
            // Find selected user
            const selectedUser = allUsers.find(user => user.id.toString() === selectedUserId);
            currentUser = selectedUser;

            // Update user name
            userNameSpan.textContent = selectedUser.name;

            // Get daily target
            dailyTarget.textContent = selectedUser.calories;

            // Get remaining calories
            const remainingResponse = await fetch(`${API_BASE_URL}/dieter/remaining`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: selectedUser.name })
            });
            
            if (!remainingResponse.ok) throw new Error('Failed to load remaining calories');
            const remainingData = await remainingResponse.json();
            remaining.textContent = remainingData.calories;

            // Calculate consumed calories
            const consumed = selectedUser.calories - remainingData.calories;
            consumedToday.textContent = consumed;

            // Get today's meals
            const mealsResponse = await fetch(`${API_BASE_URL}/dieter/mealstoday`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: selectedUser.name })
            });
            
            if (!mealsResponse.ok) throw new Error('Failed to load today\'s meals');
            const mealsData = await mealsResponse.json();

            // Render meals
            renderMealsList(mealsData, mealsList);

            // Show summary and meals
            calorieSummary.style.display = 'flex';
            todayMeals.style.display = 'block';
        } catch (error) {
            console.error('Error loading dashboard data:', error);
            showError('Failed to load dashboard data. Please try again later.');
        } finally {
            hideLoading();
        }
    });
}

// Initialize Meals page
function initMeals() {
    const userSelect = document.getElementById('meal-user-select');
    const mealManagement = document.getElementById('meal-management');
    const addMealForm = document.getElementById('add-meal-form');
    const mealFoodsSelect = document.getElementById('meal-foods');
    const mealHistoryList = document.getElementById('meal-history-list');

    // Populate user select
    populateUserSelect(userSelect);

    // Populate foods select
    allFoods.forEach(food => {
        const option = document.createElement('option');
        option.value = food.id;
        option.textContent = `${food.name} (${food.calories} cal)`;
        mealFoodsSelect.appendChild(option);
    });

    // Handle user selection
    userSelect.addEventListener('change', async () => {
        const selectedUserId = userSelect.value;
        if (!selectedUserId) {
            mealManagement.style.display = 'none';
            return;
        }

        showLoading();
        try {
            // Find selected user
            const selectedUser = allUsers.find(user => user.id.toString() === selectedUserId);
            currentUser = selectedUser;

            // Get user's meals
            const mealsResponse = await fetch(`${API_BASE_URL}/dieter/meals`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: selectedUser.name })
            });
            
            if (!mealsResponse.ok) throw new Error('Failed to load meals');
            const mealsData = await mealsResponse.json();

            // Render meal history
            renderMealHistory(mealsData, mealHistoryList);

            // Show meal management
            mealManagement.style.display = 'block';
        } catch (error) {
            console.error('Error loading meals data:', error);
            showError('Failed to load meals data. Please try again later.');
        } finally {
            hideLoading();
        }
    });

    // Handle add meal form submission
    addMealForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        if (!currentUser) return;

        const mealName = document.getElementById('meal-name').value;
        const mealCalories = document.getElementById('meal-calories').value;
        
        // Get selected foods
        const selectedFoods = Array.from(mealFoodsSelect.selectedOptions).map(option => {
            const foodId = option.value;
            const food = allFoods.find(f => f.id.toString() === foodId);
            return food.name;
        });

        showLoading();
        try {
            // Add meal
            const response = await fetch(`${API_BASE_URL}/meal`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: mealName,
                    dieter: currentUser.name,
                    calories: parseInt(mealCalories),
                    food: selectedFoods
                })
            });
            
            if (!response.ok) throw new Error('Failed to add meal');
            
            // Reset form
            addMealForm.reset();
            
            // Refresh meals
            userSelect.dispatchEvent(new Event('change'));
            
            showSuccess('Meal added successfully!');
        } catch (error) {
            console.error('Error adding meal:', error);
            showError('Failed to add meal. Please try again later.');
        } finally {
            hideLoading();
        }
    });
}

// Initialize Foods page
function initFoods() {
    const addFoodForm = document.getElementById('add-food-form');
    const foodListContainer = document.getElementById('food-list-container');

    // Render food list
    renderFoodList(allFoods, foodListContainer);

    // Handle add food form submission
    addFoodForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const foodName = document.getElementById('food-name').value;
        const foodCalories = document.getElementById('food-calories').value;
        const foodUnits = document.getElementById('food-units').value;

        showLoading();
        try {
            // Add food
            const response = await fetch(`${API_BASE_URL}/food`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: foodName,
                    calories: parseInt(foodCalories),
                    units: parseInt(foodUnits)
                })
            });
            
            if (!response.ok) throw new Error('Failed to add food');
            const newFood = await response.json();
            
            // Add to local array
            allFoods.push(newFood);
            
            // Reset form
            addFoodForm.reset();
            
            // Refresh food list
            renderFoodList(allFoods, foodListContainer);
            
            showSuccess('Food added successfully!');
        } catch (error) {
            console.error('Error adding food:', error);
            showError('Failed to add food. Please try again later.');
        } finally {
            hideLoading();
        }
    });
}

// Initialize Users page
function initUsers() {
    const addUserForm = document.getElementById('add-user-form');
    const userListContainer = document.getElementById('user-list-container');

    // Render user list
    renderUserList(allUsers, userListContainer);

    // Handle add user form submission
    addUserForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const userName = document.getElementById('new-user-name').value;
        const userCalories = document.getElementById('user-calories').value;

        showLoading();
        try {
            // Add user
            const response = await fetch(`${API_BASE_URL}/dieters`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: userName,
                    calories: parseInt(userCalories)
                })
            });
            
            if (!response.ok) throw new Error('Failed to add user');
            const newUser = await response.json();
            
            // Add to local array
            allUsers.push(newUser);
            
            // Reset form
            addUserForm.reset();
            
            // Refresh user list
            renderUserList(allUsers, userListContainer);
            
            showSuccess('User added successfully!');
        } catch (error) {
            console.error('Error adding user:', error);
            showError('Failed to add user. Please try again later.');
        } finally {
            hideLoading();
        }
    });
}

// Helper function to populate user select
function populateUserSelect(selectElement) {
    // Clear existing options
    while (selectElement.options.length > 1) {
        selectElement.remove(1);
    }

    // Add users
    allUsers.forEach(user => {
        const option = document.createElement('option');
        option.value = user.id;
        option.textContent = user.name;
        selectElement.appendChild(option);
    });
}

// Helper function to render meals list
function renderMealsList(meals, container) {
    // Clear container
    container.innerHTML = '';

    if (!meals || meals.length === 0) {
        container.innerHTML = '<p>No meals found for today.</p>';
        return;
    }

    // Add meals
    meals.forEach(meal => {
        const mealCard = document.createElement('div');
        mealCard.className = 'meal-card';
        mealCard.innerHTML = `
            <h4>${meal.name}</h4>
            <p><strong>Calories:</strong> ${meal.calories}</p>
            <div class="actions">
                <button class="delete-btn" data-id="${meal.id}">Delete</button>
            </div>
        `;

        // Add delete event listener
        const deleteBtn = mealCard.querySelector('.delete-btn');
        deleteBtn.addEventListener('click', async () => {
            if (confirm(`Are you sure you want to delete ${meal.name}?`)) {
                await deleteMeal(meal);
                mealCard.remove();
            }
        });

        container.appendChild(mealCard);
    });
}

// Helper function to render meal history
function renderMealHistory(meals, container) {
    // Clear container
    container.innerHTML = '';

    if (!meals || meals.length === 0) {
        container.innerHTML = '<p>No meal history found.</p>';
        return;
    }

    // Group meals by day
    const mealsByDay = {};
    meals.forEach(meal => {
        if (!mealsByDay[meal.day]) {
            mealsByDay[meal.day] = [];
        }
        mealsByDay[meal.day].push(meal);
    });

    // Add meals grouped by day
    Object.keys(mealsByDay).sort().reverse().forEach(day => {
        const dayHeader = document.createElement('h4');
        dayHeader.textContent = day;
        container.appendChild(dayHeader);

        mealsByDay[day].forEach(meal => {
            const mealCard = document.createElement('div');
            mealCard.className = 'meal-card';
            mealCard.innerHTML = `
                <h4>${meal.name}</h4>
                <p><strong>Calories:</strong> ${meal.calories}</p>
                <div class="actions">
                    <button class="delete-btn" data-id="${meal.id}">Delete</button>
                </div>
            `;

            // Add delete event listener
            const deleteBtn = mealCard.querySelector('.delete-btn');
            deleteBtn.addEventListener('click', async () => {
                if (confirm(`Are you sure you want to delete ${meal.name}?`)) {
                    await deleteMeal(meal);
                    mealCard.remove();
                }
            });

            container.appendChild(mealCard);
        });
    });
}

// Helper function to render food list
function renderFoodList(foods, container) {
    // Clear container
    container.innerHTML = '';

    if (!foods || foods.length === 0) {
        container.innerHTML = '<p>No foods found.</p>';
        return;
    }

    // Add foods
    foods.forEach(food => {
        const foodCard = document.createElement('div');
        foodCard.className = 'food-card';
        foodCard.innerHTML = `
            <h4>${food.name}</h4>
            <p><strong>Calories:</strong> ${food.calories}</p>
            <p><strong>Units:</strong> ${food.units}</p>
            <div class="actions">
                <button class="delete-btn" data-id="${food.id}">Delete</button>
            </div>
        `;

        // Add delete event listener
        const deleteBtn = foodCard.querySelector('.delete-btn');
        deleteBtn.addEventListener('click', async () => {
            if (confirm(`Are you sure you want to delete ${food.name}?`)) {
                await deleteFood(food);
                foodCard.remove();
                // Remove from allFoods array
                const index = allFoods.findIndex(f => f.id === food.id);
                if (index !== -1) {
                    allFoods.splice(index, 1);
                }
            }
        });

        container.appendChild(foodCard);
    });
}

// Helper function to render user list
function renderUserList(users, container) {
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
async function deleteMeal(meal) {
    try {
        const response = await fetch(`${API_BASE_URL}/meal`, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(meal)
        });
        
        if (!response.ok) throw new Error('Failed to delete meal');
        
        showSuccess('Meal deleted successfully!');
        return true;
    } catch (error) {
        console.error('Error deleting meal:', error);
        showError('Failed to delete meal. Please try again later.');
        return false;
    }
}

async function deleteFood(food) {
    try {
        const response = await fetch(`${API_BASE_URL}/food`, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(food)
        });
        
        if (!response.ok) throw new Error('Failed to delete food');
        
        showSuccess('Food deleted successfully!');
        return true;
    } catch (error) {
        console.error('Error deleting food:', error);
        showError('Failed to delete food. Please try again later.');
        return false;
    }
}

async function deleteDieter(dieter) {
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

// UI Helper Functions
function showLoading() {
    loading.style.display = 'block';
}

function hideLoading() {
    loading.style.display = 'none';
}

function showError(message) {
    // Create error message element
    const errorElement = document.createElement('div');
    errorElement.className = 'error-message';
    errorElement.textContent = message;
    
    // Add to app
    app.appendChild(errorElement);
    
    // Remove after 5 seconds
    setTimeout(() => {
        errorElement.remove();
    }, 5000);
}

function showSuccess(message) {
    // Create success message element
    const successElement = document.createElement('div');
    successElement.className = 'success-message';
    successElement.textContent = message;
    
    // Add to app
    app.appendChild(successElement);
    
    // Remove after 3 seconds
    setTimeout(() => {
        successElement.remove();
    }, 3000);
}