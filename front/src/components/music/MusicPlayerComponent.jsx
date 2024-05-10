import React, {useState} from 'react';
import "./music.scss"

const MusicPlayerComponent = () => {
    const songs = [
        {
            src: 'https://example.com/song1.mp3',
            title: 'Song 1',
            artist: 'Artist 1',
            cover: 'https://example.com/cover1.jpg',
            duration: '3:24'
        },
        {
            src: 'https://example.com/song2.mp3',
            title: 'Song 2',
            artist: 'Artist 2',
            cover: 'https://example.com/cover2.jpg',
            duration: '4:09'
        },
    ];

    const [selectedSong, setSelectedSong] = useState(songs[0]);

    return (
        <div>
            <section id="playstation">
                <div id="controlpanel">
                    <div id="backward" className="inlineblo"><i className="fa fa-backward"></i></div>
                    <div id="songpro" className="inlineblo"></div>
                    <div id="forward" className="inlineblo"><i className="fa fa-forward"></i></div>
                    <h1 className="h1">{selectedSong.title}</h1>
                    <h2 className="h2">{selectedSong.artist}</h2>
                </div>
                <progress max="100" value="80"></progress>
                <ol className="ol">
                    {songs.map((song, index) => (
                        <li className="li" key={index} onClick={() => setSelectedSong(song)}>
                            {song.title} <span className="time">{song.duration}</span>
                        </li>
                    ))}
                </ol>
            </section>
        </div>
    );
};

export default MusicPlayerComponent;


