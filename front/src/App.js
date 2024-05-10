import * as React from 'react';
import  { useState } from 'react';
import ButtomPannl from "./components/ButtomPannl";
import SingIn from "./components/SingIn"
export default function App() {

    const [isRegistered, setIsRegistered] = useState(false);

    return (
        <>
            {!isRegistered && <SingIn setIsRegistered={setIsRegistered} />}
            {isRegistered && <ButtomPannl />}
        </>
    );
}
