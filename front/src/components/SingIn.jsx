import * as React from 'react';
import  { useState } from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';


function Copyright(props) {
    return (
        <Typography variant="body2" color="text.secondary" align="center" {...props}>
            {'Copyright © '}
            <Link color="inherit" href="https://mui.com/">
                Your Website
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    );
}

const defaultTheme = createTheme();


export default function SingIn({setIsRegistered}) {
    const initialState = {
        password: '',
        email: '',
        firstName: '',
    };
    const [formData, setFormData] = useState(initialState);
    const [signIn, setSignIn] = useState(true);
    const [user, setUser] = useState(null);

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setFormData((prevFormData) => ({...prevFormData, [name]: value }));
    };

    const handleLogin = async (event) => {
        event.preventDefault();
        try {
              const response = await fetch('http://localhost:8081/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(formData),
                  });
              if (response.ok) {
                    const user = { email: formData.email };
                    console.log('Authentication successful');
                    setUser(user);
                    setIsRegistered(true)
                  } else {
                    throw new Error('Authentication failed');
                  }
            } catch (error) {
              console.error('Error during authentication:', error);

        }
    };

    const handleRegister = async (event) => {
        event.preventDefault();
        try {
            const response = await fetch('http://localhost:8081/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email: formData.email,
                    password: formData.password,
                    firstName: formData.firstName,
                }),
            });

            if (response.ok) {
                console.log('Registration successful');
                setIsRegistered(true)
            } else {
                console.error('Registration failed:', response.status, response.statusText);
                throw new Error('Registration failed');
            }
        } catch (error) {
            console.error('Error during registration:', error);
        }
    };

      const handleSignUpClick = () => {
            setSignIn(false);
          };

      const handleSignInClick = () => {
            setSignIn(true);
          };

    return (
        <>
            {signIn ? (
                <ThemeProvider theme={defaultTheme}>
                        <Container component="main" maxWidth="xs">
                            <CssBaseline />
                            <Box
                                sx={{
                                    marginTop: 8,
                                    display: 'flex',
                                    flexDirection: 'column',
                                    alignItems: 'center',
                                }}
                            >
                                <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                                    <LockOutlinedIcon />
                                </Avatar>
                                <Typography component="h1" variant="h5">
                                    Войти
                                </Typography>
                                <Box component="form" onSubmit={handleLogin}  noValidate sx={{ mt: 1 }}>
                                    <TextField
                                        margin="normal"
                                        required
                                        fullWidth
                                        id="email"
                                        label="Электронная почта"
                                        name="email"
                                        autoComplete="email"
                                        value={formData.email}
                                        onChange={handleInputChange}
                                    />
                                    <TextField
                                        margin="normal"
                                        required
                                        fullWidth
                                        name="password"
                                        label="Пароль"
                                        type="password"
                                        id="password"
                                        autoComplete="current-password"
                                        value={formData.password}
                                        onChange={handleInputChange}
                                    />

                                    <Button
                                        type="submit"
                                        fullWidth
                                        variant="contained"
                                        sx={{ mt: 3, mb: 2 }}
                                    >
                                        Войти
                                    </Button>
                                    <Grid container>
                                        <Grid item>
                                            <Link href="#" variant="body2" onClick={handleSignUpClick}>
                                                {"Не имеете аккаунта? Зарегистрируйтесь"}
                                            </Link>
                                        </Grid>
                                    </Grid>
                                </Box>
                            </Box>

                        </Container>
                    )
                </ThemeProvider>
            ) : (
                <ThemeProvider theme={defaultTheme}>
                    <Container component="main" maxWidth="xs">
                        <CssBaseline />
                        <Box
                            sx={{
                                marginTop: 8,
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'center',
                            }}
                        >
                            <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                                <LockOutlinedIcon />
                            </Avatar>
                            <Typography component="h1" variant="h5">
                                Регистрация
                            </Typography>
                            <Box component="form" onSubmit={handleRegister} sx={{mt: 3}}>

                                <label>Имя</label>
                                <input
                                    type="text"
                                    name="firstName"
                                    id="firstName"
                                    label="Имя"
                                    value={formData.firstName}
                                    onChange={handleInputChange}
                                    required
                                    className="form-input"
                                />
                                <br/>

                                <label>Электронная почта</label>
                                <input
                                    type="email"
                                    name="email"
                                    id="email"
                                    label="Электронная почта"
                                    value={formData.email}
                                    onChange={handleInputChange}
                                    required
                                    className="form-input"
                                />
                                <br/>

                                <label>Пароль</label>
                                <input
                                    type="password"
                                    name="password"
                                    id="password"
                                    label="Пароль"
                                    value={formData.password}
                                    onChange={handleInputChange}
                                    required
                                    className="form-input"
                                />
                                <br/>

                                <button type="submit" onClick={handleRegister} className="form-submit">Регистрация
                                </button>

                                <Grid container justifyContent="flex-end">
                                    <Grid item>
                                        <Link href="#" variant="body2" onClick={handleSignInClick}>
                                            {signIn ? "Не имеете аккаунта? Зарегистрируйтесь" : "Уже есть аккаунт? Войдите"}
                                        </Link>
                                    </Grid>
                                </Grid>
                            </Box>
                        </Box>
                        <Copyright sx={{mt: 5}}/>
                </Container>
                </ThemeProvider>
                )}
</>
)
;
}
