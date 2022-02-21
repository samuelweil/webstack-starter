import React from "react";

interface LoggedInState {
  isLoggedIn: true;
  logout(): void;
}

interface LoggedOutState {
  isLoggedIn: false;
  login(): void;
}

export type AuthState = LoggedInState | LoggedOutState;

export const AuthContext = React.createContext<AuthState>(null!);
