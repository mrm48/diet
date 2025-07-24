// Helper function to render food list
export async function renderFoodList(foods, container) {
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

export async function deleteFood(food) {
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
