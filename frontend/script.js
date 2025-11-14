// Fetch tasks from API and display them
fetch('http://localhost:8080/showtasks')
    .then(response => response.json())
    .then(tasks => {
        const taskList = document.getElementById('task-list');
        tasks.forEach(task => {
            const taskItem = document.createElement('li');
            taskItem.textContent = task.name;
            taskList.appendChild(taskItem);
        });
    })
    .catch(error => console.error('Error fetching tasks:', error));

// Add event listener to add task button
document.getElementById('add-task-btn').addEventListener('click', () => {
    const taskInput = document.getElementById('task-input');
    const taskName = taskInput.value.trim();
    if (taskName) {
        // Send request to add new task
        fetch('http://localhost:8080/addtask', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name: taskName })
        })
        .then(response => response.json())
        .then(newTask => {
            const taskList = document.getElementById('task-list');
            const taskItem = document.createElement('li');
            taskItem.textContent = newTask.name;
            taskList.appendChild(taskItem);
            taskInput.value = ''; // Clear input field
        })
        .catch(error => console.error('Error adding task:', error));
    }
});
