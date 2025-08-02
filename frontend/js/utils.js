export function showError(app, message) {
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

export function showSuccess(app, message) {
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
