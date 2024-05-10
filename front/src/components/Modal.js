import React from 'react';

function Modal({ editingTodo, onClose, onSubmit,fetchEmotionsData }) {
    return (
        <div className="modal">
            <div className="modal-content">
                <form onSubmit={onSubmit}>
                    <label htmlFor="todoText">Todo text:</label>
                    <input
                        type="text"
                        id="todoText"
                        name="todoText"
                        defaultValue={editingTodo? editingTodo.text : ''}
                        required
                    />
                    <button type="submit" onClick={fetchEmotionsData}>Save Todo</button>
                    <button type="button" onClick={onClose}>Close</button>
                </form>
            </div>
        </div>
    );
}

export default Modal;