import React from "react";
import { User } from "./user";

interface LoggedInState {
  isLoggedIn: true;
  logout(): void;
  user: User;
}

interface LoggedOutState {
  isLoggedIn: false;
  login(): void;
}

export type AuthState = LoggedInState | LoggedOutState;

export const AuthContext = React.createContext<AuthState>(null!);
