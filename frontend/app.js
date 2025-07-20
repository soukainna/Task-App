const API = '/api';
;
const taskList = document.getElementById('taskList');
const taskForm = document.getElementById('taskForm');
const taskTitle = document.getElementById('taskTitle');
const statsText = document.getElementById('stats');

async function loadTasks() {
  const res = await fetch(`${API}/tasks`);
  const tasks = await res.json();
  taskList.innerHTML = '';

  tasks.forEach(task => {
    const li = document.createElement('li');
    li.className = 'list-group-item d-flex justify-content-between align-items-center';

    const title = document.createElement('span');
    title.textContent = task.title;
    if (task.completed) {
      title.classList.add('text-decoration-line-through', 'text-success');
    }

    const buttons = document.createElement('div');

    const toggleBtn = document.createElement('button');
    toggleBtn.className = 'btn btn-sm btn-outline-secondary me-2';
    toggleBtn.textContent = task.completed ? 'Annuler' : 'Terminer';
    toggleBtn.onclick = async () => {
      await fetch(`${API}/tasks/${task.ID}`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ completed: !task.completed })
      });
      loadTasks();
    };

    const deleteBtn = document.createElement('button');
    deleteBtn.className = 'btn btn-sm btn-outline-danger';
    deleteBtn.textContent = 'Supprimer';
    deleteBtn.onclick = async () => {
      await fetch(`${API}/tasks/${task.ID}`, { method: 'DELETE' });
      loadTasks();
    };

    buttons.appendChild(toggleBtn);
    buttons.appendChild(deleteBtn);

    li.appendChild(title);
    li.appendChild(buttons);
    taskList.appendChild(li);
  });

  loadStats();
}

taskForm.onsubmit = async (e) => {
  e.preventDefault();
  await fetch(`${API}/tasks`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title: taskTitle.value, completed: false })
  });
  taskTitle.value = '';
  loadTasks();
};

async function loadStats() {
  const res = await fetch(`${API}/stats`);
  const data = await res.json();
  statsText.textContent = `✅ Tâches complétées : ${data.completed_tasks}`;
}

loadTasks();
