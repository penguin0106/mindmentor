import * as React from 'react';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import SelfImprovementIcon from '@mui/icons-material/SelfImprovement';
import MessageIcon from '@mui/icons-material/Message';
import Paper from '@mui/material/Paper';
import Todo from "./Todo";
import Chat from "./chat/Chat";

import SummarizeIcon from '@mui/icons-material/Summarize';
import BookComponent from "./book/Book";
import AutoStoriesIcon from '@mui/icons-material/AutoStories';
import Media from "./media/Media";


function refreshMessages() {
    const getRandomInt = (max) => Math.floor(Math.random() * Math.floor(max));

}

export default function ButtomPannl() {
    const [value, setValue] = React.useState(0);
    const ref = React.useRef(null);
    const [messages, setMessages] = React.useState(() => refreshMessages());
    const [currentPage, setCurrentPage] = React.useState(0);

    React.useEffect(() => {
        ref.current.ownerDocument.body.scrollTop = 0;
        setMessages(refreshMessages());
    }, [value, setMessages, currentPage]);

    return (
        <Box sx={{ pb: 7 }} ref={ref}>
            <CssBaseline />
            {currentPage === 0 && <Todo />}
            {currentPage === 1 && <Chat />}
            {currentPage === 2 && <Media />}
            {currentPage === 3 && <BookComponent />}
            <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
                <BottomNavigation
                    showLabels
                    value={value}
                    onChange={(event, newValue) => {
                        setValue(newValue);
                    }}
                >
                    <BottomNavigationAction label="Заметки" icon={<SummarizeIcon />} onClick={() => setCurrentPage(0)} />
                    <BottomNavigationAction label="Общий чат" icon={<MessageIcon />} onClick={() => setCurrentPage(1)} />
                    <BottomNavigationAction label="Медитация" icon={<SelfImprovementIcon />} onClick={() => setCurrentPage(2)} />
                    <BottomNavigationAction label="Библиотека" icon={<AutoStoriesIcon />} onClick={() => setCurrentPage(3)} />
                </BottomNavigation>
            </Paper>
        </Box>
    );
}
