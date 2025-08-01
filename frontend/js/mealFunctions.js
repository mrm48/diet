export async function deleteMeal(meal) {
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

// Helper function to render meals list
export function renderMealsList(meals, container) {
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
export function renderMealHistory(meals, container) {
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

// Allow user to select meals from the database
export async function populateMealSelect(allMeals, selectElement) {
  // Clear existing options
  selectElement.innerHTML = '';

  if (!allMeals || allMeals.length === 0) {
    return;
  }

  // Add meals
  allMeals.forEach(meal => {
    const option = document.createElement('option');
    option.value = meal.id;
    option.textContent = `` + meal.name + ``;
    selectElement.appendChild(option);
  });

}

export async function populateMealEntries(meal) {

  const response = await fetch(`${API_BASE_URL}/meal/entries`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      id: meal.id,
    })
  });

  if (!response.ok) throw new Error('Failed to load entries');
  listEntries = await response.json();

}


export function setCaloriesSelected(food, meal) {
  let selectedCalories = 0;

  food.forEach((item) => {
    selectedCalories += item.calories;
    const addEntryResponse = addEntry(item.calories, item.id, meal.id);
    if (!addEntryResponse.ok) throw new Error('Failed to set calories');
  })

  return selectedCalories;
}
