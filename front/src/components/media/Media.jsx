import * as React from 'react';
import Box from '@mui/material/Box';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import MusicPlayerComponent from "../music/MusicPlayerComponent";
import VideocamIcon from '@mui/icons-material/Videocam';
import AudiotrackIcon from '@mui/icons-material/Audiotrack';
import VideoGridComponent from "../video/video";
import "./media.css"

function refreshMessages() {
    const getRandomInt = (max) => Math.floor(Math.random() * Math.floor(max));
}

export default function SimpleBottomNavigation() {
    const [value, setValue] = React.useState(0);
    const containerRef = React.useRef(null);
    const [messages, setMessages] = React.useState(() => refreshMessages());
    const [currentPage, setCurrentPage] = React.useState(0);

    React.useEffect(() => {
        if (containerRef.current) {
            containerRef.current.scrollTop = 0;
            setMessages(refreshMessages());
        }
    }, [value, setMessages, currentPage]);

    const handleChange = (event, newValue) => {
        setValue(newValue);
        setCurrentPage(newValue);
    };

    return (
        <div className="media">
            <Box ref={containerRef}>
                <BottomNavigation
                    showLabels
                    value={currentPage} // Use currentPage instead of value
                    onChange={handleChange}
                >
                    <BottomNavigationAction label="Музыка" icon={<AudiotrackIcon />} />
                    <BottomNavigationAction label="Видео" icon={<VideocamIcon />} />
                </BottomNavigation>
                {currentPage === 0 && <MusicPlayerComponent />}
                {currentPage === 1 && <VideoGridComponent />}
            </Box>
        </div>
    );
}