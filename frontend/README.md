# MauIt Frontend

This directory contains the frontend code for the MauIt Diet Tracker application.

## Directory Structure

- `index.html` - The main HTML file that serves as the entry point for the application
- `css/` - Contains CSS stylesheets
  - `styles.css` - Main stylesheet for the application
- `js/` - Contains JavaScript files
  - `main.js` - Main JavaScript file that handles application logic and API integration

## How It Works

The frontend is a single-page application (SPA) that uses HTML templates and JavaScript to dynamically render different views. The application communicates with the backend API to retrieve and store data.

### Key Features

1. **Dashboard** - Shows a summary of a user's calorie consumption for the day
2. **Meals** - Allows users to add and manage meals
3. **Foods** - Allows users to add and manage foods in the database
4. **Users** - Allows management of users and their daily calorie targets

### API Integration

The frontend communicates with the backend API using fetch requests. The API base URL is configured in `js/main.js`:

```javascript
const API_BASE_URL = 'http://localhost:9090';
```

If you need to change the API URL (e.g., for deployment), update this constant.

## Customization

### Styling

To customize the appearance of the application, edit the `css/styles.css` file. The stylesheet uses CSS variables for colors, which can be easily changed:

```css
:root {
    --primary-color: #3498db;
    --secondary-color: #2ecc71;
    --accent-color: #e74c3c;
    --background-color: #f5f5f5;
    --text-color: #333;
    --card-background: #fff;
    --border-color: #ddd;
}
```

### Adding New Features

To add new features to the frontend:

1. Add a new template in `index.html`
2. Add a new navigation link in the header
3. Create a new initialization function in `js/main.js`
4. Update the `renderPage` function to handle the new page

## Deployment

The frontend is served by the Go backend, so no separate deployment is needed. When you run the backend server, it will automatically serve the frontend files.