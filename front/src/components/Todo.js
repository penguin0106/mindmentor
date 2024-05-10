import { useState, useEffect } from 'react';
import TodoItem from './TodoItem';
import "../App.css"

const TodoForm = ({ onSave, emotion, onDelete }) => {
  const [topic, setTopic] = useState('');
  const [body, setBody] = useState('');
  const [todos, setTodos] = useState([]);
  const [editingTodo, setEditingTodo] = useState(null);
  const [showModal, setShowModal] = useState(false);

  const createTodo = async (todo, event) => {
    try {
      const response = await fetch('http://localhost:8082/emotions', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(todo),
      });
      handleCloseModal()
      const responseGet = await fetch('http://localhost:8082/emotions/user');
      const data = await responseGet.json();
      setTodos(data);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const createdTodo = await response.json().catch((error) => {
        console.error('Error parsing JSON:', error);
        throw error;
      });
      return createdTodo;
    } catch (error) {
      console.error('Error creating todo:', error);
      throw error;
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const todoToSave = { topic, body, ...(editingTodo && { id: editingTodo.id }) };
      const createdTodo = await createTodo(todoToSave);
      onSave(createdTodo);
      setTopic('');
      setBody('');
      setEditingTodo(null);

    } catch (error) {
      console.error('Error creating todo:', error);
    }
  };



  const deleteTodo = async (id) => {
    try {
      await fetch(`http://localhost:8082/emotions/delete?id=${id}`, {
        method: 'DELETE',
      });

      // Update the UI to reflect the deleted todo
      setTodos((prevTodos) => prevTodos.filter((todo) => todo.id !== id));
    } catch (error) {
      console.error('Error deleting todo:', error);
    }
  };

  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const response = await fetch('http://localhost:8082/emotions/user');
        const data = await response.json();
        setTodos(data);
      } catch (error) {
        console.error('Error fetching todos:', error);
      }
    };

    fetchTodos();
  }, [,]);

  useEffect(() => {
    if (editingTodo) {
      const { topic: editingTopic, body: editingBody } = editingTodo;
      setTopic(editingTopic);
      setBody(editingBody);
    }
  }, [createTodo]);

  const handleTopicChange = (event) => {
    setTopic(event.target.value);
  };

  const handleBodyChange = (event) => {
    setBody(event.target.value);
  };


  const handleAddTodoClick = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setEditingTodo(null);
  };

  const handleTodoClick = (todo) => {
    setEditingTodo(todo);
    setShowModal(true);
  };

  return (
      <div>
        <h1 className="h1">Заметки</h1>
        <ul className="ul">

          {todos ? (
              todos.length > 0 ? (
                  todos.map((todo, id) => (
                      <TodoItem
                          key={id}
                          todo={todo}
                          onDelete={() => deleteTodo(todo.id)}
                          onClick={() => handleTodoClick(todo)}
                      />
                  ))
              ) : (
                  <div> Пусто</div>
              )
          ) : (
              <div> Loading...</div> // or some other fallback component
          )}
        </ul>
        <button className="plus" onClick={handleAddTodoClick}>
          +
        </button>
        {showModal && (
            <div className="modal">
              <div className="modal-content">
                <form onSubmit={handleSubmit}>
                  <label htmlFor="topic">Заголовок</label>
                  <input
                      defaultValue={editingTodo?.topic ?? ''}
                      type="text"
                      id="topic"
                      name="topic"
                      required
                      onChange={handleTopicChange}
                  />
                  <label htmlFor="body">Текст</label>
                  <input
                      defaultValue={editingTodo?.body ?? ''}
                      type="text"
                      id="body"
                      name="body"
                      required
                      onChange={handleBodyChange}
                  />
                  <button type="submit" >Save Todo</button>
                  <button type="button" onClick={handleCloseModal}>
                    Close
                  </button>
                </form>
              </div>
            </div>
        )}
      </div>
  );
};

export default TodoForm;