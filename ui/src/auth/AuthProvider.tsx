import React, { ReactNode } from "react";
import { AuthState, useAuthState } from "./use-auth-state";

export const AuthContext = React.createContext<AuthState>(null!);

export function AuthProvider(props: { children: ReactNode }) {
  const authState = useAuthState();

  return (
    <AuthContext.Provider value={authState}>
      {props.children}
    </AuthContext.Provider>
  );
}
