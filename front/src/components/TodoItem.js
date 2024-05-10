import React from 'react';
import { Trash2} from "lucide-react";

function TodoItem({ todo, onDelete, onClick, id }) {
    return (
        <li className="li" key={id}>
            <p style={{ display: "flex", flexDirection: "column", justifyContent: "center" }} onClick={onClick}>

                <h1 style={{fontSize:"19px"}}> {todo.topic}</h1>
                {todo.body}
            </p>
            <Trash2 size={20} className="icon" onClick={onDelete} />
        </li>
    );
}

export default TodoItem;