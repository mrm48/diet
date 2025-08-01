import * as user from './userFunctions.js'
import * as meal from './mealFunctions.js'
import * as food from './foodFunctions.js'
import * as entry from './entryFunctions.js'
// API Base URL
const API_BASE_URL = 'http://localhost:9090';



// Current state
let currentPage = 'dashboard';
let currentUser = 'Matt';
let allUsers = [];
let allMeals = [];
let allFoods = [];
let listEntries = [];

// DOM Elements
const app = document.getElementById('app');
const navLinks = document.querySelectorAll('nav a');

// Initialize the application
document.addEventListener('DOMContentLoaded', async () => {
  showLoading();
  try {
    // Set up navigation
    navLinks.forEach(link => {
      link.addEventListener('click', (e) => {
        e.preventDefault();
        const page = link.getAttribute('data-page');
        navigateTo(page);
      });
    });

    // Load initial data and render dashboard
    await loadInitialData();
  } catch (error) {
    console.error('Error during initialization:', error);
    showError('Failed to initialize application. Please refresh the page.');
  } finally {
    hideLoading();
  }
});

// Load initial data from API
async function loadInitialData() {
  showLoading();
  try {
    // Load all users
    const usersResponse = await fetch(`${API_BASE_URL}/dieters/all`);
    if (!usersResponse.ok) throw new Error('Failed to load users');
    allUsers = await usersResponse.json();

    const mealsResponse = await fetch(`${API_BASE_URL}/dieter/mealstoday`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name: currentUser })
    });

    if (!mealsResponse.ok) throw new Error('Failed to load meals for this dieter');
    allMeals = await mealsResponse.json();

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

async function navigateTo(page) {
  showLoading();
  try {
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

    // Render the page without showing/hiding loading
    await renderPageContent(page);
  } catch (error) {
    console.error('Navigation error:', error);
    showError('Failed to navigate to page. Please try again.');
  } finally {
    hideLoading();
  }
}

// Renamed and modified renderPage to not handle loading state
async function renderPageContent(page) {
  // Clear the app container
  app.innerHTML = '';

  // Get the template
  const template = document.getElementById(`${page}-template`);
  if (!template) {
    throw new Error(`Template for ${page} not found`);
  }

  // Clone the template content
  const content = template.content.cloneNode(true);
  app.appendChild(content);

  // Re-add the loading element
  const loadingElement = document.getElementById('loading');
  if (loadingElement) {
    app.appendChild(loadingElement);
  }

  // Initialize page-specific functionality
  switch (page) {
    case 'dashboard':
      await initDashboard();
      break;
    case 'meals':
      initMeals();
      break;
    case 'foods':
      initFoods();
      break;
    case 'users':
      await initUsers();
      break;
    case 'entries':
      initEntries();
      break;
  }
}

async function initDashboard() {
  try {
    const userSelect = document.getElementById('user-select');
    const dashboardUserName = document.getElementById('dashboard-user-name');
    const remainingCard = document.getElementById('remaining');
    const todayMeals = document.getElementById('today-meals');
    const mealsListContainer = document.getElementById('meals-list');
    const mealManagement = document.getElementById('meal-management');
    const caloriesSummary = document.getElementById('calorie-summary');
    const totalCaloriesCard = document.getElementById('daily-target');
    const consumedCaloriesCard = document.getElementById('consumed-today');

    if (!userSelect || !caloriesSummary || !todayMeals) {
      throw new Error('Required dashboard elements not found');
    }

    // Rest of the initialization code...
    user.populateUserSelect(allUsers, userSelect);

    userSelect.addEventListener('change', async () => {
      const selectedUserId = userSelect.value;

      if (!selectedUserId) {
        mealManagement.style.display = 'none';
        caloriesSummary.style.display = 'none';
        return;
      }

      showLoading();
      try {
        // Find selected user
        const selectedUser = allUsers.find(user => user.id.toString() === selectedUserId);
        dashboardUserName.textContent = selectedUser.name;
        currentUser = selectedUser;

        // Get user's meals -- implement in the future
        const mealsResponse = await fetch(`${API_BASE_URL}/dieter/mealstoday`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ name: selectedUser.name })
        });

        if (!mealsResponse.ok) throw new Error('Failed to load meals');
        const mealsData = await mealsResponse.json();

        meal.renderMealHistory(mealsData, mealsListContainer)
        todayMeals.style.display = 'block';

        const caloriesRemaining = await fetch(`${API_BASE_URL}/dieter/remaining`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ name: selectedUser.name })
        });

        const caloriesTotal = await fetch(`${API_BASE_URL}/dieter/calories`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ name: selectedUser.name })
        });

        if (!caloriesRemaining.ok) throw new Error('Failed to load remaining calories');
        const caloriesRemainingToday = await caloriesRemaining.json();

        if (!caloriesTotal.ok) throw new Error('Failed to load total calories');
        const caloriesTotalToday = await caloriesTotal.json();

        // Render meal history
        //renderMealHistory(mealsData, mealHistoryList);

        remainingCard.textContent = `${caloriesRemainingToday.calories}`;
        totalCaloriesCard.textContent = `${caloriesTotalToday}`;
        consumedCaloriesCard.textContent = `${caloriesTotalToday - caloriesRemainingToday.calories}`;

        caloriesSummary.style.display = 'block';

        // Show meal management
        //mealManagement.style.display = 'block';
      } catch (error) {
        console.error('Error loading meals data:', error);
        showError('Failed to load meals data. Please try again later.');
      } finally {
        hideLoading();
      }
    });
    userSelect.dispatchEvent(new Event('change'));
  } catch (error) {
    console.error('Dashboard initialization error:', error);
    hideLoading();
    showError('Failed to initialize dashboard. Please try again.');
  }

}

// Initialize Meals page
function initMeals() {
  const userSelect = document.getElementById('meal-user-select');
  const mealManagement = document.getElementById('meal-management');
  const addMealForm = document.getElementById('add-meal-form');
  const mealFoodsSelect = document.getElementById('meal-foods');
  const mealHistoryList = document.getElementById('meal-history-list');

  // Populate user select
  user.populateUserSelect(allUsers, userSelect);

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
    let totalCalories = 0;
    // Get selected foods
    const selectedFoods = Array.from(mealFoodsSelect.selectedOptions).map(option => {
      const foodId = option.value;
      return allFoods.find(f => f.id.toString() === foodId);
    });

    selectedFoods.forEach(food => {
      totalCalories += food.calories;
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
          calories: parseInt(totalCalories),
        })
      });

      if (!response.ok) throw new Error('Failed to add meal');
      const newMeal = await response.json();

      // add entries for each of the foods selected
      const selectedCalories = meal.setCaloriesSelected(selectedFoods, newMeal);

      // Update meal calories
      const mealCaloriesResponse = await fetch(`${API_BASE_URL}/meal/calories`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          id: newMeal.id,
          calories: selectedCalories,
        })
      });

      if (!mealCaloriesResponse.ok) throw new Error('Failed to update meal calories');

      allMeals.push(newMeal);

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

  userSelect.dispatchEvent(new Event('change'));
}

// Initialize Entry page
function initEntries() {
  const mealManagement = document.getElementById('entry-management');
  const addMealForm = document.getElementById('add-entry-form');
  const mealFoodsSelect = document.getElementById('entry-foods');
  const entryHistoryList = document.getElementById('entry-history-list');
  const entryMealSelect = document.getElementById('entry-meal-select');

  const populateResponse = meal.populateMealSelect(entryMealSelect);
  if (!populateResponse) {
    mealManagement.style.display = 'none';
    return;
  }

  entryMealSelect.addEventListener('change', async () => {
    let selectedMeal = allMeals.find(meal => meal.id.toString() === entryMealSelect.value);
    populateMealEntries(selectedMeal);
    const response = await fetch(`${API_BASE_URL}/meal/entries`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        id: selectedMeal.id,
      })
    });

    if (!response.ok) throw new Error('Failed to load entries');
    listEntries = await response.json();

    entry.renderEntryHistory(entryHistoryList, listEntries);
    entryHistoryList.style.display = 'block';
  });

  // Populate foods select
  allFoods.forEach(food => {
    const option = document.createElement('option');
    option.value = food.id;
    option.textContent = `${food.name} (${food.calories} cal)`;
    mealFoodsSelect.appendChild(option);
  });

  // Handle add meal form submission
  addMealForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    if (!currentUser) return;

    showLoading();
    try {
      const mealId = entryMealSelect.value;
      let selectedMeal = allMeals.find(meal => meal.id.toString() === mealId);

      // Add entries to the meal
      const selectedFoods = Array.from(mealFoodsSelect.selectedOptions).map(option => {
        const foodId = option.value;
        return allFoods.find(f => f.id.toString() === foodId);
      });

      let selectedCalories = meal.setCaloriesSelected(selectedFoods, selectedMeal);
      selectedMeal.calories = selectedMeal.calories + selectedCalories;

      const response = await fetch(`${API_BASE_URL}/meal/calories`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          id: selectedMeal.id,
          calories: parseInt(selectedMeal.calories),
        })
      });

      if (!response.ok) throw new Error('Failed to add entry');

      // Reset form
      addMealForm.reset();

      // Refresh meals
      entryMealSelect.dispatchEvent(new Event('change'));

      showSuccess('Entry added successfully!');
    } catch (error) {
      console.error('Error adding entry:', error);
      showError('Failed to add entry. Please try again later.');
    } finally {
      hideLoading();
    }
  });

  entryMealSelect.dispatchEvent(new Event('change'));

  mealManagement.style.display = 'block';
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
      food.renderFoodList(allFoods, foodListContainer);

      showSuccess('Food added successfully!');
    } catch (error) {
      console.error('Error adding food:', error);
      showError('Failed to add food. Please try again later.');
    } finally {
      hideLoading();
    }
  });
}

async function initUsers() {
  const addUserForm = document.getElementById('add-user-form');
  const addUserFormName = document.getElementById('new-user-name');
  const addUserFormCalories = document.getElementById('user-calories');
  const userListContainer = document.getElementById('user-list-container');

  if (!addUserForm || !userListContainer) {
    throw new Error('Required user page elements not found');
  }

  // Render user list without showing loading
  user.renderUserList(allUsers, userListContainer);

  // Handle add user form submission
  addUserForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    showLoading();
    try {
      await addDieter(addUserFormName.value, parseInt(addUserFormCalories.value));
    } finally {
      hideLoading();
    }
  });
}

// UI Helper Functions
function showLoading() {
  const loadingElement = document.getElementById('loading');
  if (loadingElement) {
    loadingElement.style.display = 'block';
  }
}

function hideLoading() {
  const loadingElement = document.getElementById('loading');
  if (loadingElement) {
    loadingElement.style.display = 'none';
  }
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
