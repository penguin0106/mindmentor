import React, { useState } from 'react';
import styled from 'styled-components';

const ChatContainer = styled.div`
  display: flex;
  height: 90vh;
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  box-sizing: border-box;
`;

const ChatList = styled.div`
  display: flex;
  flex-direction: column;
  width: 200px;
  border: 1px solid #ccc;
`;

const ChatListHeader = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 10px;
`;

const ChatListTitle = styled.h3`
  font-size: 18px;
  font-weight: bold;
  margin: 0;
`;

const ChatListItem = styled.div`
  padding: 5px 10px;
  cursor: pointer;
  background-color: ${(props) => (props.isActive ? '#e0e0e0' : 'white')};
`;
const ChatRoom = styled.div`
  flex-grow: 1;
  display: flex;
  flex-direction: column;
`;

const ChatHeader = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 20px;
`;

const ChatTitle = styled.h2`
  font-size: 24px;
  font-weight: bold;
  margin: 0;
`;

const ChatMessages = styled.div`
  flex-grow: 1;
  overflow-y: auto;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
`;

const ChatMessage = styled.div`
  display: flex;
  align-items: center;
  margin-bottom: 10px;

  &:last-child {
    margin-bottom: 0;
  }
`;

const ChatMessageText = styled.p`
  font-size: 16px;
  margin: 0;
  padding: 10px;
  border-radius: 4px;
  background-color: #f5f5f5;
  box-sizing: border-box;
`;

const ChatInput = styled.input`
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
`;

const ChatRooms = [
    { id: 1, name: 'Chat Room 1' },
    { id: 2, name: 'Chat Room 2' },
    { id: 3, name: 'Chat Room 3' },
];

const Chat = () => {
    const [selectedRoomId, setSelectedRoomId] = useState(1);
    const [messages, setMessages] = useState([]);
    const [inputValue, setInputValue] = useState('');

    const selectedRoom = ChatRooms.find((room) => room.id === selectedRoomId);

    const handleRoomClick = (roomId) => {
        setSelectedRoomId(roomId);
        setMessages([]);
    };

    const handleInputChange = (event) => {
        setInputValue(event.target.value);
    };

    const handleInputSubmit = (event) => {
        event.preventDefault();
        setMessages([...messages, inputValue]);
        setInputValue('');
    };

    return (
        <ChatContainer>
            <ChatList>
                <ChatListHeader>
                    <ChatListTitle>Chats</ChatListTitle>
                </ChatListHeader>
                {ChatRooms.map((room) => (
                    <ChatListItem
                        key={room.id}
                        isActive={room.id === selectedRoomId}
                        onClick={() => handleRoomClick(room.id)}
                    >
                        {room.name}
                    </ChatListItem>
                ))}
            </ChatList>
            <ChatRoom>
                <ChatHeader>
                    <ChatTitle>{selectedRoom ? selectedRoom.name : 'Select a chat room'}</ChatTitle>
                </ChatHeader>
                <ChatMessages>
                    {messages.map((message, index) => (
                        <ChatMessage key={index}>
                            <ChatMessageText>{message}</ChatMessageText>
                        </ChatMessage>
                    ))}
                </ChatMessages>
                <form onSubmit={handleInputSubmit}>
                    <ChatInput
                        type="text"
                        value={inputValue}
                        onChange={handleInputChange}
                        placeholder="Type your message here..."
                    />
                </form>
            </ChatRoom>
        </ChatContainer>
    );
};

export default Chat;