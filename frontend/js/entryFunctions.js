export async function deleteEntry(entry) {
    try {
        const response = await fetch(`${API_BASE_URL}/entry`, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(entry)
        });

        if (!response.ok) throw new Error('Failed to delete entry');

        const mealCaloriesResponse = await fetch(`${API_BASE_URL}/meal`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                id: entry.meal,
            })
        })

        if (!mealCaloriesResponse.ok) throw new Error('Failed to load meal calories');
        const mealCalories = await mealCaloriesResponse.json();

        const newCalorieTotal = mealCalories.calories - entry.calories;

        const mealResponse = await fetch(`${API_BASE_URL}/meal/calories`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                id: entry.meal,
                calories: newCalorieTotal
            })
        })

        if (!mealResponse.ok) throw new Error('Failed to update meal calories');

        showSuccess('Entry deleted successfully!');
        return true;
    } catch (error) {
        console.error('Error deleting entry:', error);
        showError('Failed to delete entry. Please try again later.');
        return false;
    }
}

export async function addEntry(foodCalories, foodID, newMealID) {
    const responseEntry = await fetch(`${API_BASE_URL}/entry`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            food: foodID,
            meal: newMealID,
            calories: foodCalories,
        })
    });
    if (!responseEntry.ok) throw new Error('Failed to add entry');
}


export function renderEntryHistory(container, listEntries) {

    container.innerHTML = '';

    if (!listEntries || listEntries.length === 0) {
        container.innerHTML = '<p>No entries found.</p>';
        return;
    }

    // Add entries
    const entryHeader = document.createElement('h4');
    entryHeader.textContent = "Entries";
    container.appendChild(entryHeader);

    listEntries.forEach(entry => {
        if (entry.food) {
            const entryCard = document.createElement('div');
            entryCard.className = 'meal-card';
            entryCard.innerHTML = `
            <h4>${allFoods[entry.food - 1].name}</h4>
            <p><strong>Calories:</strong> ${entry.calories}</p>
            <div class="actions">
            <button class="delete-btn" data-id="${entry.id}">Delete</button>
            </div>
            `;

            // Add delete event listener
            const deleteBtn = entryCard.querySelector('.delete-btn');
            deleteBtn.addEventListener('click', async () => {
                if (confirm(`Are you sure you want to delete ${allFoods[entry.food - 1].name}?`)) {
                    await deleteEntry(entry);
                    entryCard.remove();
                }
            });

            container.appendChild(entryCard);

        }
    });
}
