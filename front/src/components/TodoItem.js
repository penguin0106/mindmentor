import React from 'react';
import { Trash2} from "lucide-react";

function TodoItem({ todo, onDelete, onClick,body,topic }) {
    return (
        <li className="li">
            <p style={{ width: "90%", margin: "0", padding: "10px" }} onClick={onClick}>

                {
                    body
                }

            </p>
            <Trash2 size={20} className="icon" onClick={onDelete} />
        </li>
    );
}

export default TodoItem;